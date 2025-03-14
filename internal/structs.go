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

// toUpperSnakeCase takes a CamelCase string and converts it to UPPER_SNAKE_CASE.
// Underscores are added per hump in the string e.g., CamelCase turns into CAMEL_CASE.
func toUpperSnakeCase(s string) string {
	upperSnakeCase := []rune{}
	for i, c := range s {
		if i != 0 && unicode.IsUpper(c) {
			upperSnakeCase = append(upperSnakeCase, '_')
		}

		upperSnakeCase = append(upperSnakeCase, unicode.ToUpper(c))
	}

	return string(upperSnakeCase)
}

func SetStructValues(structPtr interface{}, envVarValues map[int]string) error {
	val := reflect.ValueOf(structPtr)
	if val.Kind() != reflect.Pointer {
		return errors.New("struct passed is not a pointer")
	}
	val = val.Elem()

	// Loop through each field in struct
	for fieldIndex := 0; fieldIndex < val.NumField(); fieldIndex++ {
		name := val.Type().Field(fieldIndex).Name
		if !isExportedField(name) {
			fmt.Printf("skipping unexported field in struct: %s\n", name)
			continue
		}

		upperSnakeCaseName := toUpperSnakeCase(name)
		fmt.Printf("checking for env var: %s\n", upperSnakeCaseName)
		envVarVal, ok := envVarValues[fieldIndex]
		if !ok {
			continue
		}

		field := val.Field(fieldIndex)
		parsedVal, err := parseEnvVarValue(envVarVal, field.Kind())
		if err != nil {
			return fmt.Errorf("failed to parse value for %s: %w", upperSnakeCaseName, err)
		}

		parsedValReflect := reflect.ValueOf(parsedVal)
		convertedVal := parsedValReflect.Convert(field.Type())
		if field.CanSet() {
			field.Set(convertedVal)
		} else {
			fmt.Println("cannot set value")
		}
		fmt.Printf("setting env var: %s = %s\n", upperSnakeCaseName, envVarVal)
	}

	return nil
}

func ParseFieldNames(structPtr interface{}) (envVarKeys map[int]string) {
	val := reflect.ValueOf(structPtr)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	parsedFieldNames := make(map[int]string)
	for fieldIndex := 0; fieldIndex < val.NumField(); fieldIndex++ {
		fieldName := val.Type().Field(fieldIndex).Name
		parsedFieldNames[fieldIndex] = toUpperSnakeCase(fieldName)
	}

	return parsedFieldNames
}
