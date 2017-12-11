package main

import (
	"fmt"
	"strings"
	"strconv"
)

type CountTheWay []int

func (c *CountTheWay) String() string  {

	result := ""

	for _, v := range *c {
		if len(result) > 0 {
			result += " ... "
		}
		result += fmt.Sprint(v)
	}

	return result
}

func (c *CountTheWay) Set(value string) error  {
	values := strings.Split(value, ",")

	for _, v := range values {
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}

		*c = append(*c, i)
	}

	return nil
}