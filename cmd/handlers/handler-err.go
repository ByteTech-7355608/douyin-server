package handlers

import (
	"ByteTech-7355608/douyin-server/pkg/constants"
	"reflect"
)

func HandlerErr(response interface{}, err error) {
	e := reflect.ValueOf(response)
	var code = int32(200)
	var msg = "操作成功"
	if err != nil {
		if status, ok := err.(*constants.RespStatus); ok {
			code = status.StatusCode
			msg = status.Error()
		} else {
			code = 500
			msg = "服务器内部错误"
		}
	}
	e.MethodByName("SetStatusCode").Call([]reflect.Value{reflect.ValueOf(code)})
	e.MethodByName("SetStatusMsg").Call([]reflect.Value{reflect.ValueOf(&msg)})
	return
}
