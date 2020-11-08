package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	app "github.com/zengqiang96/service-skeleton/application"
	"github.com/zengqiang96/service-skeleton/proto"
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

	mUser := app.GetUser(ctx, getUserReq.Id)
	ctx.JSON(http.StatusOK, mUser)
	return
}
