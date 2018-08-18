package main
import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_"github.com/go-sql-driver/mysql"
	"database/sql"
)
type User struct {
	Id int64 `db:"id"`
	Name string `db:"name"`
	Age int `db:"age"`
}
func QueryRow(Db *sqlx.DB) {
	id := 230900
	var user User
	err := Db.Get(&user,"select id, name, age from user where id=?", id)
	if err == sql.ErrNoRows {
		fmt.Printf("not record found\n")
		return
	}
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}

	fmt.Printf("get user succ, user:%#v\n", user)
}

func Query(Db *sqlx.DB) {
	var user []*User
	id := 0
	err := Db.Select(&user, "select id, name, age from user where id>?", id)
	if err != nil {
		return
	}

	fmt.Printf("user :%#v\n", user)
}


func Insert(Db *sqlx.DB) {
	username := "user01"
	age := 18
	
	result, err := Db.Exec("insert into user(name, age)values(?, ?)", username, age)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("last insert id failed, err:%v\n", err)
		return
	}

	affectRows, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("affectRows failed, err:%v\n", err)
		return
	}

	fmt.Printf("last insert id:%d affect rows:%d\n", id, affectRows)
}

func main() {
	dns := "root:123456@tcp(192.168.20.200:3306)/golang"
	Db, err := sqlx.Connect("mysql", dns)
	if err != nil {
		fmt.Printf("open mysql failed, err:%v\n", err)
		return
	}

	err = Db.Ping()
	if err != nil {
		fmt.Printf("ping failed, err:%v\n", err)
		return
	}

	fmt.Printf("connect to db succ\n")
	//QueryRow(Db)
	//Query(Db)
	Insert(Db)
	
	//Transaction(Db)
}