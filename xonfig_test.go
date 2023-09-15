package xonfig

import (
	"os"
	"testing"
)

func TestMustLoad(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("unexpected error: \"%v\"", err)
		}
	}()

	os.Setenv("CONFIG", `
		a = "env"
	`)

	e := MustLoad[struct {
		A string
	}]()

	if e.A != "env" {
		t.Fail()
	}

	os.Unsetenv("CONFIG")

	f := MustLoad[struct {
		A string
	}]()

	if f.A != "file" {
		t.Fail()
	}
}

func TestMissingField(t *testing.T) {
	defer func() {
		err := recover().(error)
		if err == nil {
			t.Fail()
		} else if err.Error() !=
			"xonfig: strict mode: fields in the document are missing in the target struct" {
			t.Logf("unexpected error: \"%v\"", err)
		}
	}()
	MustLoad[struct{}]()
}
