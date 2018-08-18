package data_layer

import (
	"github.com/gostudy03/library_mgr/model"
	"testing"
	"time"
)

func init() {
	dns := "root:123456@tcp(192.168.20.200:3306)/library_mgr?parseTime=True"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestInsertBook(t *testing.T) {
	var book model.Book
	book.Author = "jim"
	book.BookId = "83883488344"
	book.BookName = "C语言从入门到放弃"
	book.PublishTime = time.Now()
	book.StockNum = 10

	err := InsertBook(&book)
	if err != nil {
		t.Errorf("insert book failed, err:%v", err)
		return
	}

	t.Logf("insert book succ")
}

func TestUpdateBook(t *testing.T) {
	var book model.Book
	book.Author = "jim"
	book.BookId = "83883488344"
	book.BookName = "C语言从入门到放弃"
	book.PublishTime = time.Now()
	book.StockNum = 10

	err := UpdateBook(&book)
	if err != nil {
		t.Errorf("update book failed, err:%v", err)
		return
	}

	t.Logf("update book succ")
}


func TestQueryBook(t *testing.T) {
	var book *model.Book
	bookId := "83883488344"
	book, err := QueryBook(bookId)
	if err != nil {
		t.Errorf("query book failed, err:%v", err)
		return
	}

	t.Logf("query book succ, book:%#v", book)
}
