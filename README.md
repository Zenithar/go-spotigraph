# go-spotigraph

[![Go Report Card](https://goreportcard.com/badge/github.com/Zenithar/go-spotigraph)](https://goreportcard.com/report/github.com/Zenithar/go-spotigraph)
[![CircleCI](https://circleci.com/gh/Zenithar/go-spotigraph.svg?style=svg)](https://circleci.com/gh/Zenithar/go-spotigraph)
[![LICENSE](https://img.shields.io/github/license/Zenithar/go-spotigraph.svg)](https://github.com/Zenithar/go-spotigraph/blob/master/LICENSE)

Spotify Agile model mapping microservice

## Build

```sh
> go run .circleci/mage.go
```

## Todo

### Project

- [x] Golang project best practice
- [x] Go modules
- [x] Go tools vendoring
- [x] Magefile migration
- [x] Feature flags support
- [x] Wire code generation
- [x] CircleCI integration
- [ ] Documentation

### Services

- [x] Identity management
- [x] Squad management
- [x] Chapter management
- [x] Guild management
- [x] Tribe management
- [ ] Graph query
- [ ] External identity resolver (LDAP, etc.)

### Code generation

- [x] Protocol Validator generation
- [x] gRpc Client code generation
- [x] Decorator generation
- [x] CLI command generation from protobuf
- [x] Local or Remote service abstraction (direct vs gRPC call)

### Database

- [x] RethinkDB support
- [x] MongoDB support
- [x] PostgreSQL support
- [x] PostgreSQL json column support
- [ ] BoltDB backend study
- [ ] DGraph backend study
- [ ] DynamoDB backend study

### Dispatchers

- [x] gRPC Server
- [x] GraphQL Server
- [x] REST Api

### Decorators

- [x] Code generation
- [x] Cache support
- [x] Logger support
- [x] Tracing support
- [x] Metric support
- [x] OpenCensus support
- [x] Instrumented service providers for wire
- [ ] Circuit breaker
- [ ] Authorization

### Tests

- [x] Unit tests template
- [ ] Integration tests
- [ ] E2E tests

### Deployment

- [x] Multi Stage Build for Docker
- [ ] Kubernetes manifest
- [ ] Test on K3S
