package bytestring

import (
	"fmt"
	"bytes"
	"bufio"
)

func WorkWithBuffer() error {
	rawString := "its easy to encode unicode into a byte array"

	b := Buffer(rawString)
	fmt.Println(b.String())

	s, err := toString(b)
	if err != nil {
		return err
	}

	fmt.Println(s)

	reader := bytes.NewBuffer([]byte(rawString))
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return nil
}