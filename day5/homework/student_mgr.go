package main

import (
	"errors"
	"fmt"
)

type StudentMgr struct {
	StudentList []*Student
}

func (s*StudentMgr)AddStudent(student *Student) (err error) {
	if student == nil {
		err = errors.New("invalid parameter, student is nil")
		return
	}

	for _, val := range s.StudentList {
		if val.ID == student.ID {
			err  = fmt.Errorf("student ID:%s is exists", val.ID)
			return
		}
	}

	s.StudentList  = append(s.StudentList, student)
	return
}


func (s*StudentMgr)QueryStudentByBookID(bookID string) (studentList []*Student, err error) {
	
	for _, val := range s.StudentList {
		for _, book := range val.BorrowsBook {
			if book.ID == bookID {
				studentList = append(studentList, val)
				continue
			}
		}
	}

	return
}


func (s*StudentMgr)GetStudentByID(id string) (student *Student, err error) {
	
	for _, val := range s.StudentList {
		if val.ID == id {
			student = val
			return
		}
	}

	err = fmt.Errorf("not found student by id:%s", id)
	return
}


func (s*StudentMgr)GetStudentList() (student[]*Student, err error) {
	return s.StudentList, nil
}