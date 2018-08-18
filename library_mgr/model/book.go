package model

import (
	"time"
)


type Book struct {
	Id int64 `db:"id"`
	Author string `db:"author"`
	BookName string 	`db:"book_name"`
	PublishTime time.Time `db:"publish_time"` 
	StockNum uint `db:"stock_num"`
	BookId  string `db:"book_id"`
}