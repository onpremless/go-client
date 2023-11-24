#!/usr/bin/env bash
set -e

rm -rf client-tmp *.go go.mod go.sum docs README.md openapi.yaml
wget https://raw.githubusercontent.com/onpremless/api/master/openapi.yaml
openapi-generator-cli generate \
  -i ./openapi.yaml \
  -g go \
  -o client-tmp \
  --additional-properties=packageName="opless" \
  --git-user-id "onpremless" \
  --git-repo-id "go-client"
cd client-tmp && cp -r *.go go.mod go.sum docs README.md .. && cd ..
rm -rf client-tmp openapi.yaml
go mod tidy