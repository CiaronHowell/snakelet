package snakelet

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/ciaronhowell/snakelet/internal"
)

// Order of precedence (top to bottom, highest to lowest)
//  - Environment Variables
//  - Flags

func mergeEnvKeys(implicitEnvKeys map[int]string, customEnvKeys map[int]string) map[int]string {
	for fieldIndex, customEnvKey := range customEnvKeys {
		implicitEnvKeys[fieldIndex] = customEnvKey
	}

	return implicitEnvKeys
}

func Unmarshal(structPtr interface{}) error {
	val := reflect.ValueOf(structPtr)
	if val.Kind() != reflect.Pointer {
		return errors.New("struct passed needs to be a pointer")
	}

	tags, err := internal.ParseGoTags(structPtr)
	if err != nil {
		return fmt.Errorf("failed to parse go struct tags: %w", err)
	}
	fmt.Printf("go tags: %v\n", tags)

	customEnvKeys := make(map[int]string)
	// Get the properties for a specific field
	for fieldIndex, tag := range tags {
		props, err := internal.ExtractProps(tag)
		if err != nil {
			return fmt.Errorf("failed to extract properties from tag: %w", err)
		}

		fmt.Printf("field index: %d, props: %v\n", fieldIndex, props)

		if customEnvKey, ok := props["name"]; ok {
			customEnvKeys[fieldIndex] = customEnvKey
		}
	}

	// Merging implicit env var keys with explicit (custom) env var keys
	envVarKeys := mergeEnvKeys(internal.ParseFieldNames(structPtr), customEnvKeys)
	fmt.Printf("parsed and merged field names: %v\n", envVarKeys)

	envVars := internal.GetEnvVarValues(envVarKeys)

	fmt.Printf("env var key + val: %v\n", envVars)
	if err := internal.SetStructValues(structPtr, envVars); err != nil {
		return fmt.Errorf("failed to set struct fields: %w", err)
	}

	return nil
}
