package xonfig

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("a", "it")
	os.Setenv("b", "works")
}

func TestItWorks(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("unexpected error: \"%v\"", err)
		}
	}()

	value := struct {
		A string `env:"a"`
		B string `env:"b"`
	}{}
	MustLoad(&value)

	if value.A != "it" || value.B != "works" {
		t.Fail()
	}
}

func TestPointerArgument(t *testing.T) {
	assertPanics(t, func() {
		MustLoad("")
	}, "xonfig: argument is not a pointer")
}

func TestNonStructArgument(t *testing.T) {
	value := ""
	assertPanics(t, func() {
		MustLoad(&value)
	}, "xonfig: argument is not a pointer to a struct")
}

func TestAnonymousField(t *testing.T) {
	value := struct {
		A      string `env:"a"`
		string `env:"b"`
	}{}
	assertPanics(t, func() {
		MustLoad(&value)
	}, "xonfig: struct contains anonymous field")
}

func TestUnexportedField(t *testing.T) {
	value := struct {
		A string `env:"a"`
		b string `env:"b"`
	}{}
	assertPanics(t, func() {
		MustLoad(&value)
	}, "xonfig: struct contains unexported field \"b\"")
}

func TestNonStringield(t *testing.T) {
	value := struct {
		A string `env:"a"`
		B int    `env:"b"`
	}{}
	assertPanics(t, func() {
		MustLoad(&value)
	}, "xonfig: type of field \"B\" is not string")
}

func TestMissingEnvTag(t *testing.T) {
	value := struct {
		A string `env:"a"`
		B string
	}{}
	assertPanics(t, func() {
		MustLoad(&value)
	}, "xonfig: missing \"env\" tag for field \"B\"")
}

func TestMissingEnv(t *testing.T) {
	value := struct {
		A string `env:"a"`
		B string `env:"bz2g1zpn7fgd5"`
	}{}
	assertPanics(t, func() {
		MustLoad(&value)
	}, "xonfig: missing environment variable \"bz2g1zpn7fgd5\"")
}

func assertPanics(t *testing.T, f func(), expected string) {
	defer func() {
		err := recover()
		if err == nil {
			t.Fail()
		} else if err != expected {
			t.Errorf("unexpected error: \"%v\"", err)
		}
	}()
	f()
}
