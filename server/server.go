package server

import (
	"context"
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
		return s.server.Serve(s.config.Network, s.config.Address)
	}

	// rpcx server, 注册服务
	s.server = server.NewServer()
	s.server.Register(new(handler.User), "")
	return s
}

func initComponents(ctx context.Context, conf *config.Config) {

}
