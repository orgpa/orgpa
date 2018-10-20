package configuration

import (
	"encoding/json"
	"orgpa/orgpa-database-api/database/dblayer"
	"os"
)

const (
	DBTypeDefault       = dblayer.DBTYPE("mongodb")
	DBConnectionDefault = "mongodb://127.0.0.1:27017"
	EndpointAPIDefault  = "localhost:9900"
)

// ServiceConfig contains the configuration of the micro-service
type ServiceConfig struct {
	DBType       dblayer.DBTYPE `json:"dbtype"`
	DBConnection string         `json:"dbconnection"`
	EndpointAPI  string         `json:"endpointapi"`
}

// ExtractConfiguration from a given filename
func ExtractConfiguration(filename string) (ServiceConfig, error) {
	config := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		EndpointAPIDefault,
	}

	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	err = json.NewDecoder(file).Decode(&config)
	return config, err
}
