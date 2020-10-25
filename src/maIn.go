package main

import (
	"unsafe"
)

/*
#cgo linux CFLAGS: -I/usr/include/nodejs/src
#cgo darwin CFLAGS: -I/usr/local/include/node
#cgo darwin LDFLAGS: -L. -lnode_api

#include <node_api.h>
*/
import "C"

// NapiValue is binding to C.napi_value
type NapiValue C.napi_value

// NapiEnv is binding to C.napi_env
type NapiEnv C.napi_env

// NapiStatus is binding to C.napi_status
type NapiStatus C.napi_status

// NapiCallbackInfo is binding to C.napi_callback_info
type NapiCallbackInfo C.napi_callback_info

// Method fu
func Method(env NapiEnv, info NapiCallbackInfo) NapiValue {
	var world NapiValue
	return world
}

// Init addon function
func Init(env NapiEnv, exports NapiValue) NapiValue {
	// var status NapiStatus
	desc := C.napi_property_descriptor{
		C.CString("hello"),
		nil,
		(C.napi_callback)(unsafe.Pointer(&Method)),
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
