package config

type (
	Config struct {
		// The application host address.
		Host string `envconfig:"host" default:"0.0.0.0"`

		// The application host port.
		Port string `envconfig:"port" default:"8000"`
	}
)
