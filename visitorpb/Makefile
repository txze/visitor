#!/usr/bin/env bash

# Protocol Buffers for Go with Gadgets
#
# Copyright (c) 2013, The GoGo Authors. All rights reserved.
# http://github.com/gogo/protobuf


gogo:
	protoc \
	-I. \
	-I ${GOPATH}/src \
	--gogo_out=plugins=grpc:. *.proto

go:
	protoc \
	-I. \
	-I ${GOPATH}/src \
	--go_out=plugins=grpc:. *.proto