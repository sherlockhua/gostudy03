package dal

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gostudy03/xlog"
	"github.com/jmoiron/sqlx"
)

var (
	Db *sqlx.DB
)

func InitDb(dns string) (err error) {

	Db, err = sqlx.Connect("mysql", dns)
	if err != nil {
		xlog.LogError("open mysql failed, err:%v", err)
		return
	}

	err = Db.Ping()
	if err != nil {
		xlog.LogError("ping failed, err:%v", err)
		return
	}

	return
}
