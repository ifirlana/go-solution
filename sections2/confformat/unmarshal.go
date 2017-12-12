package confformat

import (
	"fmt"
)

const (
	exampleTOML = `name="EXAMPLE1"
age=99`
	exampleJSON = `{"name":"EXAMPLE2", "age":98}`
	exampleYAML = `name: EXAMPLE3 age:97`
)

func UnMarshalAll() error {
	t:=TOMLData{}
	y:=YAMLData{}
	j:=JSONData{}

	if _,err := t.Decode([]byte(exampleTOML));err != nil {
		return err
	}
	fmt.Println("TOML Unmarshal =", t)

	if err := j.Decode([]byte(exampleJSON));err != nil {
		return err
	}
	fmt.Println("JSON Unmarshal = ", j)

	if err := y.Decode([]byte(exampleYAML));err !=nil {
		return err
	}
	fmt.Println("YAML Unmarshal = ", y)

	return nil
}