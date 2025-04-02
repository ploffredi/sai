package data

import (
	"encoding/json"
	"errors"
	"os"
)

type Software struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
	ConfigFile  string   `json:"config_file"`
	Tags        []string `json:"tags"`
}

var softwareData []Software

func LoadData(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&softwareData)
	if err != nil {
		return err
	}
	return nil
}

func GetSoftware(name string) (*Software, error) {
	for _, s := range softwareData {
		if s.Name == name {
			return &s, nil
		}
	}
	return nil, errors.New("software not found")
}
