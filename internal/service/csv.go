package service

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSV(csvPath string) (map[string]string, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return nil, fmt.Errorf("open csv file %s: %w", csvPath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("read csv headers: %w", err)
	}

	oldIDIndex := -1
	newIDIndex := -1
	for i, h := range headers {
		switch h {
		case "old_id":
			oldIDIndex = i
		case "new_id":
			newIDIndex = i
		}
	}

	if oldIDIndex == -1 || newIDIndex == -1 {
		return nil, fmt.Errorf("csv must contain old_id and new_id columns")
	}

	result := make(map[string]string)

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		if len(record) <= oldIDIndex || len(record) <= newIDIndex {
			continue
		}

		oldID := record[oldIDIndex]
		newID := record[newIDIndex]

		if oldID != "" && newID != "" {
			result[oldID] = newID
		}
	}

	return result, nil
}
