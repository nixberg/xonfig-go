package xonfig

import (
	"fmt"
	"os"
	"reflect"
)

func MustLoad(config interface{}) {
	configType := reflect.TypeOf(config)
	if configType.Kind() != reflect.Ptr {
		panic("xonfig: argument is not a pointer")
	}

	configType = configType.Elem()
	if configType.Kind() != reflect.Struct {
		panic("xonfig: argument is not a pointer to a struct")
	}

	configValue := reflect.ValueOf(config).Elem()

	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)

		if field.Anonymous {
			panic("xonfig: struct contains anonymous field")
		}

		if field.PkgPath != "" {
			panic(fmt.Sprintf("xonfig: struct contains unexported field \"%s\"", field.Name))
		}

		if field.Type.Kind() != reflect.String {
			panic(fmt.Sprintf("xonfig: type of field \"%s\" is not string", field.Name))
		}

		envTag := field.Tag.Get("env")
		if len(envTag) == 0 {
			panic(fmt.Sprintf("xonfig: missing \"env\" tag for field \"%s\"", field.Name))
		}

		env, exists := os.LookupEnv(envTag)
		if !exists {
			panic(fmt.Sprintf("xonfig: missing environment variable \"%s\"", envTag))
		}

		configValue.FieldByName(field.Name).SetString(env)
	}
}
