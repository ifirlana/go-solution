package main

import (
	"github.com/ifirlana/go-solution/sections2/confformat"
)

func main()  {
	if err := confformat.MarshalAll();err != nil {
		panic(err)
	}

	if err := confformat.UnMarshalAll();err!=nil {
		panic(err)
	}

	if err:=confformat.OtherJSONExample();err!=nil {
		panic(err)
	}
}