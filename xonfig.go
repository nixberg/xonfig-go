package xonfig

import (
	"errors"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

var (
	ErrMissingEnv    = errors.New("XONFIG environment variable not present")
	ErrEmptyEnv      = errors.New("XONFIG environment variable contains only white space")
	ErrInvalidConfig = errors.New("invalid config")
)

func load[Config any]() (Config, error) {
	var config Config
	configString, present := os.LookupEnv("XONFIG")

	if !present {
		return config, ErrMissingEnv
	} else if strings.TrimSpace(configString) == "" {
		return config, ErrEmptyEnv
	}

	decoder := toml.NewDecoder(strings.NewReader(configString))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&config); err != nil {
		return config, errors.Join(ErrInvalidConfig, err)
	}
	return config, nil
}

func MustLoad[Config any]() Config {
	config, err := load[Config]()
	if err != nil {
		panic(err)
	}
	return config
}
