package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type record struct {
	Date time.Time
	Open float64
}

func foo(res http.ResponseWriter, req *http.Request) {
	records := prs("table.csv") // parse csv

	// parse template
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// execute template
	err = tpl.Execute(res, records)
	if err != nil {
		log.Fatalln(err)
	}
}

func prs(filePath string) []record {
	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	rdr := csv.NewReader(src)
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	records := make([]record, 0, len(rows))

	for i, row := range rows {

		if i == 0 || i == 1 {
			continue
		}

		s := strings.Split(row[0], ";")

		date, _ := time.Parse("2006-01-02", s[0])

		open, _ := strconv.ParseFloat(s[1], 64)

		records = append(records, record{
			Date: date,
			Open: open,
		})
	}
	return records
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}
