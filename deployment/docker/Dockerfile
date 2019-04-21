## -------------------------------------------------------------------------------------------------

FROM golang:1.12 as builder

RUN set -eux; \
    apt-get update -y && \
    apt-get install -y apt-utils upx zip unzip;

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
    go run mage.go -d tools go:tools && \
    go run mage.go go:deps

# Build final target
RUN set -eux; \
    go run mage.go

# Compress binaries
RUN set -eux; \
    upx -9 bin/*

## -------------------------------------------------------------------------------------------------

FROM gcr.io/distroless/static:latest

COPY --from=builder /go/src/workspace/bin/spotigraph /usr/local/bin/spotigraph

ENTRYPOINT [ "/usr/local/bin/spotigraph" ]
