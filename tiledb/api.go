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

func TileDBVersion() (major, minor, rev int) {
	var majorC C.int
	var minorC C.int
	var revC C.int
	C.tiledb_version(&majorC, &minorC, &revC)
	return int(majorC), int(minorC), int(revC)

}

// int tiledb_ctx_create(tiledb_ctx_t** ctx)
func TileDBCtxCreate() (*Context, error) {
	// (*(*C.tiledb_ctx_t))(unsafe.Pointer(&ctx))
	ctx := &Context{}
	ctxPtr := unsafe.Pointer(ctx)
	result := C.tiledb_ctx_create((*(*C.struct_tiledb_ctx_t))(ctxPtr), nil)
	fmt.Println(result)
	if result != 0 {
		return nil, errors.New("Cannot create TileDB Context")
	}
	return ctx, nil
}
