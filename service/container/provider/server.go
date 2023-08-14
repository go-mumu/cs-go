/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-10
 * File: server.go
 * Desc: server provider
 */

package provider

import (
	"github.com/go-mumu/cs-go/library/common/flags"
	"github.com/go-mumu/cs-go/library/config"
	"github.com/go-mumu/cs-go/library/log"
	"github.com/go-mumu/cs-go/proto/pb"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hashicorp/consul/api"
	"golang.org/x/net/context"
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

var defaultHttpQuitTimeout = 5 * time.Second

type HttpRegister func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

func HttpRegisterFunc(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption, registerFunc ...HttpRegister) error {
	for _, f := range registerFunc {
		err := f(ctx, mux, endpoint, opts)
		if err != nil {
			return err
		}
	}

	return nil
}

type Server struct {
	GrpcServer *grpc.Server
	// GrpcAddr        string
	grpcRegister    func(server *grpc.Server)
	grpcMiddlewares []grpc.UnaryServerInterceptor

	serverIp   string
	serverPort int
	serverId   string

	httpServer    *http.Server
	httpServerMux *http.ServeMux
	httpRegister  HttpRegister
	// httpAddr        string
	httpIp   string
	httpPort int

	httpMiddlewares []func(h http.Handler) http.Handler
	httpInitDone    chan struct{}
	httpOptions     []runtime.ServeMuxOption

	gatewayPattern string
	maxMsgSize     int

	maxConnectionIdle time.Duration
	httpIdleTimeout   time.Duration
	//
	defaultContextTimeout time.Duration
	// httpHandlerTimeout time.Duration
	//
	httpReadTimeout  time.Duration
	httpWriteTimeout time.Duration
	// httpQuitTimeout  time.Duration
	//
	errChan chan error
}

func NewServer(h *Handler) *Server {
	return &Server{
		maxMsgSize:        config.V.GetInt("rpc.max_msg_size_byte"),
		maxConnectionIdle: time.Duration(config.V.GetInt("rpc.max_connection_idle_ms")) * time.Millisecond,
		grpcRegister: func(s *grpc.Server) {
			pb.RegisterUserServiceServer(s, h.UserServiceHandler)
		},
		grpcMiddlewares: []grpc.UnaryServerInterceptor{},
		errChan:         make(chan error),

		defaultContextTimeout: time.Duration(config.V.GetInt("rpc.default_context_timeout_ms")) * time.Millisecond,

		serverIp:   flags.NodeIp,
		serverPort: flags.NodePort,
		serverId:   flags.NodeId,

		httpOptions:   []runtime.ServeMuxOption{},
		httpServerMux: http.NewServeMux(),
		httpIp:        flags.NodeIp,
		httpPort:      flags.HttpPort,

		httpInitDone:    make(chan struct{}, 1),
		httpMiddlewares: []func(h http.Handler) http.Handler{},
		gatewayPattern:  "/",

		httpReadTimeout:  time.Duration(config.V.GetInt("rpc.http_read_timeout_ms")) * time.Millisecond,
		httpWriteTimeout: time.Duration(config.V.GetInt("rpc.http_write_time_ms")) * time.Millisecond,

		httpIdleTimeout: time.Duration(config.V.GetInt("rpc.http_idle_timeout_ms")) * time.Millisecond,

		httpRegister: func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
			return HttpRegisterFunc(ctx, mux, endpoint, opts,
				[]HttpRegister{
					pb.RegisterUserServiceHandlerFromEndpoint,
				}...,
			)
		},
		/*GrpcAddr:        string

		httpServer:      *http.Server
		httpAddr:        string
		httpRegister:    HttpRegister

		httpInitDone:    chan struct{}


		gatewayPattern: string

		httpIdleTimeout: time.Duration

		grpcHandlerTimeout: time.Duration
		httpHandlerTimeout: time.Duration


		httpQuitTimeout:  time.Duration

		*/
	}
}

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

	var opts []grpc.ServerOption

	if len(s.grpcMiddlewares) > 0 {
		grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(s.grpcMiddlewares...))
	}

	if s.maxMsgSize > 0 {
		opts = append(opts, grpc.MaxRecvMsgSize(s.maxMsgSize), grpc.MaxSendMsgSize(s.maxMsgSize))
	}

	if s.maxConnectionIdle > 0 {
		opts = append(opts, grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionIdle: s.maxConnectionIdle}))
	}

	s.GrpcServer = grpc.NewServer(opts...)
	s.grpcRegister(s.GrpcServer)

	if err := s.registerConsul(); err != nil {
		return err
	} else {
		log.Cli.Info("register consul success.")
	}

	lis, err := net.Listen("tcp", s.serverIp+":"+strconv.FormatInt(int64(s.serverPort), 10))
	if err != nil {
		return err
	}

	log.Cli.Info("grpc server running!", "addr", s.serverIp+":"+strconv.FormatInt(int64(s.serverPort), 10))

	if err = s.GrpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (s *Server) registerConsul() error {
	// 获取consul操作对象
	consulClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return err
	}

	// 注册服务信息
	registerService := api.AgentServiceRegistration{
		ID:      config.V.GetString("service.service_name") + "-" + s.serverId,
		Name:    config.V.GetString("service.service_name"),
		Tags:    []string{"grpc", "consul"},
		Port:    s.serverPort,
		Address: s.serverIp,
		Check: &api.AgentServiceCheck{
			TCP:      s.serverIp + ":" + strconv.FormatInt(int64(s.serverPort), 10),
			Timeout:  "5s",
			Interval: "5s",
		},
	}

	if err = consulClient.Agent().ServiceRegister(&registerService); err != nil {
		return err
	}

	return nil
}

// runHTTPServer 启动http服务
func (s *Server) runHttpServer() error {
	ctx, _ := context.WithCancel(context.Background())

	runtimeMux := runtime.NewServeMux(s.httpOptions...)

	runtime.DefaultContextTimeout = s.defaultContextTimeout

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if s.maxMsgSize > 0 {
		opts = append(
			opts,
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(s.maxMsgSize)),
			grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(s.maxMsgSize)),
		)
	}

	if err := s.httpRegister(ctx, runtimeMux, ":"+strconv.Itoa(s.serverPort), opts); err != nil {
		return err
	}

	// gateway handle
	s.httpServerMux.Handle(s.gatewayPattern, runtimeMux)

	s.httpServer = &http.Server{
		Addr:    ":" + strconv.Itoa(s.httpPort),
		Handler: s.httpServerMux,
	}

	// set Timeout
	{
		if s.defaultContextTimeout != 0 {
			s.httpServer.Handler = http.TimeoutHandler(s.httpServer.Handler, s.defaultContextTimeout, "503 Handler timeout")
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

	log.Cli.Info("http server running!", "addr", ":"+strconv.Itoa(s.httpPort))

	s.httpInitDone <- struct{}{}

	if err := s.httpServer.ListenAndServe(); err != nil {
		return err
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
			ctx, _ := context.WithTimeout(context.Background(), defaultHttpQuitTimeout)
			err = s.httpServer.Shutdown(ctx)
			if err != nil {
				return err
			}
			s.GrpcServer.GracefulStop()
		case syscall.SIGQUIT:
			<-s.httpInitDone
			err = s.httpServer.Close()
			if err != nil {
				return err
			}
			s.GrpcServer.Stop()
		}

		log.Cli.Error("server closed got signal, shutdown success", "signal", sig.String())
		return err
	case err = <-s.errChan:
		return err
	}
}
