package main

import (
	"fmt"

	"github.com/ciaronhowell/snakelet"
)

type AppConfig struct {
	FooBar string `snakelet:"required"`
	Bar    int
	Baz    bool
}

func main() {
	appCfg := AppConfig{}
	if err := snakelet.Unmarshal(&appCfg); err != nil {
		panic(err)
	}

	fmt.Printf("updated struct: %v", appCfg)
}
