package xonfig

import (
	"os"
	"testing"
)

func TestItWorks(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("unexpected error: \"%v\"", err)
		}
	}()

	value := MustLoad[struct {
		A string
		B string
	}]()

	if value.A != "it" || value.B != "works" {
		t.Fail()
	}
}

func TestItWorksEnv(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("unexpected error: \"%v\"", err)
		}
	}()

	os.Setenv("CONFIG", `c = "also works"`)

	value := MustLoad[struct {
		C string
	}]()

	if value.C != "also works" {
		t.Fail()
	}
}

func TestMissingField(t *testing.T) {
	assertPanics(t, func() {
		MustLoad[struct {
			A string
		}]()
	}, "xonfig: strict mode: fields in the document are missing in the target struct")
}

func assertPanics(t *testing.T, f func(), expected string) {
	defer func() {
		err := recover().(error)
		if err == nil {
			t.Fail()
		} else if err.Error() != expected {
			t.Logf("unexpected error: \"%v\"", err)
		}
	}()
	f()
}
