package main

import (
	"ToolKit/lib"
	"os"
)

func main() {
	data := os.Args[1]
	lib.GetSecretFromImageString(data)

}
