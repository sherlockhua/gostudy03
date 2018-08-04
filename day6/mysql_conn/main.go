package main


import (
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

type User struct {
	Id int64 `db:"id"`
	Name string `db:"name"`
	Age int `db:"age"`
}

func QueryRow(Db *sql.DB) {
	id := 23
	row := Db.QueryRow("select id, name, age from user where id=?", id)

	var user User
	err := row.Scan(&user.Id, &user.Name, &user.Age)
	if err == sql.ErrNoRows {
		fmt.Printf("not found data of id:%d\n", id)
		return
	}

	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}

	fmt.Printf("user:%#v\n", user)
}


func Query(Db *sql.DB) {
	id := 0
	rows, err := Db.Query("select id, name, age from user where id>?", id)
	//这个rows结果集，用完之后，一定释放！rows.Close()
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err == sql.ErrNoRows {
		fmt.Printf("not found data of id:%d\n", id)
		return
	}
	
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
	
		fmt.Printf("user:%#v\n", user)
	}
	
}

func Insert(DB *sql.DB) {
	username := "user01"
	age := 18
	
	result, err := DB.Exec("insert into user(name, age)values(?, ?)", username, age)
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

func Update(DB *sql.DB) {
	username := "user02"
	age := 108
	id := 3
	result, err := DB.Exec("update user set name=?, age=? where id=?", username, age, id)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}

	affectRows, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("affectRows failed, err:%v\n", err)
		return
	}

	fmt.Printf("last insert id:%d affect rows:%d\n", id, affectRows)
}

func Delete(DB *sql.DB) {
	id := 3
	result, err := DB.Exec("delete from user where id=?", id)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}

	affectRows, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("affectRows failed, err:%v\n", err)
		return
	}

	fmt.Printf("last insert id:%d affect rows:%d\n", id, affectRows)
}

func PrepareQuery(DB* sql.DB) {
	 stmt, err := DB.Prepare("select id, name, age from user where id>?")
	 if err != nil {
		 fmt.Printf("prepare failed, err:%v\n", err)
		 return
	 }
	 id := 1
	 rows, err := stmt.Query(id)
	 //这个rows结果集，用完之后，一定释放！rows.Close()
	defer func() {
		if rows != nil {
			rows.Close()
		}
		if stmt != nil {
			stmt.Close()
		}
	}()
	if err == sql.ErrNoRows {
		fmt.Printf("not found data of id:%d\n", id)
		return
	}
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
	
		fmt.Printf("user:%#v\n", user)
	}
}

func PrepareInsert(DB* sql.DB) {
	stmt, err := DB.Prepare("insert into user(name, age)values(?, ?)")
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}

	result, err := stmt.Exec("user10111", 108)
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

func Transaction(Db *sql.DB) {

	tx, err := Db.Begin()
	if err != nil {
		fmt.Printf("begin failed, err:%v\n", err)
		return
	}

	_, err = tx.Exec("insert into user(name, age)values(?, ?)", "user0101", 108)
	if err != nil {
		tx.Rollback()
		return
	}

	_, err = tx.Exec("update user set name=?, age=?", "user0101", 108)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return
	}
}

func main() {
	dns := "root:123456@tcp(192.168.20.200:3306)/golang"
	Db, err := sql.Open("mysql", dns)
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
	//Insert(Db)
	//Update(Db)
	//Delete(Db)
	//PrepareQuery(Db)
	//PrepareInsert(Db)
	Transaction(Db)
}