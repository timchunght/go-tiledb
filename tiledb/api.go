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
	Context   C.struct_tiledb_ctx_t
	Config    C.struct_tiledb_config_t
	DataType  C.tiledb_datatype_t
	Attribute C.tiledb_attribute_t
)

const (
	TITLEDB_INT32 = C.TILEDB_INT32
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
func GroupCreate(ctx *Context, groupName string) error {
	ret := C.tiledb_group_create(
		(*C.struct_tiledb_ctx_t)(unsafe.Pointer(ctx)),
		C.CString(groupName),
	)
	fmt.Println(ret)

	if ret != 0 {
		return errors.New("Cannot create TileDB Group")
	}

	return nil
}

// int tiledb_attribute_create(
//     tiledb_ctx_t* ctx,
//     tiledb_attribute_t** attr,
//     const char* name,
//     tiledb_datatype_t type);
func AttributeCreate(ctx *Context, attr *Attribute, name string, dataType DataType) error {
	ret := C.tiledb_attribute_create(
		(*C.struct_tiledb_ctx_t)(unsafe.Pointer(ctx)),
		(*(*C.struct_tiledb_attribute_t))(unsafe.Pointer(&attr)),
		C.CString(name),
		C.tiledb_datatype_t(dataType),
	)
	if ret != 0 {
		return errors.New("Cannot create attribute")
	}
	return nil

}
