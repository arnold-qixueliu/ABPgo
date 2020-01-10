package main

import (
	"net/http"
	"reflect"
	"strings"
)

var mapStruct = make(map[string]reflect.Type)

type AutoRouter struct {
}
type Context struct {
	Response http.ResponseWriter
	Request  http.Request
}

func RegisterType(stru interface{}) {
	typ := reflect.TypeOf(stru)
	structName := typ.Name()
	mapStruct[strings.ToLower(structName)] = typ
}
func (router AutoRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	arr := strings.Split(path[1:], "/")
	typeName := strings.ToLower(arr[0])
	actor := mapStruct[typeName]
	v := reflect.New(actor).Elem()
	v.Field(0).SetInt(2000)
	method := v.MethodByName("GetUser")
	param := make([]interface{}, 1)
	param[0] = 200
	params := make([]reflect.Value, len(param)) //参数
	for k, v := range param {
		params[k] = reflect.ValueOf(v)
	}
	method.Call(nil)

	///////
	res.Write([]byte(req.URL.Path))
}
