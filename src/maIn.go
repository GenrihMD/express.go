package main

import (
	"reflect"
	"unsafe"
)

/*
// #cgo linux CFLAGS: -I/usr/include/nodejs/src
// #cgo darwin CFLAGS: -I/usr/local/include/node
#cgo darwin LDFLAGS: -L. -lnode_api

#include "node_api.h"
*/
import "C"

// Method fu
func Method(env NapiEnv, info NapiCallbackInfo) NapiValue {
	var world NapiValue
	return world
}

// Init addon function
func Init(env NapiEnv, exports NapiValue) NapiValue {
	f := reflect.ValueOf(Method).Pointer()
	// var status NapiStatus
	desc := C.napi_property_descriptor{
		C.CString("hello"),
		nil,
		(C.napi_callback)(unsafe.Pointer(&f)),
		nil,
		nil,
		nil,
		C.napi_default,
		unsafe.Pointer(nil),
	}
	C.napi_define_properties(env, exports, 1, &desc)
	return exports
}

func main() {}
