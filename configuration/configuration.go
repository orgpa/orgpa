package configuration

import (
	"encoding/json"
	"orgpa-database-api/database/dblayer"
	"os"
)

const (
	DBTypeDefault        = dblayer.DBTYPE("mysql")
	DBConnectionDefault  = "127.0.0.1:3306"
	EndpointAPIDefault   = "localhost:9900"
	PasswordMySQLDefault = "test"
	DatabaseNameDefault  = "orgpa_user_api"
)

// ServiceConfig contains the configuration of the micro-service
type ServiceConfig struct {
	DBType        dblayer.DBTYPE `json:"dbtype"`
	DBConnection  string         `json:"dbconnection"`
	EndpointAPI   string         `json:"endpointapi"`
	PasswordMySQL string         `json:"passwordMySQL"`
	DatabaseName  string         `json:"databaseName"`
}

// ExtractConfiguration from a given filename
func ExtractConfiguration(filename string) (ServiceConfig, error) {
	config := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		EndpointAPIDefault,
		PasswordMySQLDefault,
		DatabaseNameDefault,
	}

	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	err = json.NewDecoder(file).Decode(&config)
	return config, err
}
