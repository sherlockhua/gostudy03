package main

import (
	"errors"
	"fmt"
)

type BookMgr struct {
	BookList []*Book
}

func (b *BookMgr) AddBook(book *Book) (err error) {
	if book == nil {
		err = errors.New("invalid parameter, book is nil")
		return
	}

	for _, val := range b.BookList {
		if val.ID == book.ID {
			val.Num += book.Num
			return
		}
	}

	b.BookList = append(b.BookList, book)
	return
}


func (b *BookMgr) FindByID(ID string) (book *Book, err error) {
	
	for _, val := range b.BookList {
		if val.ID == ID {
			book = val
			return
		}
	}

	err = fmt.Errorf("not found book id:%s", ID)
	return
}

func (b *BookMgr) FindByName(name string) (book []*Book, err error) {
	
	for _, val := range b.BookList {
		if val.Name == name {
			book = append(book, val)
			continue
		}
	}

	return
}

func (b *BookMgr) FindByAuthor(author string) (book []*Book, err error) {
	
	for _, val := range b.BookList {
		if val.Author == author {
			book = append(book, val)
			continue
		}
	}

	return
}

func (b *BookMgr) Borrow(bookID string, student *Student) (err error) {
	book, err := b.FindByID(bookID)
	if err != nil {
		return
	}

	if book.Num <= 0 {
		err = fmt.Errorf("book ID:%s have not copys", bookID)
		return
	}

	book.Num--
	//student.BorrowsBook = append(student.BorrowsBook, book)
	err = student.AddBook(book)
	return
}

func (b *BookMgr) GetBookList() (bookList []*Book) {
	bookList = b.BookList
	return
}