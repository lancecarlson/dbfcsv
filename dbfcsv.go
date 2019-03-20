package main

import (
	"encoding/csv"
	"os"
	"flag"
	"unicode/utf8"
	"github.com/LindsayBradford/go-dbf/godbf"
)

func main() {
	delimiter := flag.String("d", "|", "delimiter used to separate fields")
	headers := flag.Bool("h", false, "display headers")
	flag.Parse()
	path := flag.Arg(0)
	if path == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	
	dbfTable, err := godbf.NewFromFile(path, "UTF8")
	if err != nil { panic(err) }

	comma, _ := utf8.DecodeRuneInString(*delimiter)
	out := csv.NewWriter(os.Stdout)
	out.Comma = comma

	if *headers { 
		fields := dbfTable.Fields()
		fieldRow := make([]string, len(fields))
		for i := 0; i < len(fields); i++ {
			fieldRow[i] = fields[i].Name()
		}
		out.Write(fieldRow)
		out.Flush()
	}

	// Output rows
	for i := 0; i < dbfTable.NumberOfRecords(); i++ {
		row := dbfTable.GetRowAsSlice(i)
		out.Write(row)
		out.Flush()
	}
}
