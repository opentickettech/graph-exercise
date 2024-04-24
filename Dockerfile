# Simple docker image which takes a base golang version and prepares an environment which should be able to properly deploy golang applications
ARG GO_VERSION=1.22
FROM golang:$GO_VERSION-alpine

# Set goprivate to default values
ENV GOFLAGS="-mod=vendor" \
    CGO_ENABLED=0

# Install basic required packages.
# Remove caches to shrink image size.
RUN apk add git openssh-client make grep parallel \
    && go clean -cache -modcache

# TODO implement the actual building and running
