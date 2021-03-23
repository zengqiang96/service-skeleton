package adaptor

import "database/sql"

var DB sql.DB

func InitMySQL() {
	DB = sql.DB{}
}
