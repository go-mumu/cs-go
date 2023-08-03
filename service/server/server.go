/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-26
 * File: server.go
 * Desc:
 */

package server

import (
	"fmt"
	"github.com/go-mumu/cs-go/library/log"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	defaultGrpcHandlerTimeout = 5 * time.Second
	defaultHttpQuitTimeout    = 5 * time.Second

	defaultGatewayPattern = "/"
)

type Server struct {
	grpcServer      *grpc.Server
	grpcAddr        string
	grpcRegister    GrpcRegister
	grpcMiddlewares []grpc.UnaryServerInterceptor

	httpServer      *http.Server
	httpServerMux   *http.ServeMux
	httpAddr        string
	httpRegister    HttpRegister
	httpMiddlewares []func(h http.Handler) http.Handler
	httpInitDone    chan struct{}
	httpOptions     []runtime.ServeMuxOption

	gatewayPattern string
	maxBodySize    int

	grpcIdleTimeout time.Duration
	httpIdleTimeout time.Duration

	grpcHandlerTimeout time.Duration
	httpHandlerTimeout time.Duration

	httpReadTimeout  time.Duration
	httpWriteTimeout time.Duration
	httpQuitTimeout  time.Duration

	errChan chan error
}

func NewServer() *Server {
	return &Server{
		grpcMiddlewares: make([]grpc.UnaryServerInterceptor, 0),
		httpMiddlewares: make([]func(next http.Handler) http.Handler, 0),
		httpOptions:     make([]runtime.ServeMuxOption, 0),
		httpServerMux:   http.NewServeMux(),
		gatewayPattern:  defaultGatewayPattern,

		grpcHandlerTimeout: defaultGrpcHandlerTimeout,
		httpQuitTimeout:    defaultHttpQuitTimeout,
		httpHandlerTimeout: 0,
		httpReadTimeout:    0,
		httpWriteTimeout:   0,
		httpIdleTimeout:    0,

		httpInitDone: make(chan struct{}, 1),
		errChan:      make(chan error),
	}
}

func (s *Server) SetGrpcRegister(register GrpcRegister) {
	s.grpcRegister = register
}

func (s *Server) SetHttpRegister(register HttpRegister) {
	s.httpRegister = register
}

func (s *Server) SetGrpcAddr(addr string) {
	s.grpcAddr = addr
}

func (s *Server) SetHttpAddr(addr string) {
	s.httpAddr = addr
}

func (s *Server) SetMaxBodySize(size int) {
	s.maxBodySize = size
}

func (s *Server) SetGrpcIdleTimeout(t int) {
	s.grpcIdleTimeout = time.Millisecond * time.Duration(t)
}

func (s *Server) AddGrpcMiddleware(interceptor grpc.UnaryServerInterceptor) {
	s.grpcMiddlewares = append(s.grpcMiddlewares, interceptor)
}

func (s *Server) AddHttpMiddleware(middleware func(next http.Handler) http.Handler) {
	s.httpMiddlewares = append(s.httpMiddlewares, middleware)
}

func (s *Server) AddHttpOption(option runtime.ServeMuxOption) {
	s.httpOptions = append(s.httpOptions, option)
}

func (s *Server) AddHttpHandle(pattern string, handler http.Handler) {
	s.httpServerMux.Handle(pattern, handler)
}

func (s *Server) AddHttpHandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	s.httpServerMux.HandleFunc(pattern, handler)
}

func (s *Server) SetGatewayPattern(pattern string) {
	s.gatewayPattern = pattern
}

func (s *Server) SetGrpcHandlerTimeout(t int) {
	s.grpcHandlerTimeout = time.Millisecond * time.Duration(t)
}

func (s *Server) SetHttpHandlerTimeout(t int) {
	s.httpHandlerTimeout = time.Millisecond * time.Duration(t)
}

func (s *Server) SetHttpReadTimeout(t int) {
	s.httpReadTimeout = time.Millisecond * time.Duration(t)
}

func (s *Server) SetHttpWriteTimeout(t int) {
	s.httpWriteTimeout = time.Millisecond * time.Duration(t)
}

func (s *Server) SetHttpIdleTimeout(t int) {
	s.httpIdleTimeout = time.Millisecond * time.Duration(t)
}

func (s *Server) SetHttpQuitTimeout(t int) {
	s.httpQuitTimeout = time.Millisecond * time.Duration(t)
}

/*// SetDefaultMarshaler 设置默认编解码方式
func (s *Server) SetDefaultMarshaler() {
	marshaler := &runtime.JSONBuiltin{}

	opt1 := runtime.WithMarshalerOption("application/x-www-form-urlencoded", &FORMPb{Json: marshaler})
	opt2 := runtime.WithMarshalerOption(runtime.MIMEWildcard, marshaler)

	s.httpOptions = append(s.httpOptions, opt1, opt2)
}

// SetJsonPbMarshaler 设置为使用JsonPb编解码
func (s *Server) SetJsonPbMarshaler() {
	marshaler := &runtime.JSONPb{OrigName: true, EmitDefaults: true}

	opt1 := runtime.WithMarshalerOption("application/x-www-form-urlencoded", &FORMPb{Json: marshaler})
	opt2 := runtime.WithMarshalerOption(runtime.MIMEWildcard, marshaler)

	s.httpOptions = append(s.httpOptions, opt1, opt2)
}*/

func (s *Server) Run() error {
	go func() {
		s.errChan <- s.runGrpcServer()
	}()

	go func() {
		s.errChan <- s.runHttpServer()
	}()

	return s.wait()
}

func (s *Server) runGrpcServer() error {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(s.grpcMiddlewares...),
		),
	}

	if s.maxBodySize > 0 {
		opts = append(
			opts,
			grpc.MaxRecvMsgSize(s.maxBodySize),
			grpc.MaxSendMsgSize(s.maxBodySize),
		)
	}

	if s.grpcIdleTimeout > 0 {
		opts = append(
			opts,
			grpc.KeepaliveParams(keepalive.ServerParameters{
				MaxConnectionIdle: s.grpcIdleTimeout,
			}),
		)
	}

	s.grpcServer = grpc.NewServer(opts...)
	s.grpcRegister(s.grpcServer)

	lis, err := net.Listen("tcp", s.grpcAddr)
	if err != nil {
		log.Cli.Error("grpc listen failed!", "error", err)
		return fmt.Errorf("grpc listen failed! error:%s", err)
	}

	log.Cli.Info("grpc server running!", "addr", s.grpcAddr)

	reflection.Register(s.grpcServer)

	if err = s.grpcServer.Serve(lis); err != nil {
		log.Cli.Error("grpc serve failed!", "error", err)
		return fmt.Errorf("grpc Serve failed! error:%s", err)
	}

	return nil
}

// runHTTPServer 启动http服务
func (s *Server) runHttpServer() error {
	ctx, _ := context.WithCancel(context.Background())

	runtimeMux := runtime.NewServeMux(s.httpOptions...)

	runtime.DefaultContextTimeout = s.grpcHandlerTimeout

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if s.maxBodySize > 0 {
		opts = append(
			opts,
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(s.maxBodySize)),
			grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(s.maxBodySize)),
		)
	}

	if err := s.httpRegister(ctx, runtimeMux, s.grpcAddr, opts); err != nil {
		log.Cli.Error("http handler failed!", "error", err)
		return fmt.Errorf("http handler failed! error:%s", err)
	}

	// gateway handle
	s.httpServerMux.Handle(s.gatewayPattern, runtimeMux)

	s.httpServer = &http.Server{
		Addr:    s.httpAddr,
		Handler: s.httpServerMux,
	}

	// set Timeout
	{
		if s.httpHandlerTimeout != 0 {
			s.httpServer.Handler = http.TimeoutHandler(s.httpServer.Handler, s.httpHandlerTimeout, "503 Handler timeout")
		}

		if s.httpReadTimeout != 0 {
			s.httpServer.ReadTimeout = s.httpReadTimeout
		}

		if s.httpWriteTimeout != 0 {
			s.httpServer.WriteTimeout = s.httpWriteTimeout
		}

		if s.httpIdleTimeout != 0 {
			s.httpServer.IdleTimeout = s.httpIdleTimeout
		}
	}

	// add middleware, 先添加的在最外层
	for i := len(s.httpMiddlewares) - 1; i >= 0; i-- {
		s.httpServer.Handler = s.httpMiddlewares[i](s.httpServer.Handler)
	}

	log.Cli.Info("http server running!", "addr", s.httpAddr)

	s.httpInitDone <- struct{}{}

	if err := s.httpServer.ListenAndServe(); err != nil {
		log.Cli.Error("http listen stoped!", "error", err)
		return fmt.Errorf("http listen failed! error:%s", err)
	}

	return nil
}

func (s *Server) wait() error {
	var err error

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case sig := <-c:
		switch sig {
		case syscall.SIGTERM, syscall.SIGINT:
			<-s.httpInitDone
			ctx, _ := context.WithTimeout(context.Background(), s.httpQuitTimeout)
			err = s.httpServer.Shutdown(ctx)
			if err != nil {
				return err
			}
			s.grpcServer.GracefulStop()
		case syscall.SIGQUIT:
			<-s.httpInitDone
			err = s.httpServer.Close()
			if err != nil {
				return err
			}
			s.grpcServer.Stop()
		}

		log.Cli.Error("server closed got signal, shutdown success", "signal", sig.String())
		return fmt.Errorf("server closed got signal %s, shutdown success", sig)
	case err = <-s.errChan:
		return fmt.Errorf("server closed %s", err)
	}
}
