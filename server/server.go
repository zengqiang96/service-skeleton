package server

import (
	"context"

	"github.com/zengqiang96/service-skeleton/adaptor"
	"github.com/zengqiang96/service-skeleton/config"
)

// 初始化数据库组件，远程调用等的client
func initComponents(ctx context.Context, conf *config.Config) {
	adaptor.InitMySQL()
}
