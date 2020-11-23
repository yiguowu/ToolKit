package lib

import "flag"

type ParamSet struct {
	File *string
}

func ParseParameter() ParamSet {
	var params ParamSet
	params.File = flag.String("f", "", "Input file")
	flag.Parse()
	return params
}
