// Package internal contains all decorrators
package internal

//go:generate gowrap gen -g -p go.zenithar.org/spotigraph/internal/services -i Tribe -t ../../../../../tools/templates/services/decorators/logger.txt -o logger.gen.go
//go:generate gowrap gen -g -p go.zenithar.org/spotigraph/internal/services -i Tribe -t ../../../../../tools/templates/services/decorators/opencensus.txt -o opencensus.gen.go
