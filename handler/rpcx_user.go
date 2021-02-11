package handler

import (
	"context"

	"github.com/zengqiang96/service-skeleton/proto"
	"github.com/zengqiang96/service-skeleton/service"
)

type User struct {
}

func (u *User) GetUser(ctx context.Context, req *proto.GetUserReq, rsp *proto.GetUserRsp) error {
	mUser := service.GetUser(ctx, req.Id)
	rsp.Id = mUser.Id
	rsp.Name = mUser.Name
	return nil
}
