package utils

import (
	"encoding/csv"
	"os"
	"path/filepath"
)

func GetFile() (*os.File, error) {
	filepath := filepath.Join(os.Getenv("HOME"), ".cache", "task.csv")

	_, err := os.Stat(filepath)
	isFileExists := !os.IsNotExist(err)
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	if !isFileExists {
		writer := csv.NewWriter(file)
		headers := []string{"Task", "Priority"}
		if err := writer.Write(headers); err != nil {
			file.Close()
			return nil, err
		}
		writer.Flush()                         // Flush the writer to ensure headers are written
		if err := writer.Error(); err != nil { // Check for errors during flushing
			file.Close()
			return nil, err
		}
	}

	return file, nil
}
