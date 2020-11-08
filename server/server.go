package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/smallnest/rpcx/server"
	"github.com/urfave/cli"
	"github.com/zengqiang96/service-skeleton/config"
	"github.com/zengqiang96/service-skeleton/handler"
)

type RpcxServer struct {
	*cli.App
	config *config.Config
	server *server.Server
}

type HttpServer struct {
	*cli.App
	config *config.Config
	server *gin.Engine
}

func NewHttpServer(ctx context.Context) *HttpServer {
	s := new(HttpServer)
	s.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "c",
			Usage: "配置文件",
		},
	}

	s.Action = func(c *cli.Context) error {
		initComponents(ctx, s.config)
		// http server 注册路由
		if err := routes(s.server); err != nil {
			return err
		}
		return nil
	}
	s.server = gin.Default()

	return s
}

func NewRpcServer(ctx context.Context) *RpcxServer {
	s := &RpcxServer{}
	s.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "c",
			Usage: "配置文件",
		},
	}

	s.Action = func(c *cli.Context) error {
		initComponents(ctx, s.config)
		// rpcx server, 注册服务
		if err := register(s.server); err != nil {
			return err
		}
		return s.server.Serve(s.config.Network, s.config.Address)
	}
	s.server = server.NewServer()
	return s
}

func register(server *server.Server) (err error) {
	err = server.Register(new(handler.User), "")
	return err
}

func routes(engine *gin.Engine) error {
	engine.GET("/user", handler.GetUser)

	return nil
}

func initComponents(ctx context.Context, conf *config.Config) {


}
