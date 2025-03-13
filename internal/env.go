package internal

import "os"

// GetEnvVars loops through array and tries to fetch the value associated with each env var key.
// If there is no environment variable set, the value will be an empty string.
func GetEnvVars(envVarKeys []string) map[string]string {
	envVars := make(map[string]string)
	for _, envVarKey := range envVarKeys {
		// Just want to return everything, even if it's not available
		// will decide whether something is required or optional outside of this func
		envVars[envVarKey] = os.Getenv(envVarKey)
	}

	return envVars
}
