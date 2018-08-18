package data_layer

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	Db *sqlx.DB
)

func Init(dns string) (err error) {

	Db, err = sqlx.Connect("mysql", dns)
	if err != nil {
		fmt.Printf("open mysql failed, err:%v\n", err)
		return
	}

	err = Db.Ping()
	if err != nil {
		fmt.Printf("ping failed, err:%v\n", err)
		return
	}

	return
}
