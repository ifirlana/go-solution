package confformat

import(
	"fmt"
)

func MarshalAll() error  {
	t := TOMLData{
		Name:"toml name",
		Age:12,
	}

	j := JSONData{
		Name:"json name",
		Age:10,
	}

	y := YAMLData{
		Name:"yml name",
		Age:1,
	}

	tomRes, err := t.ToTOML()
	if err != nil {
		return err
	}
	fmt.Println("TOML Marshal =", tomRes)

	jsonRes, err := j.ToJSON()
	if err != nil {
		return err
	}
	fmt.Println("JSON Marshal = ", jsonRes)


	yamlRes, err := y.ToYAML()
	if err != nil {
		return err
	}
	fmt.Println("YAML Marshal = ", yamlRes)

	return nil
}