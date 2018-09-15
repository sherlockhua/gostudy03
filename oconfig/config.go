package oconfig

import (
	"strconv"
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


	var sectionName string
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

			sectionName = strings.TrimSpace(line[1:len(line)-1])
			if len(sectionName) == 0 {
				tips := fmt.Sprintf("syntax error, invalid section:\"%s\" line:%d", line, lineNo)
				panic(tips)
			}

			fmt.Printf("section:%s\n", sectionName)
		} else {
			if len(sectionName) == 0 {
				tips := fmt.Sprintf("syntax error, key-value:%s 不属于任何section， lineNo:%d", line, lineNo)
				panic(tips)
			}

			index := strings.Index(line, "=")
			if index == -1 {
				tips := fmt.Sprintf("syntax error, not found =, line:%s, lineNo:%d", line, lineNo)
				panic(tips)
			}

			key := strings.TrimSpace(line[0:index])
			value := strings.TrimSpace(line[index+1:])

			if len(key)  == 0 {
				tips := fmt.Sprintf("syntax error, not found =, line:%s, lineNo:%d", line, lineNo)
				panic(tips)
			}

			//1. 找到sectionName在result中对应的结构体s1
			for i := 0; i <t.Elem().NumField(); i++ {
				//field := v.Field(i)
				tfield := t.Elem().Field(i)
				vField := v.Elem().Field(i)
				if tfield.Tag.Get("ini") != sectionName {
					continue
				}

				//2. 通过key找到对应结构体s1中的对应字段
				tfieldType := tfield.Type
				if tfieldType.Kind() != reflect.Struct {
					tips := fmt.Sprintf("field %s is not struct", tfieldType.Name)
					panic(tips)
				}

				for j := 0; j < tfieldType.NumField(); j++ {
					tKeyField := tfieldType.Field(j)
					vKeyField := vField.Field(j)
					if tKeyField.Tag.Get("ini") != key {
						continue
					}

					//找到了子结构体中的字段
					switch tKeyField.Type.Kind() {
					case reflect.String:
						vKeyField.SetString(value)
					case reflect.Int, reflect.Uint, reflect.Int16, reflect.Uint16:
						fallthrough
					case reflect.Int64, reflect.Uint64, reflect.Int32, reflect.Uint32:
						valueInt, err := strconv.ParseInt(value, 10, 64)
						if err != nil {
							tips := fmt.Sprintf("value:%s can not convert to int64, lineNo:%d", value, lineNo)
							panic(tips)
						}
						vKeyField.SetInt(valueInt)
					case reflect.Float32, reflect.Float64:
						valueFloat, err := strconv.ParseFloat(value, 64)
						if err != nil {
							tips := fmt.Sprintf("value:%s can not convert to float64, lineNo:%d", value, lineNo)
							panic(tips)
						}
						vKeyField.SetFloat(valueFloat)
					default:
						tips := fmt.Sprintf("key:\"%s\" can not convert to %v, lineNo:%d", 
							key, tKeyField.Type.Kind(), lineNo)
						panic(tips)
					}
				}
				break
			}
			
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

func Marshal(result interface{})(data []byte, err error) {

	t := reflect.TypeOf(result)
	v := reflect.ValueOf(result)

	if t.Kind() != reflect.Struct {
	    panic("please input struct type")
		return
	}

	var strSlice []string
	for i := 0; i < t.NumField(); i++ {
		tField := t.Field(i)
		vFiled := v.Field(i)
		if tField.Type.Kind() != reflect.Struct {
			continue
		}

		sectionName := tField.Name
		if len(tField.Tag.Get("ini")) > 0  {
			sectionName = tField.Tag.Get("ini")
		}

		sectionName = fmt.Sprintf("[%s]\n", sectionName)
		strSlice = append(strSlice, sectionName)
		for j := 0; j < tField.Type.NumField(); j++ {
			//1. 拿到类型信息
			subTField := tField.Type.Field(j)
			if subTField.Type.Kind() == reflect.Struct || subTField.Type.Kind() == reflect.Ptr{
				//跳过结构体字段
				continue
			}

			subTFieldName := subTField.Name
			subTFieldTagName := subTField.Tag.Get("ini")
			if len(subTFieldTagName) > 0 {
				subTFieldName = subTFieldTagName
			}

			//2. 拿到值信息
			subVField := vFiled.Field(j)
			fieldStr := fmt.Sprintf("%s=%v\n", subTFieldName, subVField.Interface())
			fmt.Printf("conf:%s\n", fieldStr)

			strSlice = append(strSlice, fieldStr)
		}
	}

	for _, v := range strSlice {
		data = append(data, []byte(v)...)
	}
	return
}

func MarshalFile(filename string, result interface{}) (err error) {

	data, err := Marshal(result)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(filename, data, 0755)
	return
}