package internal

import "os"

func GetEnvVars(envVarKeys []string) map[string]string {
	envVars := make(map[string]string)
	for _, envVarKey := range envVarKeys {
		// Just want to return everything, even if it's not available
		// will decide whether something is required or optional outside of this func
		envVars[envVarKey] = os.Getenv(envVarKey)
	}

	return envVars
}
