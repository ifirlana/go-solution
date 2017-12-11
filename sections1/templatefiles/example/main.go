package main

import (
	"github.com/ifirlana/go-solution/sections1/templatefiles"
)

func main() {
	if err := templatefiles.WorkWithTemp();err != nil {
		panic(err)
	}
}