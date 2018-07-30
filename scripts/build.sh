#! /bin/bash

protoc -I api/ \
-I${GOPATH}/src \
--go_out=plugins=grpc:api \
api/abstract_olt_api.proto

go build cmd/AbstractOLT/AbstractOLT.go
