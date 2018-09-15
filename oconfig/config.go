package oconfig

import (
	"strings"
	"io/ioutil"
	"reflect"
	"fmt"
)

func UnMarshal(data []byte, result interface{}) (err error) {

	t := reflect.TypeOf(result)
	v := reflect.ValueOf(result)

	_ = v
	kind := t.Kind()
	if kind != reflect.Ptr {
		panic("please pass a address")
	}



	lines := strings.Split(string(data), "\n")
	lineNo := 0
	for _, line := range lines {
		lineNo++
		line = strings.Trim(line, " \t\r\n")
		if len(line) == 0 {
			continue
		}

		if line[0] == '#' || line[0] == ';' {
			continue
		}

		//fmt.Printf("line:%s\n", line)
		if line[0] == '[' {
			//解析section/group
			if len(line) <= 2  || line[len(line)-1] != ']' {
				tips := fmt.Sprintf("syntax error, invalid section:\"%s\" line:%d", line, lineNo)
				panic(tips)
			}

			sectionName := strings.TrimSpace(line[1:len(line)-1])
			if len(sectionName) == 0 {
				tips := fmt.Sprintf("syntax error, invalid section:\"%s\" line:%d", line, lineNo)
				panic(tips)
			}

			fmt.Printf("section:%s\n", sectionName)
			for i := 0; i <t.Elem().NumField(); i++ {
				//field := v.Field(i)
				field := t.Elem().Field(i)
				if field.Tag.Get("ini") == sectionName {
					fmt.Printf("found field by group name:%s\n", sectionName)
				}
			}
		} else {

		}


	}
	return
}

func UnMarshalFile(filename string, result interface{}) (err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	return UnMarshal(data, result)
}