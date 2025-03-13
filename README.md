# Snakelet

A lightweight, easy-to-use environment variable loader.

## TODO

### Critical

- [ ] Propagate parsing error to `Unmarshal` func
- [ ] Differentiate between unset env var and parsing errors
- [ ] Implement a way to set custom env var name
- [ ] Handle nested structs
  - Could use name of struct as a prefix for fields of the nested struct but variable names could get v long if we use the name of the struct. That said, if we go without a prefix, names could clash.

### Backlog

- [ ] Move logs from print statements to debug statements
- [ ] Ingest from .env file
- [ ] Add optional fields e.g., using pointers or follow similar practices as emitzero
- [ ] Add prefix option for all env vars
- [ ] Benchmark how quickly variables are loaded in vs other alternatives???
