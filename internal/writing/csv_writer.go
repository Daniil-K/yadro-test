package writing

import (
	"encoding/csv"
	"os"
)

func WriterCsv(filePath string, records [][]string) error {
	f, err := os.Create(filePath)
	defer f.Close()

	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(records)

	if err != nil {
		return err
	}

	return nil
}
