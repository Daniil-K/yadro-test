package transformation

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	reAll    = regexp.MustCompile("[^A-Za-z0-9_]+")
	reChar   = regexp.MustCompile("[0-9]+")
	reDigit  = regexp.MustCompile("[^0-9]+")
	reSymbol = regexp.MustCompile("[A-Za-z0-9_]+")
)

func mapColumns(records [][]string) map[string]int {
	columns := make(map[string]int)

	for key, value := range records[0] {
		if value == "" {
			continue
		}
		columns[value] = key
	}

	return columns
}

func mapRows(records [][]string) map[string]int {
	rows := make(map[string]int)

	for key, value := range records {
		for k, val := range value {
			if k == 0 && val != "" {
				rows[val] = key
			}
		}
	}

	return rows
}

func Transform(records [][]string) [][]string {
	columns := mapColumns(records)
	rows := mapRows(records)

	for key, value := range records {
		for k, val := range value {
			if strings.Contains(val, "=") {
				newValue, err := formulaToValue(val[1:], columns, rows, records)
				if err != nil {
					log.Println(err)
					continue
				}
				records[key][k] = newValue
			}
		}
	}

	return records
}

func formulaToValue(template string, cols, rows map[string]int, records [][]string) (string, error) {
	var formula []string

	cells := reAll.Split(template, -1)

	for _, cell := range cells {
		splitCol := reChar.Split(cell, -1)[0]
		splitRow := reDigit.Split(cell, -1)[1]

		colNumber, ok := cols[splitCol]
		if !ok {
			return template, fmt.Errorf("Column %s not found \n", splitCol)
		}

		rowNumber, ok := rows[splitRow]
		if !ok {
			return template, fmt.Errorf("Row %s not found \n", splitRow)
		}

		formula = append(formula, records[rowNumber][colNumber])
	}

	var res string
	var err error

	switch {
	case strings.Contains(template, "+"):
		res, err = sum(formula)
	case strings.Contains(template, "-"):
		res, err = sub(formula)
	case strings.Contains(template, "*"):
		res, err = multiply(formula)
	case strings.Contains(template, "/"):
		res, err = div(formula)
	default:
		return template, errors.New("unknown operation " + template)
	}

	if err != nil {
		return template, err
	}

	return res, nil

}

func sum(values []string) (string, error) {
	digit1, err := strconv.Atoi(values[0])
	if err != nil {
		return "", fmt.Errorf("Failed to convern to int")
	}

	digit2, err := strconv.Atoi(values[1])
	if err != nil {
		return "", fmt.Errorf("Failed to convern to int")
	}

	return strconv.Itoa(digit1 + digit2), nil
}

func sub(values []string) (string, error) {
	digit1, err := strconv.Atoi(values[0])
	if err != nil {
		return "", fmt.Errorf("Failed to convern to int")
	}

	digit2, err := strconv.Atoi(values[1])
	if err != nil {
		return "", fmt.Errorf("Failed to convern to int")
	}

	return strconv.Itoa(digit1 - digit2), nil
}

func multiply(values []string) (string, error) {
	digit1, err := strconv.Atoi(values[0])
	if err != nil {
		return "", fmt.Errorf("Failed to convern to int")
	}

	digit2, err := strconv.Atoi(values[1])
	if err != nil {
		return "", fmt.Errorf("Failed to convern to int")
	}

	return strconv.Itoa(digit1 * digit2), nil
}

func div(values []string) (string, error) {
	digit1, err := strconv.Atoi(values[0])
	if err != nil {
		return "", fmt.Errorf("Failed to convern to int")
	}

	digit2, err := strconv.Atoi(values[1])
	if err != nil {
		return "", fmt.Errorf("Failed to convern to int")
	}

	if digit2 == 0 {
		return "", fmt.Errorf("Failed to division on 0")
	}

	return strconv.Itoa(digit1 / digit2), nil
}
