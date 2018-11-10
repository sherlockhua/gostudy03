package data_layer

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gostudy03/library_mgr/model"
)

func InsertBook(book *model.Book) (err error) {
	if book == nil {
		err = fmt.Errorf("invalid book parameter")
		return
	}
// \r\n
	sqlstr := "select book_id from book where book_id=?"
	var bookId string
	err = Db.Get(&bookId, sqlstr, book.BookId)
	if err == sql.ErrNoRows {
		//插入操作
		sqlstr = `insert into book(
					  author, book_name, publish_time, stock_num, book_id
				  )values(?, ?, ?, ?, ?)`
		_, err = Db.Exec(sqlstr, book.Author, book.BookName, book.PublishTime, book.StockNum, book.BookId)
		if err != nil {
			return
		}
		return
	}

	if err != nil {
		return
	}

	err = fmt.Errorf("book_id:%s is already exists", bookId)
	return
}

func UpdateBook(book *model.Book) (err error) {
	if book == nil {
		err = fmt.Errorf("invalid book parameter")
		return
	}

	//插入操作
	sqlstr := `update book set
						author = ?, book_name=?, publish_time=?,  stock_num= stock_num+?
				  where 
				        book_id = ?`
	result, err := Db.Exec(sqlstr, book.Author, book.BookName, book.PublishTime, book.StockNum, book.BookId)
	if err != nil {
		return
	}

	affects, err := result.RowsAffected()
	if err != nil {
		return
	}

	if affects == 0 {
		err = fmt.Errorf("update book failed, book_id:%s, not found", book.BookId)
		return
	}
	return
}

func QueryBook(bookId string) (book *model.Book, err error) {
	//插入操作
	sqlstr := `select 
					author , book_name, publish_time,  stock_num, book_id, id
				from book
				where 
					book_id = ?`
	book = &model.Book{}
	err = Db.Get(book, sqlstr, bookId)
	if err != nil {
		return
	}
	return
}
