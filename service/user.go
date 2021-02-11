package service

import (
	"context"

	"github.com/zengqiang96/service-skeleton/component/user"
)

func GetUser(ctx context.Context, userId string) user.MUser {
	return user.API().GetUserById(ctx, userId)
}
