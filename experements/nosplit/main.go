package main

import (
	"reflect"
	"unsafe"

	"napisys"
)

import "C"

//go:noescape
func napi_define_properties()

//export Method
func Method(env napisys.Env, info napisys.CallbackInfo) napisys.Value {
	var res napisys.Value
	napisys.CreateInt32(env, C.int(5), &res)
	return res
}

//export Init
func Init(env Env, exports Value) napisys.Value {
	f := reflect.ValueOf(Method).Pointer()
	// var status NapiStatus
	desc := napisys.PropertyDescriptor{
		C.CString("hello"),
		nil,
		(napisys.Callback)(unsafe.Pointer(&f)),
		nil,
		nil,
		nil,
		napisys.Default,
		unsafe.Pointer(nil),
	}
	napi_define_properties(env, exports, 1, &desc)
	return exports
}

func main() {}
