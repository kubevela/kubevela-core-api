#!/bin/bash

rm -r apis
cp -r ../kubevela/apis .

rm -r pkg/oam
cp -r ../kubevela/pkg/oam pkg/

find . -type f -name "*.go" -print0 | xargs -0 sed -i '' 's|github.com/oam-dev/kubevela/|github.com/oam-dev/kubevela-core-api/|g'

go build test/main.go