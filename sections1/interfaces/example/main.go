package example

import (
	"fmt"
	"github.com/ifirlana/go-solution/sections1/interfaces"
	"bytes"
)

func main() {
	in := bytes.NewBuffer([]byte("example"))
	out := &bytes.Buffer{}

	fmt.Println("stdout on Copy = ")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}

	fmt.Println("out bytes buffer ", out.String())
	fmt.Println("stdout on Pipeexample = ")
	if err:= interfaces.PipeExample(); err != nil {
		panic(err)
	}

}