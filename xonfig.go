package xonfig

import (
	"bytes"
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

func MustLoad[Config any]() Config {
	configString, present := os.LookupEnv("CONFIG")
	configBytes := []byte(configString)

	if !present {
		var err error
		configBytes, err = os.ReadFile("config.toml")
		if err != nil {
			panic(fmt.Errorf("xonfig: %w", err))
		}
	}

	decoder := toml.NewDecoder(bytes.NewReader(configBytes))
	decoder.DisallowUnknownFields()

	var config Config
	if err := decoder.Decode(&config); err != nil {
		panic(fmt.Errorf("xonfig: %w", err))
	}

	return config
}
