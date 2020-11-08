package user

import "context"

type APIImpl struct {
}

func (api APIImpl) GetUserById(ctx context.Context, id string) MUser {
	panic("implement me")
}
