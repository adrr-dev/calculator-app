package repository

import (
	"encoding/json"
	"os"
)

type Repo struct {
	DataFile string
}

func (r Repo) FetchDisplay() (*Display, error) {
	file, err := os.Open(r.DataFile)
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(file)

	defer func() { _ = file.Close() }()

	var data Display
	err = dec.Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r Repo) WriteData(data *Display) error {
	file, err := os.Create(r.DataFile)
	if err != nil {
		return err
	}

	defer func() { _ = file.Close() }()

	enc := json.NewEncoder(file)
	err = enc.Encode(data)
	return err
}
