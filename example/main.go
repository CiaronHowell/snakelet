package main

import (
	"fmt"

	"github.com/ciaronhowell/snakelet"
)

type AppConfig struct {
	Port       int
	ApiBaseUrl string
	RunProd    bool
	Namespace  string `snakelet:"name=K8S_NAMESPACE"`
}

func getConfig() (*AppConfig, error) {
	appCfg := AppConfig{}
	if err := snakelet.Unmarshal(&appCfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal app config: %w", err)
	}

	// Further custom code e.g., advanced checks, if needed...

	return &appCfg, nil
}

func main() {
	appCfg, err := getConfig()
	if err != nil {
		panic(err)
	}

	env := "development"
	if appCfg.RunProd {
		env = "production"
	}

	fmt.Printf("running as %s\n", env)
	fmt.Printf("running on port %d\n", appCfg.Port)
	fmt.Printf("using kubernetes namespace %s\n", appCfg.Namespace)
	fmt.Printf("sending POST request to %s/foo/bar\n", appCfg.ApiBaseUrl)
}
