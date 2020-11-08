package handler

import (
	"context"

	app "github.com/zengqiang96/service-skeleton/application"
	"github.com/zengqiang96/service-skeleton/proto"
)

type User struct {

}

func (u *User) GetUser(ctx context.Context, req *proto.GetUserReq, rsp *proto.GetUserRsp) error {
	mUser :=  app.GetUser(ctx, req.Id)
	rsp.Id = mUser.Id
	rsp.Name = mUser.Name
	return nil
}
