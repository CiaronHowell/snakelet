package snakelet

import (
	"fmt"

	"github.com/ciaronhowell/snakelet/internal"
)

// Order of precedence (top to bottom, highest to lowest)
//  - Environment Variables
//  - Flags

func Unmarshal(structPtr interface{}) error {
	// TODO: Make sure we have address rather than obj
	tags := internal.ExtractTags(structPtr)
	fmt.Printf("tags: %v\n", tags)

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
	internal.SetStructValues(structPtr, envVars)

	return nil
}
