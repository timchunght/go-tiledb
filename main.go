package main

import (
	"fmt"
	"go-tiledb/tiledb"
)

func main() {
	major, minor, rev := tiledb.TileDBVersion()

	tiledb.TileDBCtxCreate()
	fmt.Printf("TileDB v%v.%v.%v\n", major, minor, rev)
}
