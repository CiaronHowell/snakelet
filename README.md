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

## TODO

### Critical

- [x] Propagate parsing error to `Unmarshal` func
- [ ] Differentiate between unset env var and parsing errors
- [x] Implement a way to set custom env var name
- [ ] Handle nested structs
  - Could use name of struct as a prefix for fields of the nested struct but variable names could get v long if we use the name of the struct. That said, if we go without a prefix, names could clash.

### Backlog

- [ ] Move logs from print statements to debug statements
- [ ] Ingest from .env file
- [ ] Add optional fields e.g., using pointers or follow similar practices as emitzero
- [ ] Add prefix option for all env vars
- [ ] Benchmark how quickly variables are loaded in vs other alternatives???
