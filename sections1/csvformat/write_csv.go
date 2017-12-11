package csvformat

import (
	"io"
	"encoding/csv"
	"os"
	"bytes"
)

type Book struct  {
	Author string
	Title string
}

type Books []Book

func (book *Books) ToCSV(w io.Writer) error  {
	n := csv.NewWriter(w)
	err := n.Write([]string{"Author", "Title"})
	if err != nil {
		return err
	}
	return nil
}

func WriteCSVOutput() error {
	b := Books{
		Book{
			Author:"F Scott Fitzgerald",
			Title:"The Great Gastby",
		},
		Book{
			Author:"J D Salinger",
			Title: "The Catcher in the Rye",
		},
	}

	return b.ToCSV(os.Stdout)
}

func WriteCSVBuffer() (*bytes.Buffer, error) {
	//b := Books{
	//	Book{
	//		Author:"F Scott Fitzgerald",
	//		Title:"The Great Gastby",
	//	},
	//	Book{
	//		Author:"J D Salinger",
	//		Title: "The Catcher in the Rye",
	//	},
	//}

	//w := &bytes.Buffer{}

	return nil, nil
}