package lib

import (
	"io/ioutil"
	"os"
	"strconv"
)

func ReadFile(fileName string, optional ...string) []byte {
	path := ""
	if len(optional) > 0 {
		path = optional[0]
		fileName = path + "/" + fileName
	}
	dat, _ := ioutil.ReadFile(fileName)
	return dat
}

func WriteFile(dat []byte, fileName string, optional ...string) {
	path := ""
	mode := os.FileMode.Perm(0644)
	if len(optional) > 0 {
		fileMode, _ := strconv.Atoi(optional[0])
		mode = os.FileMode.Perm(os.FileMode(fileMode))
	}
	if len(optional) > 1 {
		path = optional[1]
		fileName = path + "/" + fileName
	}
	_ = ioutil.WriteFile(fileName, dat, mode)
}
