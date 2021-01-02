// This file is part of go-savetime.
// Copyright (c) 2021 Claudio Guarnieri
// See the file 'LICENSE' for copying permission.

package files

import (
	"encoding/json"
	"os"
)

// ReadJSON reads a JSON file at the specified path into an interface.
func ReadJSON(jsonPath string, data interface{}) error {
	file, err := os.Open(jsonPath)
	if err != nil {
		return err
	}

	return json.NewDecoder(file).Decode(data)
}

// SaveJSON encodes an interface into JSON and saves it to a file.
func SaveJSON(jsonPath string, data interface{}) error {
	file, err := os.OpenFile(jsonPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		return err
	}

	return nil
}
