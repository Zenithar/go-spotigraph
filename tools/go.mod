module go.zenithar.org/spotigraph/tools

go 1.13

replace (
	github.com/envoyproxy/protoc-gen-validate => github.com/Zenithar/protoc-gen-validate v0.2.0-java.0.20190808132234-82cadd0ebcfe
	github.com/go-critic/go-critic => github.com/go-critic/go-critic v0.3.5-0.20190904082202-d79a9f0c64db
	github.com/golangci/errcheck => github.com/golangci/errcheck v0.0.0-20181223084120-ef45e06d44b6
	github.com/golangci/go-tools => github.com/golangci/go-tools v0.0.0-20190124090046-35a9f45a5db0
	github.com/golangci/gofmt => github.com/golangci/gofmt v0.0.0-20181222123516-0b8337e80d98
	github.com/golangci/gosec => github.com/golangci/gosec v0.0.0-20180901114220-8afd9cbb6cfb
	github.com/golangci/ineffassign => github.com/golangci/ineffassign v0.0.0-20180808204949-2ee8f2867dde
	github.com/golangci/lint-1 => github.com/golangci/lint-1 v0.0.0-20181222135242-d2cdd8c08219
	github.com/timakin/bodyclose => github.com/timakin/bodyclose v0.0.0-20190721030226-87058b9bfcec
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20190905072037-92dd089d5514
	mvdan.cc/unparam => mvdan.cc/unparam v0.0.0-20190720180237-d51796306d8f
)

require (
	github.com/99designs/gqlgen v0.9.3
	github.com/envoyproxy/protoc-gen-validate v0.0.0-00010101000000-000000000000
	github.com/fatih/color v1.7.0
	github.com/frapposelli/wwhrd v0.2.4
	github.com/gobuffalo/packr v1.30.1
	github.com/gogo/protobuf v1.3.0
	github.com/golang/mock v1.3.1
	github.com/golangci/golangci-lint v1.18.0
	github.com/google/wire v0.3.0
	github.com/hexdigest/gowrap v1.1.7
	github.com/iancoleman/strcase v0.0.0-20190422225806-e506e3ef7365 // indirect
	github.com/izumin5210/gex v0.5.1 // indirect
	github.com/magefile/mage v1.8.0
	github.com/sqs/goreturns v0.0.0-20181028201513-538ac6014518
	github.com/srikrsna/protoc-gen-mock v0.0.0-20190420084455-3bcb9cec43b1
	github.com/uber/prototool v1.8.1
	go.zenithar.org/protoc-gen-cobra v0.0.3
	golang.org/x/tools v0.0.0-20190909030654-5b82db07426d
	google.golang.org/appengine v1.4.0 // indirect
	gotest.tools/gotestsum v0.3.5
	mvdan.cc/gofumpt v0.0.0-20190729090447-96300e3d49fb
)
