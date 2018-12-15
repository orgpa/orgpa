package configuration

import (
	"orgpa-database-api/database/dblayer"

	"github.com/kelseyhightower/envconfig"
)

// ServiceConfig contains the configuration of the micro-service
// Tags are used for the package "envconfig".
type ServiceConfig struct {
	DBType        dblayer.DBTYPE `envconfig:"DATABASE_TYPE" required:"true"`
	DBConnection  string         `envconfig:"DATABASE_CONNECTION" required:"true"`
	EndpointAPI   string         `split_words:"true" required:"true"`
	PasswordMySQL string         `envconfig:"DATABASE_PASSWORD_MYSQL" required:"true"`
	DatabaseName  string         `split_words:"true" required:"true"`
}

// ExtractConfiguration from a given filename
func ExtractConfiguration() (ServiceConfig, error) {
	var config ServiceConfig
	err := envconfig.Process("orgpa", &config)
	return config, err
}
