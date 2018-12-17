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

// ExtractConfiguration will extract the configuration from
// the environement and return a ServiceConfig struct containing
// the whole service configuration.
//
// If an environment variable is missing a non nil error will be
// returned.
func ExtractConfiguration() (ServiceConfig, error) {
	var config ServiceConfig
	err := envconfig.Process("orgpa", &config)
	return config, err
}
