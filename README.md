# go-spotigraph

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

### Dispatchers

- [x] gRPC Server
- [ ] GraphQL Server
- [ ] REST Api

### Decorators

- [x] Code generation
- [ ] Cache support
- [ ] Logger support
- [ ] OpenTracing support
- [ ] Prometheus support
- [ ] OpenCensus support
- [ ] Instrumented service providers for wire

### Tests

- [x] Unit tests template
- [ ] Integration tests
- [ ] E2E tests

### Deployment

- [x] Multi Stage Build for Docker
- [ ] Kubernetes manifest
- [ ] Test on K3S
