package repository

import (
	"os"
	"strings"
)

func Read() ([][]string, error) {
	readData := [][]string{}
	data, err := os.ReadFile("./tickets.csv")
	if err != nil {
		return [][]string{}, err
	}
	stringData := string(data)
	sliceData := strings.Split(stringData, "\n")
	for _, row := range sliceData {
		sliceRow := strings.Split(row, ",")
		readData = append(readData, sliceRow)
	}

	return readData, nil
}
