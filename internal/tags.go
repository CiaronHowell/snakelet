package internal

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

type properties int

const (
	required properties = iota
	min
	max
)

func (properties) propKeys() []string {
	return []string{
		"required",
		"min",
		"max",
	}
}

func (p properties) String() string {
	return p.propKeys()[p]
}

func (p properties) IsProp(s string) bool {
	for _, prop := range p.propKeys() {
		if s == prop {
			return true
		}
	}

	return false
}

func ExtractProps(tags string) (map[string]string, error) {
	roughProps := strings.Split(tags, ",")

	props := make(map[string]string)
	for _, roughProp := range roughProps {
		kv := strings.SplitN(roughProp, "=", 2)

		if !properties.IsProp(0, kv[0]) {
			return nil, fmt.Errorf("invalid property has been passed: %s", kv[0])
		}

		props[kv[0]] = ""
		if len(kv) > 1 {
			props[kv[0]] = kv[1]
		}
	}

	return props, nil
}

func ExtractTags(foo interface{}) map[int]string {
	val := reflect.ValueOf(foo)
	if val.Kind() == reflect.Pointer {
		// Get value from the pointer
		val = val.Elem()
	}

	tags := make(map[int]string)
	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag.Get("snakelet")

		if tag == "" {
			continue
		}

		tags[i] = tag
	}

	return tags
}

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

func ParseFieldNames(foo interface{}) []string {
	val := reflect.ValueOf(foo)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	parsedFieldNames := []string{}
	for i := 0; i < val.NumField(); i++ {
		name := val.Type().Field(i).Name
		parsedFieldNames = append(parsedFieldNames, toUpperSnakeCase(name))
	}

	return parsedFieldNames
}
