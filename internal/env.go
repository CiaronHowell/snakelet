package internal

import "os"

// GetEnvVarValues loops through map and tries to fetch the value associated with each env var key.
// If there is no environment variable set, the value will be an empty string.
//
// Returns a map containing the field index and the associated value.
func GetEnvVarValues(envVarKeys map[int]string) map[int]string {
	envVars := make(map[int]string)
	for fieldIndex, envVarKey := range envVarKeys {
		// Just want to return everything, even if it's not available
		// will decide whether something is required or optional outside of this func
		envVars[fieldIndex] = os.Getenv(envVarKey)
	}

	return envVars
}
