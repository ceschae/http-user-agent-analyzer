## What is it

An API

## What should it do

analyze an HTTP user agent and make a decision as to whether to block or allow it
- if the user agent header is from a Safari browser, it should return a decision to block the request
- if it is from a Firefox browser, it should allow the request

## How should it do it

- a gRPC client and server in Rust or Go
  - implements a single endpoint that receives the user agent string, analyzes it, and returns the result
- Create a terminal CLI which uses the client to call the API
- Include appropriate tests and documentation
- Present the completed project within a Git repo with a README

## Helpful commands
- `protoc --go_out=./ ./httpanalyzer/http_analyzer.proto`
- `protoc --go-grpc_out=./ ./httpanalyzer/http_analyzer.proto`

## Questions
- 'single endpoint' is a requirement i'm not sure how to complete. my confusion comes from not knowing how to _enforce_ single endpoint, because my understanding is you could clone the program to any number of devices and not be able to know for sure it's a _single_ endpoint. it's likely i'm overthinking this phrase, but wanted to call it out regardless.
- behavior not defined for non-Safari and non-Firefox browser user agent headers. what should we do in this case? (defaulting to blocking everything not explicitly allowed by spec)

## Outstanding To-Dos
- the TUI experience is unwritten
- `README`
- testing