package service

import (
	"encoding/json"
	"fmt"
	"os"

	"e4cm/internal/model"
)

func ReadJSON(jsonPath string) ([]model.TwikooComment, error) {
	file, err := os.Open(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("open json file %s: %w", jsonPath, err)
	}
	defer file.Close()

	var comments []model.TwikooComment
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&comments); err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}

	return comments, nil
}
