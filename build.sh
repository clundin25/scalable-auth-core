#! /bin/sh

set -ex

go build -buildmode=c-archive -o libscalable_auth.a main.go
