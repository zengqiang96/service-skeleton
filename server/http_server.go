package server

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"github.com/zengqiang96/service-skeleton/config"
	"github.com/zengqiang96/service-skeleton/handler"
)

type HttpServer struct {
	*cli.App
	config *config.Config
	server *gin.Engine
}

func NewHttpServer(ctx context.Context) *HttpServer {
	s := &HttpServer{App: cli.NewApp()}
	s.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "c",
			Usage: "配置文件",
		},
	}

	s.Action = func(c *cli.Context) error {
		var err error
		if c.String("c") == "" {
			return fmt.Errorf("usage: service-skeleton -c 配置文件路径")
		}
		log.Printf("开始读取配置文件: %s", c.String("c"))

		s.config, err = config.Load(c.String("c"))
		if err != nil {
			return errors.Wrap(err, "加载配置文件失败")
		}

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

func routes(engine *gin.Engine) error {
	engine.GET("/user", handler.GetUser)

	return nil
}
