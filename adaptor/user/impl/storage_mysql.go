package impl

import "github.com/zengqiang96/service-skeleton/adaptor/user"

type StorageMySQL struct{}

func (StorageMySQL) GetUserById(id string) user.MUser {
	panic("implement me!")
}
