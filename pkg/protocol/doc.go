package protocol

//go:generate protoc -I . -I ${GOPATH}/src -I ${GOPATH}/src/github.com/lyft/protoc-gen-validate \
// --gogofast_out=":." --validate_out="lang=gogo:." v1/spotigraph/spotigraph.proto
