package templates

import (
	"html/template"
	"strings"
	"os"
)

const sampleTemplate = `
This template demonstrate printing a {{.Variable | printf "%#v"}}.

{{if .Condition }}
If condition is set, we ll print this
{{else}}
Otherwise, we'll print this instead
{{end}}

Next we'll iterate over an array of strings:
{{range $index, $item := .Items}}
{{$index}}: {{$item}}
{{end}}
`

const sencondTemplate = `
{{define "block_example"}}
{{.Otherwise}}
{{end}}
`

func RunTemplate() error {

	data := struct {
		Condition bool
		Variable string
		Items []string
		Words string
		OtherVariable string
	}{
		Condition: true,
		Variable: "variable",
		Items: []string{"Item1", "Item2", "Item3"},
		Words:"another_item1,another_item2,another_item3",
		OtherVariable: "I'm defined in a second template!",
	}

	funcmap := template.FuncMap{
		"split": strings.Split,
	}

	t := template.New("example ")
	t = t.Funcs(funcmap)
	t, err := t.Parse(sampleTemplate)
	if err != nil {
		return err
	}

	t2, err := t.Clone()
	if err != nil {
		return err
	}

	t2, err = t2.Parse(sencondTemplate)
	if err != nil {
		return err
	}

	err = t2.Execute(os.Stdout, &data)
	if err != nil {
		return err
	}

	return nil
}