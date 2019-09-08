# Copyright 2019 Thibault NORMAND
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Arguments
ARG BUILD_DATE
ARG VERSION
ARG VCS_REF

## -------------------------------------------------------------------------------------------------

FROM golang:1.13 as builder

RUN set -eux; \
    apt-get update -y && \
    apt-get install -y apt-utils upx zip unzip;

# Create a non-root privilege account to build
RUN adduser --disabled-password --gecos "" -u 1000 golang && \
    mkdir -p $GOPATH/src/workspace && \
    chown -R golang:golang $GOPATH/src/workspace;

# Force go modules
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org/

WORKDIR $GOPATH/src/workspace

# Prepare an unprivilegied user for run
RUN set -eux; \
    echo 'nobody:x:65534:65534:nobody:/:' > /tmp/passwd && \
    echo 'nobody:x:65534:' > /tmp/group && \
    mkdir /tmp/.config && \
    chown 65534:65534 /tmp/.config

# Drop privileges to build
USER golang
COPY --chown=golang:golang mage.go .
COPY --chown=golang:golang tools tools/

# Install tools
RUN set -eux; \
    go run mage.go -d tools 

# Copy project go module
COPY --chown=golang:golang . .

# Build final target
RUN set -eux; \
    go run mage.go build

# Compress binaries
RUN set -eux; \
    upx -9 bin/* && \
    chmod +x bin/*

## -------------------------------------------------------------------------------------------------

FROM gcr.io/distroless/static:latest

# Arguments
ARG BUILD_DATE
ARG VERSION
ARG VCS_REF

# Metadata
LABEL \
    org.label-schema.build-date=$BUILD_DATE \
    org.label-schema.name="Spotigraph" \
    org.label-schema.description="Spotify Agile model mapping microservice" \
    org.label-schema.url="https://go.zenithar.org/spotigraph" \
    org.label-schema.vcs-url="https://github.com/Zenithar/go-spotigraph.git" \
    org.label-schema.vcs-ref=$VCS_REF \
    org.label-schema.vendor="Thibault NORMAND" \
    org.label-schema.version=$VERSION \
    org.label-schema.schema-version="1.0" \
    org.zenithar.licence="MIT"

COPY --from=builder /go/src/workspace/bin/spotigraph /usr/bin/spotigraph
COPY --from=builder /tmp/group /tmp/passwd /etc/
COPY --from=builder --chown=65534:65534 /tmp/.config /

USER nobody:nobody
WORKDIR /

ENTRYPOINT [ "/usr/bin/spotigraph" ]
CMD ["--help"]

