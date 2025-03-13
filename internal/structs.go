package internal

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"unicode"
)

func isExportedField(fieldName string) bool {
	return unicode.IsUpper(rune(fieldName[0]))
}

func parseEnvVarValue(envVarValue string, fieldKind reflect.Kind) (interface{}, error) {
	switch fieldKind {
	case reflect.String:
		return envVarValue, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.ParseInt(envVarValue, 10, 64)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.ParseUint(envVarValue, 10, 64)
	case reflect.Float32, reflect.Float64:
		return strconv.ParseFloat(envVarValue, 64)
	case reflect.Bool:
		return strconv.ParseBool(envVarValue)
	default:
		return nil, fmt.Errorf("unsupported type: %s", fieldKind.String())
	}
}

func SetStructValues(foo interface{}, envVars map[string]string) error {
	val := reflect.ValueOf(foo)
	if val.Kind() != reflect.Pointer {
		return errors.New("struct passed is not a pointer")
	}
	val = val.Elem()

	// Loop through each field in struct
	for i := 0; i < val.NumField(); i++ {
		name := val.Type().Field(i).Name
		if !isExportedField(name) {
			fmt.Printf("skipping unexported field in struct: %s\n", name)
			continue
		}

		upperSnakeCaseName := toUpperSnakeCase(name)
		fmt.Printf("checking for env var: %s\n", upperSnakeCaseName)
		envVarVal, ok := envVars[upperSnakeCaseName]
		if !ok {
			continue
		}

		fieldVal := val.Field(i)
		parsedVal, err := parseEnvVarValue(envVarVal, fieldVal.Kind())
		if err != nil {
			return fmt.Errorf("failed to parse value for %s: %w", upperSnakeCaseName, err)
		}

		parsedValReflect := reflect.ValueOf(parsedVal)
		convertedVal := parsedValReflect.Convert(fieldVal.Type())
		if fieldVal.CanSet() {
			fieldVal.Set(convertedVal)
		} else {
			fmt.Println("cannot set value")
		}
		fmt.Printf("setting env var: %s = %s\n", upperSnakeCaseName, envVarVal)
	}

	return nil
}
