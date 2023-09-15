package xonfig

import (
	"bytes"
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

func MustLoad[Config any]() (config Config) {
	configString, present := os.LookupEnv("CONFIG")

	if present {
		mustDecode(&config, []byte(configString))
	} else {
		configBytes, err := os.ReadFile("config.toml")
		if err != nil {
			panic(fmt.Errorf("xonfig: %w", err))
		}
		mustDecode(&config, configBytes)
	}

	return config
}

func MustLoadOr[Config any](config Config) Config {
	configString, present := os.LookupEnv("CONFIG")

	if present {
		mustDecode(&config, []byte(configString))
	}

	return config
}

func mustDecode[Config any](config *Config, configBytes []byte) {
	decoder := toml.NewDecoder(bytes.NewReader(configBytes))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&config); err != nil {
		panic(fmt.Errorf("xonfig: %w", err))
	}
}
