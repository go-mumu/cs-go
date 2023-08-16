/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-26
 * File: server.go
 * Desc:
 */

package server

import (
	"context"
	"fmt"
	"github.com/go-mumu/cs-go/library/config"
	"github.com/go-mumu/cs-go/library/log"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	defaultGrpcHandlerTimeout = 5 * time.Second
	defaultHttpQuitTimeout    = 5 * time.Second

	defaultGatewayPattern = "/"
	defaultMaxMsgSize     = 10 * 1024 * 1024
)

type Server struct {
	// About gRPC
	// server
	gRPCServer *grpc.Server

	// options
	gRPCIp             string
	gRPCPort           string
	gRPCRegister       GRPCRegister
	gRPCMiddlewares    []grpc.UnaryServerInterceptor
	gRPCHandlerTimeout time.Duration

	maxConnectionIdle time.Duration

	// About HTTP
	// server
	httpServer   *http.Server
	httpServeMux *http.ServeMux

	// options
	httpIp             string
	httpPort           string
	httpRegister       HTTPRegister
	httpMiddlewares    []func(handler http.Handler) http.Handler
	httpServeMuxOption []runtime.ServeMuxOption
	httpReadTimeout    time.Duration
	httpWriteTimeout   time.Duration
	httpIdleTimeout    time.Duration
	httpHandlerTimeout time.Duration
	httpQuitTimeout    time.Duration

	httpInitDone chan struct{}

	gatewayPattern string

	// maxMsgSize set the max message size in bytes the server can receive/send.
	maxMsgSize int

	// error chan
	errChan chan error
}

// NewServer init server struct
func NewServer() *Server {
	return &Server{
		gRPCMiddlewares:    make([]grpc.UnaryServerInterceptor, 0),
		httpMiddlewares:    make([]func(next http.Handler) http.Handler, 0),
		httpServeMuxOption: make([]runtime.ServeMuxOption, 0),
		httpServeMux:       http.NewServeMux(),

		gatewayPattern: defaultGatewayPattern,
		maxMsgSize:     defaultMaxMsgSize,

		httpQuitTimeout:    defaultHttpQuitTimeout,
		gRPCHandlerTimeout: defaultGrpcHandlerTimeout,

		httpInitDone: make(chan struct{}, 1),
		errChan:      make(chan error),
	}
}

// SetGRPCIp set gRPC ip
func (s *Server) SetGRPCIp(ip string) {
	s.gRPCIp = ip
}

// SetGRPCPort set gRPC port
func (s *Server) SetGRPCPort(port string) {
	s.gRPCPort = port
}

// SetGRPCRegister register grpc handler
func (s *Server) SetGRPCRegister(register GRPCRegister) {
	s.gRPCRegister = register
}

func (s *Server) SetGRPCHandlerTimeout(t int) {
	s.gRPCHandlerTimeout = time.Duration(t) * time.Millisecond
}

// AddGRPCMiddleware add gRPC middleware
func (s *Server) AddGRPCMiddleware(interceptor grpc.UnaryServerInterceptor) {
	s.gRPCMiddlewares = append(s.gRPCMiddlewares, interceptor)
}

// SetMaxConnectionIdle set max connection idle
func (s *Server) SetMaxConnectionIdle(t int) {
	s.maxConnectionIdle = time.Duration(t) * time.Millisecond
}

// SetHTTPIp set http ip
func (s *Server) SetHTTPIp(ip string) {
	s.httpIp = ip
}

// SetHTTPPort set http port
func (s *Server) SetHTTPPort(port string) {
	s.httpPort = port
}

// SetHTTPRegister set http register
func (s *Server) SetHTTPRegister(register HTTPRegister) {
	s.httpRegister = register
}

// AddHTTPMiddleware add http middleware
func (s *Server) AddHTTPMiddleware(middleware func(next http.Handler) http.Handler) {
	s.httpMiddlewares = append(s.httpMiddlewares, middleware)
}

// AddHTTPServeMuxOption add http serve mux option
func (s *Server) AddHTTPServeMuxOption(option runtime.ServeMuxOption) {
	s.httpServeMuxOption = append(s.httpServeMuxOption, option)
}

// SetHTTPReadTimeout set http read timeout
func (s *Server) SetHTTPReadTimeout(t int) {
	s.httpReadTimeout = time.Duration(t) * time.Millisecond
}

// SetHTTPWriteTimeout set http read timeout
func (s *Server) SetHTTPWriteTimeout(t int) {
	s.httpWriteTimeout = time.Duration(t) * time.Millisecond
}

// SetHTTPHandlerTimeout set http handler timeout
func (s *Server) SetHTTPHandlerTimeout(t int) {
	s.httpHandlerTimeout = time.Duration(t) * time.Millisecond
}

// SetHTTPIdleTimeout set http idle timeout
func (s *Server) SetHTTPIdleTimeout(t int) {
	s.httpIdleTimeout = time.Duration(t) * time.Millisecond
}

// SetHttpQuitTimeout set http quit timeout
func (s *Server) SetHttpQuitTimeout(t int) {
	s.httpQuitTimeout = time.Millisecond * time.Duration(t)
}

// SetGatewayPattern set gateway pattern
func (s *Server) SetGatewayPattern(pattern string) {
	s.gatewayPattern = pattern
}

// SetMaxMsgSize set max msg size
func (s *Server) SetMaxMsgSize(size int) {
	s.maxMsgSize = size
}

// AddHTTPHandle add http handle
func (s *Server) AddHTTPHandle(pattern string, handler http.Handler) {
	s.httpServeMux.Handle(pattern, handler)
}

// AddHTTPHandleFunc add http handle func
func (s *Server) AddHTTPHandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	s.httpServeMux.HandleFunc(pattern, handler)
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
		s.errChan <- s.runGRPCServer()
	}()

	go func() {
		s.errChan <- s.runHTTPServer()
	}()

	return s.wait()
}

// runGRPCServer run gRPC server
func (s *Server) runGRPCServer() (err error) {
	var opts []grpc.ServerOption

	if len(s.gRPCMiddlewares) > 0 {
		opts = append(opts, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(s.gRPCMiddlewares...)))
	}

	if s.maxMsgSize > 0 {
		opts = append(opts, grpc.MaxRecvMsgSize(s.maxMsgSize), grpc.MaxSendMsgSize(s.maxMsgSize))
	}

	if s.maxConnectionIdle > 0 {
		opts = append(opts, grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionIdle: s.maxConnectionIdle}))
	}

	s.gRPCServer = grpc.NewServer(opts...)
	s.gRPCRegister(s.gRPCServer)

	lis, err := net.Listen("tcp", s.gRPCIp+":"+s.gRPCPort)
	if err != nil {
		log.Cli.Error("grpc listen failed!", "error", err)
		return
	}

	if err = s.registerConsul(); err != nil {
		log.Cli.Info(config.V.GetString("server.service_name") + " register consul failed!")
	}

	log.Cli.Info(config.V.GetString("server.service_name") + " register consul success!")

	log.Cli.Info("grpc server running!", "addr", s.gRPCIp+":"+s.gRPCPort)

	if err = s.gRPCServer.Serve(lis); err != nil {
		log.Cli.Error("grpc serve failed!", "error", err)
		return
	}

	return
}

// runHTTPServer run http server
func (s *Server) runHTTPServer() (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	runtimeMux := runtime.NewServeMux(s.httpServeMuxOption...)

	runtime.DefaultContextTimeout = s.gRPCHandlerTimeout

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if s.maxMsgSize > 0 {
		opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(s.maxMsgSize)), grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(s.maxMsgSize)))
	}

	if err = s.httpRegister(ctx, runtimeMux, s.gRPCIp+":"+s.gRPCPort, opts); err != nil {
		log.Cli.Error("http handler failed!", "error", err)
		return
	}

	// gateway handle
	s.httpServeMux.Handle(s.gatewayPattern, runtimeMux)

	s.httpServer = &http.Server{
		Addr:    s.httpIp + ":" + s.httpPort,
		Handler: s.httpServeMux,
	}

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

	// add middleware, 先添加的在最外层
	for i := len(s.httpMiddlewares) - 1; i >= 0; i-- {
		s.httpServer.Handler = s.httpMiddlewares[i](s.httpServer.Handler)
	}

	log.Cli.Info("http server running!", "addr", s.httpIp+":"+s.httpPort)

	s.httpInitDone <- struct{}{}

	if err = s.httpServer.ListenAndServe(); err != nil {
		log.Cli.Error("http listen stoped!")
		return
	}

	return
}

func (s *Server) registerConsul() (err error) {
	// 获取consul操作对象
	consulClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return err
	}

	// string port -> int port
	port, _ := strconv.Atoi(s.gRPCPort)

	// 注册服务信息
	registerService := api.AgentServiceRegistration{
		ID:      config.V.GetString("server.service_name") + "-" + s.gRPCIp + ":" + s.gRPCPort,
		Name:    config.V.GetString("server.service_name"),
		Tags:    []string{"grpc", "consul"},
		Port:    port,
		Address: s.gRPCIp,
		Check: &api.AgentServiceCheck{
			TCP:      s.gRPCIp + ":" + s.gRPCPort,
			Timeout:  "1s",
			Interval: "10s",
		},
	}

	return consulClient.Agent().ServiceRegister(&registerService)
}

func (s *Server) deregisterConsul() (err error) {
	// 获取consul操作对象
	consulClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return err
	}

	return consulClient.Agent().ServiceDeregister(config.V.GetString("server.service_name") + "-" + s.gRPCIp + ":" + s.gRPCPort)
}

func (s *Server) wait() (err error) {

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case sig := <-c:
		switch sig {
		case syscall.SIGTERM, syscall.SIGINT:
			<-s.httpInitDone

			ctx, cancel := context.WithTimeout(context.Background(), s.httpQuitTimeout)
			defer cancel()

			if err = s.httpServer.Shutdown(ctx); err != nil {
				return
			}

			s.gRPCServer.GracefulStop()

			if err = s.deregisterConsul(); err != nil {
				return
			}
		case syscall.SIGQUIT:
			<-s.httpInitDone

			if err = s.httpServer.Close(); err != nil {
				return
			}

			s.gRPCServer.Stop()

			if err = s.deregisterConsul(); err != nil {
				return
			}
		}

		log.Cli.Error(fmt.Sprintf("server closed got signal %s, shutdown success.", sig.String()))
		return
	case err = <-s.errChan:
		return
	}
}
