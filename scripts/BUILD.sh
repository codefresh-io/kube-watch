#!/bin/sh

echo "Building kube-watch"
CGO_ENABLED=0 go build -v -o "./dist/bin/kube-watch" *.go