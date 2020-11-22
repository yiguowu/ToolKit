package main

import (
	"ToolKit/lib"
	"fmt"
	"os"
)

func main() {
	data := os.Args[1]
	secret := lib.GetSecretFromImageString(data)
	fmt.Println(lib.GetCurrentCode(secret))
}
