package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zengqiang96/service-skeleton/proto"
	"github.com/zengqiang96/service-skeleton/service"
)

func GetUser(ctx *gin.Context) {
	var (
		getUserReq proto.GetUserReq
		err        error
	)

	if err = ctx.ShouldBindQuery(&getUserReq); err != nil {
		ctx.JSON(http.StatusBadRequest, "")
		return
	}

	mUser := service.GetUser(ctx, getUserReq.Id)
	ctx.JSON(http.StatusOK, mUser)
	return
}
