SHELL := /bin/bash

GO111MODULE = on

proto:
	protoc --gogo_out=./pb/ --proto_path=./pb/ sofer.proto
.PHONY: proto

