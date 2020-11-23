package main

import (
	"ToolKit/lib"
	"fmt"
	"os"
	"strings"
)

const Version = "1.0"

func creationDispatch(params lib.ParamSet) {

}

func updateDispatch(params lib.ParamSet) {

}

func deleteDispatch(params lib.ParamSet) {

}

func listDispatch(params lib.ParamSet) {

}

func dispatch(action string, params lib.ParamSet) {
	switch action {
	case "create":
		creationDispatch(params)
	case "update":
		updateDispatch(params)
	case "delete":
		deleteDispatch(params)
	case "list":
		listDispatch(params)
	default:
		if strings.HasPrefix(action, "-") {
			fmt.Println("Missing action type")
			os.Exit(2)
		} else {
			fmt.Println("Unknown action type", action)
			os.Exit(3)
		}
	}
}

func main() {
	params := lib.ParseParameter()
	if len(os.Args) == 1 {
		fmt.Println("awsctl version", Version)
	} else {
		dispatch(os.Args[1], params)
	}
}
