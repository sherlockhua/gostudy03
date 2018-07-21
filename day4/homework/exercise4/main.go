package main
import (
	"os"
	"fmt"
)

type Student struct {
	Id string
	Name string
	Age int
	Sex string
	Score float32
}

func showMenu() {
	fmt.Printf("please select:\n")
	fmt.Printf("1. 添加学生信息\n")
	fmt.Printf("2. 修改学生信息\n")
	fmt.Printf("3. 显示学生列表\n")
	fmt.Printf("4. 退出\n")
}

func getStudentInfo() Student {
	var stu Student
	fmt.Printf("please input id:\n")
	fmt.Scanf("%s\n", &stu.Id)

	fmt.Printf("please input name:\n")
	fmt.Scanf("%s\n", &stu.Name)

	fmt.Printf("please input age:\n")
	fmt.Scanf("%d\n", &stu.Age)

	fmt.Printf("please input sex:\n")
	fmt.Scanf("%s\n", &stu.Sex)

	fmt.Printf("please input score:\n")
	fmt.Scanf("%f\n", &stu.Score)

	return stu
}

func addStudent(allStudent map[string]Student) {
	stu := getStudentInfo()
	_, ok := allStudent[stu.Id]
	if ok {
		fmt.Printf("student %s is exists\n", stu.Id)
		return
	}

	allStudent[stu.Id] = stu
}

func modifyStudent(allStudent map[string]Student) {
	stu := getStudentInfo()
	_, ok := allStudent[stu.Id]
	if !ok {
		fmt.Printf("student %s is not exists\n", stu.Id)
		return
	}

	allStudent[stu.Id] = stu
}

func showStudentList(allStudent map[string]Student) {
	for _, val := range allStudent {
		fmt.Printf("id:%s\n", val.Id)
		fmt.Printf("name:%s\n", val.Name)
		fmt.Printf("sex:%s\n", val.Sex)
		fmt.Printf("age:%d\n", val.Age)
		fmt.Printf("score:%f\n", val.Score)

		fmt.Printf("-------------------------\n")
		fmt.Println("\n\n")
	}
}

func main() {
	var allStudent map[string]Student = make(map[string]Student, 100)
	for {
		showMenu()
		var sel int
		fmt.Scanf("%d\n", &sel)
		switch sel {
		case 1:
			addStudent(allStudent)
		case 2:
			modifyStudent(allStudent)
		case 3:
			showStudentList(allStudent)
		case 4:
			os.Exit(0)
		}
	}
}