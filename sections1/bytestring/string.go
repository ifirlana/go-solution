package bytestring

import (
	"fmt"
	"strings"
	"io"
	"os"
)

func SearchString()  {
	s := "this is test"
	fmt.Println(strings.Contains(s, "this"))
	fmt.Println(strings.Contains(s, "abc"))
	fmt.Println(strings.Contains(s, "this"))
	fmt.Println(strings.Contains(s, "this"))
	fmt.Println(strings.Contains(s, "test"))
}

func ModifyString() {
	s := "this is test"
	fmt.Println(strings.Split(s, " "))
	fmt.Println(strings.Title(s))
	s = " simple string "
	fmt.Println(strings.TrimSpace(s))
}

func StringReader() {
	s := "simple string\n"
	r := strings.NewReader(s)
	io.Copy(os.Stdout, r)

}