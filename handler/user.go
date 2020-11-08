package handler

import (
	"context"
	"github.com/zengqiang96/service-skeleton/proto"
)

type User struct {
}

func (u *User) GetUser(ctx context.Context, req *proto.GetUserReq, rsp *proto.GetUserRsp) error {
	return nil
}
