package user

import (
	"database/sql"
	"github.com/zengqiang96/service-skeleton/component/user"
)


type MysqlUser struct {
	*sql.DB
}

func (u MysqlUser) GetUserById(id string) user.MUser {
	panic("implement me")
}
