// +build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// GenerateServiceProtocol call protobuf compilet to generate service protocol
func GenerateServiceProtocol() error {
    return sh.Run(
        "protoc",
        "-I", ".",
        "-I", "${GOPATH}/src",
        "-I", "${GOPATH}/src/github.com/gogo/protobuf/protobuf",
        "--gogo_out", "Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:.",
        "--validate_out", "lang=gogo:.",
        "pkg/protocol/v1/spotigraph/spotigraph.proto",
    )
}

// GenerateGRPC call protobuf compiler to generate gRPC stubs
func GenerateGRPC() error {
    return sh.Run(
        "protoc",
        "-I", ".",
        "-I", "${GOPATH}/src",
        "-I", "${GOPATH}/src/github.com/gogo/protobuf/protobuf",
        "--gogo_out", "plugins=grpc,Mpkg/protocol/v1/spotigraph/spotigraph.proto=go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph:.",
        "pkg/grpc/v1/spotigraph/pb/spotigraph.proto",
    )
}
