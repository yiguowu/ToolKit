package main

import (
	"ToolKit/lib"
	"fmt"
)

func main() {
	lib.InitDefault()
	fmt.Println(lib.DefaultToolKitConfig)
}
