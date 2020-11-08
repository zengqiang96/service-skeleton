package user

import "context"

type IAPI interface {
	GetUserById(ctx context.Context, id string) MUser
}

var _APIInstance IAPI

func API() IAPI {
	return _APIInstance
}

// 单元测试时，实现fake api 替换掉默认api实例
func SetAPI(api IAPI) {
	_APIInstance = api
}

func init() {
	_APIInstance = &APIImpl{}
}
