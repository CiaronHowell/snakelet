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

func Unmarshal(structPtr interface{}) error {
	val := reflect.ValueOf(structPtr)
	if val.Kind() != reflect.Pointer {
		return errors.New("struct passed needs to be a pointer")
	}

	tags := internal.ExtractTags(structPtr)
	fmt.Printf("tags: %v\n", tags)

	// Get the properties for a specific field
	for fieldIndex, tag := range tags {
		props, err := internal.ExtractProps(tag)
		if err != nil {
			return fmt.Errorf("failed to extract properties from tag: %w", err)
		}
		fmt.Printf("field index: %d, props: %v\n", fieldIndex, props)
	}

	envVarKeys := internal.ParseFieldNames(structPtr)
	fmt.Printf("parsed field names: %v\n", envVarKeys)
	envVars := internal.GetEnvVars(envVarKeys)

	fmt.Printf("env var key + val: %v\n", envVars)
	if err := internal.SetStructValues(structPtr, envVars); err != nil {
		return fmt.Errorf("failed to set struct fields: %w", err)
	}

	return nil
}
