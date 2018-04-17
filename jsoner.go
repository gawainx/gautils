package gautils

import (
	"encoding/json"
	"reflect"
)

type JsonMap map[string]interface{}

func JsonParser(src string) *JsonMap{
	//src: json-format string
	//return: JsonMap pointer
	//Unmarshal json-string to map[string]interface{}
	var res JsonMap
	err := json.Unmarshal([]byte(src),&res)
	if err != nil{
		return nil
	}else{
		return &res
	}
}

func GetJsonField(src string,key string) (interface{},reflect.Type){
	//return value and it's type
	var tmp JsonMap
	err := json.Unmarshal([]byte(src),&tmp)
	if err != nil{
		return nil, reflect.TypeOf(nil)
	}else{
		res := tmp[key]
		return res, reflect.TypeOf(res)
	}
}



