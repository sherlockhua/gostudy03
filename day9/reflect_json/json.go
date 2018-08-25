package reflect_json

import (
	"reflect"
	"fmt"
)


func Marshal(data interface{}) (jsonStr string) {
	
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	switch t.Kind(){
	case reflect.String,reflect.Int,reflect.Int32:
		jsonStr = fmt.Sprintf("\"%v\"", data)
	case reflect.Struct:
		numField := t.NumField()
		for i := 0; i < numField; i++ {
			//类型信息
			name := t.Field(i).Name
			tag := t.Field(i).Tag.Get("json")
			if len(tag) > 0 {
				name = tag
			}
			//值信息
			vField := v.Field(i)
			vFieldValue := vField.Interface()
			//拼接json
			if t.Field(i).Type.Kind()  == reflect.String {
				jsonStr += fmt.Sprintf("\"%s\":\"%v\"", name, vFieldValue)
			} else {
				jsonStr += fmt.Sprintf("\"%s\":%v", name, vFieldValue)
			}

			if i != numField - 1 {
				jsonStr += ","
			}
		}

		jsonStr = "{" +jsonStr + "}"
	}
	return
}