## -------------------------------------------------------------------------------------------------

FROM golang:1.12 as builder

# build golang environment
ENV PROTOC_VER 3.7.0

RUN set -eux; \
    apt-get update -y && \
    apt-get install -y apt-utils upx zip unzip;

# Install protoc
RUN set -eux; \
    wget -q -P /tmp/temp/ https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VER}/protoc-${PROTOC_VER}-linux-x86_64.zip && \
    cd /usr && unzip /tmp/temp/protoc-${PROTOC_VER}-linux-x86_64.zip && \
    chmod o+x /usr/bin/protoc && \
    rm -f /tmp/temp/protoc-${PROTOC_VER}-linux-x86_64.zip;

# Drop root privileges to build
RUN adduser --disabled-password --gecos "" -u 1000 golang && \
    mkdir -p $GOPATH/src/workspace && \
    chown -R golang:golang $GOPATH/src/workspace;

# Force go modules
ENV GO111MODULE=on

WORKDIR $GOPATH/src/workspace

USER golang
COPY --chown=golang:golang . .

# Install dependencies
RUN set -eux; \
    go run .circleci/mage.go go:deps && \
    gex --build

# Add locally installed tools to path
ENV PATH $PWD/bin:$PATH

# Build final target
RUN set -eux; \
    go run .circleci/mage.go

## -------------------------------------------------------------------------------------------------

FROM gcr.io/distroless/static:latest

COPY --from=builder /src/go/bin/spotigraph /usr/local/bin/spotigraph

ENTRYPOINT [ "/usr/local/bin/spotigraph" ]
