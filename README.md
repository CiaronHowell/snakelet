# Snakelet

**An opinionated, lightweight, easy-to-use environment variable loader.**

## Usage

### Install

```shell
go get github.com/ciaronhowell/snakelet
```

### Quick Start

You can use the example `main.go` file [here](example/main.go).

```go
type Example struct {
  Foo int
  Bar string `snakelet:"name=BAZ"`
}

func main() {
  exampleStruct := Example{}
  if err := snakelet.Unmarshal(&exampleStruct); err != nil {
    panic(err)
  }

  fmt.Printf("foo: %d, bar: %s\n", exampleStruct.Foo, exampleStruct.Bar)
}
```
