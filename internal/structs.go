package internal

import (
	"fmt"
	"reflect"
)

func SetStructValues(foo interface{}, envVars map[string]string) {
	val := reflect.ValueOf(foo)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		name := val.Type().Field(i).Name
		upperSnakeCaseName := toUpperSnakeCase(name)
		fmt.Printf("updated name: %s", upperSnakeCaseName)
		envVarVal, ok := envVars[upperSnakeCaseName]
		if !ok {
			continue
		}

		fmt.Println("setting env var")

		// TODO: Need to check if the value is unexported field
		val.Field(i).SetString(envVarVal)
	}
}
