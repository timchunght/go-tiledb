package main

import (
	"fmt"
	"go-tiledb/tiledb"
)

func main() {
	major, minor, rev := tiledb.Version()
	fmt.Printf("TileDB v%v.%v.%v\n", major, minor, rev)
	ctx, _ := tiledb.CtxCreate()
	tiledb.GroupCreate(ctx, "myGroup2")

}
