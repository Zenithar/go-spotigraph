proto:
	protoc \
		-I . \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/gogo/protobuf/protobuf \
		-I ${GOPATH}/src/github.com/lyft/protoc-gen-validate \
		--gogofast_out=\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:. \
		--validate_out="lang=gogo:." \
		pkg/protocol/v1/spotigraph/spotigraph.proto
