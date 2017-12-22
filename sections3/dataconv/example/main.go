package main

import (
	"github.com/ifirlana/go-solution/sections3/dataconv"
)

func main()  {
	dataconv.ShowConv()
	if err:=dataconv.Strconv();err != nil {
		panic(err)
	}

	dataconv.Interface()
}