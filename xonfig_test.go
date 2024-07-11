package xonfig

import (
	"errors"
	"os"
	"testing"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	A string
}

func TestLoad(t *testing.T) {
	os.Setenv("XONFIG", `
		a = "env"
	`)

	config, err := load[Config]()

	if err != nil {
		t.Fail()
	}

	if config.A != "env" {
		t.Fail()
	}
}

func TestMissingEnv(t *testing.T) {
	os.Unsetenv("XONFIG")

	_, err := load[Config]()

	if !errors.Is(err, ErrMissingEnv) {
		t.Fail()
	}
}

func TestEmptyEnv(t *testing.T) {
	os.Setenv("XONFIG", " \t\n")

	_, err := load[Config]()

	if !errors.Is(err, ErrEmptyEnv) {
		t.Fail()
	}
}

func TestMissingField(t *testing.T) {
	os.Setenv("XONFIG", `
		a = "env"
		b = "cat"
	`)

	_, err := load[Config]()

	if !errors.Is(err, ErrInvalidConfig) {
		t.Fail()
	}

	var strictMissingError *toml.StrictMissingError
	if !errors.As(err, &strictMissingError) {
		t.Fail()
	}
}
