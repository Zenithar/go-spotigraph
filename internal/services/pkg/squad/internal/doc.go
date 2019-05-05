// Package internal contains all decorrators
package internal

//go:generate gowrap gen -g -p go.zenithar.org/spotigraph/internal/services -i Squad -t ../../../../../tools/templates/services/decorators/logger.gotmpl -o logger.gen.go
//go:generate gowrap gen -g -p go.zenithar.org/spotigraph/internal/services -i Squad -t ../../../../../tools/templates/services/decorators/opencensus.gotmpl -o opencensus.gen.go
