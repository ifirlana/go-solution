package main

import (
	"github.com/ifirlana/go-solution/sections1/filedir"
)

func main()  {
	if err := filedir.Operate();err != nil {
		panic(err)
	}

	if err := filedir.CapitalizerExample();err != nil {
		panic(err)
	}
}