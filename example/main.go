package main

import (
	"fmt"

	"github.com/ciaronhowell/snakelet"
)

type Foo struct {
	FooBar string `snakelet:"required,min=6"`
	Bar string
	Baz string `snakelet:""`
}

func main() {
	test := Foo{}
	if err := snakelet.Unmarshal(&test); err != nil {
		panic(err)
	}

	fmt.Printf("updated struct: %v", test)

}
