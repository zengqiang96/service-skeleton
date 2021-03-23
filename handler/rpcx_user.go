package handler

import (
	"context"

	"github.com/zengqiang96/service-skeleton/proto"
)

type UserHandler struct{}

func (UserHandler) RpcxGetUser(ctx context.Context, req *proto.GetUserReq, rsp *proto.GetUserRsp) error {
	rsp.Id = ""
	rsp.Name = ""
	return nil
}
