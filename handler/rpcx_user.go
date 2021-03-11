package handler

import (
	"context"

	"github.com/zengqiang96/service-skeleton/component/user"
	"github.com/zengqiang96/service-skeleton/proto"
)

type User struct {
}

func (u *User) GetUser(ctx context.Context, req *proto.GetUserReq, rsp *proto.GetUserRsp) error {
	mUser := user.GetUserById(ctx, req.Id)
	rsp.Id = mUser.Id
	rsp.Name = mUser.Name
	return nil
}
