#!/bin/bash

set -e
set GOARCH=amd64
set GOOS=linux

go mod vendor
go mod verify

PROJECT_NAME="skydrive"
BINARY="skydrive_release"

OUTPUT_DIR=output
GOOS=linux

APP_NAME=${PROJECT_NAME}
APP_VERSION=$(git log -1 --oneline)
BUILD_VERSION=$(git log -1 --oneline)
BUILD_TIME=$(date "+%FT%T%z")
GIT_REVISION=$(git rev-parse --short HEAD)
GIT_BRANCH=$(git name-rev --name-only HEAD)
GO_VERSION=$(go version)
Debug=false

echo  ${OUTPUT_DIR}
echo  ${OUTPUT_DIR}/${BINARY}

CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -a -installsuffix cgo -v -mod=vendor \
-ldflags "-s -w -X 'main.AppName=${APP_NAME}' \
			-X 'main.AppVersion=${APP_VERSION}' \
			-X 'main.BuildVersion=${BUILD_VERSION}' \
			-X 'main.BuildTime=${BUILD_TIME}' \
			-X 'main.GitRevision=${GIT_REVISION}' \
			-X 'main.GitBranch=${GIT_BRANCH}' \
			-X 'main.GoVersion=${GO_VERSION}' \
			-X 'main.Debug=${Debug}'" \
-o ${OUTPUT_DIR}/${BINARY}