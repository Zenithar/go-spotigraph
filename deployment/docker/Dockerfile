FROM golang:1.12 as builder

# build golang environment
ENV PROTOC_VER 3.7.0

RUN apt-get update -y && \
    apt-get install -y apt-utils zip unzip; \
    wget -q -P /tmp/temp/ https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VER}/protoc-${PROTOC_VER}-linux-x86_64.zip && \
    cd /usr && unzip /tmp/temp/protoc-${PROTOC_VER}-linux-x86_64.zip && \
    rm -f /tmp/temp/protoc-${PROTOC_VER}-linux-x86_64.zip;

RUN adduser -D -u 1000 golang

WORKDIR $GOPATH/src/workspace

USER golang


FROM gcr.io/distroless/static:latest

COPY --from=builder source /usr/local/bin/spotigraph
