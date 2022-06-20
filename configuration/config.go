package configuration

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port struct {
		Insert string `json:"insert"`
		Delete string `json:"delete"`
		Update string `json:"update"`
	} `json:"port"`

	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Dbname   string `json:"dbname"`
		SSLMode  string `json:"sslmode"`
	} `json:"database"`
}

func LoadConfig(file string, cfg interface{}) error {
	r, err := os.Open(file)
	if err != nil {
		return err
	}
	defer r.Close()

	return json.NewDecoder(r).Decode(&cfg)
}
