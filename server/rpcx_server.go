package server

import (
	"context"
	"fmt"

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

	s := &RpcxServer{App: cli.NewApp()}
	s.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "c",
			Usage: "配置文件",
		},
	}

	s.Action = func(c *cli.Context) error {
		var err error
		s.config, err = config.Load("")
		if err != nil {
			return fmt.Errorf("加载配置文件失败: %w", err)
		}
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
	err = server.Register(new(handler.UserHandler), "")
	return err
}
