package main

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
