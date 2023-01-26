package main

import (
	"flag"
	"log"
	"yadro/dklimov/test/internal/parsing"
	"yadro/dklimov/test/internal/transformation"
	"yadro/dklimov/test/internal/writing"
)

func main() {
	input := flag.String("input", "input.csv", "Файл для парсинга")
	output := flag.String("output", "output.csv", "Файл после парсинга")
	flag.Parse()

	records := parsing.Parser(*input)
	transRecords := transformation.Transform(records)

	if err := writing.WriterCsv(*output, transRecords); err != nil {
		log.Fatal(err)
	}
}
