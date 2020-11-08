package main

import (
	"context"
	"math/rand"
	"os"
	"time"

	"service-skeleton/server"

	log "github.com/sirupsen/logrus"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	if err := server.NewRpcServer(ctx).Run(os.Args); err != nil {
		log.WithField("err", err).Fatal("启动服务失败")
	}
}
