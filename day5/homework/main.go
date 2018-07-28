package main
import (
	"os"
	"fmt"
)

func showMenu() {
	fmt.Printf("please select:\n")
	fmt.Printf("1. 添加学生信息\n")
	fmt.Printf("2. 显示学生列表\n")
	fmt.Printf("3. 添加书籍\n")
	fmt.Printf("4. 显示书籍列表\n")
	fmt.Printf("5. 借书\n")
	fmt.Printf("6. 退出\n")
}

func getStudentInfo() *Student {
	var stu Student
	fmt.Printf("please input id:\n")
	fmt.Scanf("%s\n", &stu.ID)

	fmt.Printf("please input name:\n")
	fmt.Scanf("%s\n", &stu.Name)

	fmt.Printf("please input sex:\n")
	fmt.Scanf("%s\n", &stu.Sex)

	fmt.Printf("please input grade:\n")
	fmt.Scanf("%f\n", &stu.Grade)

	return &stu
}

func addStudent(studentMgr *StudentMgr) {
	stu := getStudentInfo()
	err := studentMgr.AddStudent(stu)
	if err != nil {
		fmt.Printf("add student failed, err:%v\n", err)
		return
	}
}


func getBook() *Book {
	var book Book
	fmt.Printf("please input book id:\n")
	fmt.Scanf("%s\n", &book.ID)

	fmt.Printf("please input name:\n")
	fmt.Scanf("%s\n", &book.Name)

	fmt.Printf("please input num:\n")
	fmt.Scanf("%d\n", &book.Num)

	fmt.Printf("please input author:\n")
	fmt.Scanf("%s\n", &book.Author)

	return &book
}

func addBook(bookMgr *BookMgr) {
	book := getBook()
	err := bookMgr.AddBook(book)
	if err != nil {
		fmt.Printf("add book failed, err:%v\n", err)
		return
	}
}

func showStudentList(studentMgr *StudentMgr) {
	studentList, err := studentMgr.GetStudentList()
	if err != nil {
		fmt.Printf("get student list failed, err:%v\n", err)
		return
	}

	for _, val := range studentList {
		fmt.Printf("id:%s\n", val.ID)
		fmt.Printf("name:%s\n", val.Name)
		fmt.Printf("sex:%s\n", val.Sex)
		
		for _, book := range val.GetBorrowList() {
			fmt.Printf("book ID:%s\n", book.ID)
			fmt.Printf("book Name:%s\n", book.Name)
		}
		fmt.Printf("-------------------------\n")
		fmt.Println("\n\n")
	}
}


func showBookList(bookMgr *BookMgr) {
	bookList := bookMgr.GetBookList()
	for _, val := range bookList {
		fmt.Printf("id:%s\n", val.ID)
		fmt.Printf("name:%s\n", val.Name)
		fmt.Printf("author:%s\n", val.Author)
		fmt.Printf("num:%d\n", val.Num)
		
		fmt.Printf("-------------------------\n")
		fmt.Println("\n\n")
	}
}


func borrowBook(bookMgr *BookMgr, studentMgr* StudentMgr) {
	fmt.Printf("please input student ID\n")
	var studentID string
	fmt.Scanf("%s\n", &studentID)

	student, err := studentMgr.GetStudentByID(studentID)
	if err != nil {
		fmt.Printf("get studetn failed, err:%v\n", err)
		return
	}

	fmt.Printf("please input book ID\n")
	var bookId string
	fmt.Scanf("%s\n", &bookId)

	err = bookMgr.Borrow(bookId, student)
	if err != nil {
		fmt.Printf("borrow book failed, err:%v\n", err)
		return
	}

	fmt.Printf("借书成功\n")

	fmt.Printf("借书列表\n\n")
	bookList := student.GetBorrowList()
	for _, book := range bookList {
		fmt.Printf("id:%s\n", book.ID)
		fmt.Printf("name:%s\n", book.Name)
		fmt.Printf("author:%s\n", book.Author)
		fmt.Printf("-------------------------\n")
		fmt.Println("\n\n")
	}
}

func main() {
	var studentMgr *StudentMgr = &StudentMgr{}
	var bookMgr *BookMgr = &BookMgr{}

	for {
		showMenu()
		var sel int
		fmt.Scanf("%d\n", &sel)
		switch sel {
		case 1:
			addStudent(studentMgr)
		case 2:
			showStudentList(studentMgr)
		case 3:
			addBook(bookMgr)
		case 4:
			showBookList(bookMgr)
		case 5:
			borrowBook(bookMgr, studentMgr)
		case 6:
			os.Exit(0)
		}
	}
}