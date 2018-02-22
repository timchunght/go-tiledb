package tiledb

/*
#include "tiledb.h"
#cgo LDFLAGS: -ltiledb
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

type (
	Context C.struct_tiledb_ctx_t
	Config  C.struct_tiledb_config_t
)

func Version() (major, minor, rev int) {
	var majorC C.int
	var minorC C.int
	var revC C.int
	C.tiledb_version(&majorC, &minorC, &revC)
	return int(majorC), int(minorC), int(revC)

}

// int tiledb_ctx_create(tiledb_ctx_t** ctx)
func CtxCreate() (*Context, error) {
	// (*(*C.tiledb_ctx_t))(unsafe.Pointer(&ctx))
	ctx := &Context{}
	ctxPtr := unsafe.Pointer(&ctx)
	ret := C.tiledb_ctx_create((*(*C.struct_tiledb_ctx_t))(ctxPtr), nil)
	fmt.Println(ret)
	if ret != 0 {
		return nil, errors.New("Cannot create TileDB Context")
	}
	return ctx, nil
}

// int tiledb_group_create(tiledb_ctx_t* ctx, const char* group)

func GroupCreate(ctx *Context, groupName string) {
	// ctx := &Context{}
	// C.tiledb_ctx_create((*(*C.struct_tiledb_ctx_t))(unsafe.Pointer(&ctx)), nil)
	ret := C.tiledb_group_create((*C.struct_tiledb_ctx_t)(unsafe.Pointer(ctx)), C.CString(groupName))
	fmt.Println(ret)
}
