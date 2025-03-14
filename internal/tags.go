package internal

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type properties int

const (
	required properties = iota
	min
	max
)

func (properties) propKeys() []string {
	return []string{
		// "required", // INFO: Pointers could be used to show optional values????
		// "min",
		// "max",
		"name",
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

func ParseGoTags(structPtr interface{}) (map[int]string, error) {
	val := reflect.ValueOf(structPtr)
	if val.Kind() != reflect.Pointer {
		return nil, errors.New("struct passed needs to be a pointer")
	}

	// Get value from the pointer
	val = val.Elem()

	tags := make(map[int]string)
	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag.Get("snakelet")

		// If no snakelet tag is available, `tag` should be an empty string
		if tag == "" {
			continue
		}

		tags[i] = tag
	}

	return tags, nil
}
