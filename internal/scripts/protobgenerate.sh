#!/bin/bash

protoc api/grpc/v1/drivers/core_service.proto --go_opt=paths=source_relative --go_out=plugins=grpc:./pkg/ --experimental_allow_proto3_optional
protoc api/grpc/v1/routes/core_service.proto --go_opt=paths=source_relative --go_out=plugins=grpc:./pkg/ --experimental_allow_proto3_optional


protoc api/proto/v1/drivers.proto --go_out=. --go-grpc_out=. --experimental_allow_proto3_optional