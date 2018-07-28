package main

import (
	"errors"
)

type Student struct {
	ID string
	Name string
	Sex string
	Grade string
	BorrowsBook []*Book
}

func (s *Student)AddBook (book *Book) (err error) {
	if book == nil {
		err = errors.New("parameter error")
		return
	}

	s.BorrowsBook = append(s.BorrowsBook, book)
	return
}

func (s *Student)GetBorrowList() (bookList []*Book) {
	return s.BorrowsBook
}