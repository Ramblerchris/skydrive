#!/bin/bash

set -e

PROJECT_NAME="skydrive"
BINARY="skydrive_debug"

OUTPUT_DIR=output
# GOOS=$(go env GOOS)
# GOOS=windows
GOOS=darwin
go mod vendor
go mod verify

APP_NAME=${PROJECT_NAME}
APP_VERSION=$(git log -1 --oneline)
BUILD_VERSION=$(git log -1 --oneline)
BUILD_TIME=$(date "+%FT%T%z")
GIT_REVISION=$(git rev-parse --short HEAD)
GIT_BRANCH=$(git name-rev --name-only HEAD)
GO_VERSION=$(go version)
Debug=true

echo  ${OUTPUT_DIR}
echo  ${OUTPUT_DIR}/${BINARY}

CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -v -mod=vendor \
-ldflags "-s  -X 'main.AppName=${APP_NAME}' \
			-X 'main.AppVersion=${APP_VERSION}' \
			-X 'main.BuildVersion=${BUILD_VERSION}' \
			-X 'main.BuildTime=${BUILD_TIME}' \
			-X 'main.GitRevision=${GIT_REVISION}' \
			-X 'main.GitBranch=${GIT_BRANCH}' \
			-X 'main.GoVersion=${GO_VERSION}' \
			-X 'main.Debug=${Debug}'" \
-o ${OUTPUT_DIR}/${BINARY}