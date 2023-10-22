package utils

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/KunalDuran/duranz-stats-cli/models"
)

func ListJSONFiles(path string) []string {
	var fileList []string

	files, err := os.ReadDir(path)
	if err != nil {
		return fileList
	}

	for _, f := range files {
		if strings.Contains(f.Name(), "json") {
			fileList = append(fileList, f.Name())
		}
	}
	return fileList
}

func ReadMatchJSON(f_path string) (models.Match, error) {
	var matchData models.Match
	body, err := os.ReadFile(f_path)
	if err != nil {
		return matchData, err
	}

	err = json.Unmarshal(body, &matchData)
	if err != nil {
		return matchData, err
	}
	return matchData, nil
}
