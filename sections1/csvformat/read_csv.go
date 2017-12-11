package csvformat

import (
	"io"
	"encoding/csv"
	"strconv"
	"bytes"
	"fmt"
)

type Movie struct {
	Title string
	Director string
	Year int
}

func ReadCSV(b io.Reader) ([]Movie, error) {
	r := csv.NewReader(b)
	r.Comma = ';'
	r.Comment = '-'
	var movies []Movie
	_, err := r.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		year, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			return nil, err
		}

		m := Movie{record[0], record[1], int(year)}
		movies = append(movies, m)
	}

	return movies, nil
}

func AddMoviesFromText() error  {
	in := `
- first our headers
movie title; director;year released

- then some data
Guardian of the Galaxy Vol.2;James Gun;2017
Star Wars: Episode VIII;Rian Johnson;2017
`
	b := bytes.NewBufferString(in)
	m, err := ReadCSV(b)
	if err != nil {
		return err
	}

	fmt.Printf("%#v\n", m)
	return nil
}