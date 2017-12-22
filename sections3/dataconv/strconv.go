package dataconv

import (
	"strconv"
	"fmt"
)

func Strconv() error  {
	s := "1234"
	res, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return err
	}

	fmt.Println(res)

	res, err = strconv.ParseInt("FF", 16, 64)
	if err != nil {
		return err
	}

	fmt.Println(res)

	val, err := strconv.ParseBool("True")
	if err != nil {
		return err
	}

	fmt.Println(val)


	return nil
}