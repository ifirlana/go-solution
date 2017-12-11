package main

import (
	"github.com/ifirlana/go-solution/sections1/bytestring"
)

func main()  {
	err := bytestring.WorkWithBuffer()
	if err != nil {
		panic(err)
	}

	bytestring.SearchString()
	bytestring.ModifyString()
	bytestring.StringReader()
}