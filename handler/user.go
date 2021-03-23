package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	ctx.JSON(http.StatusOK, gin.H{})
	return
}
