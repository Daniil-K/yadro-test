package parsing

import (
	"encoding/csv"
	"log"
	"os"
)

func Parser(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read input file %s %v \n", filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to parse file as CSV for %s %v \n", filePath, err)
	}

	return records
}
