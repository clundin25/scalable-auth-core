#! /bin/sh

set -ex

go build -buildmode=c-archive -o libscalable-auth.a main.go
