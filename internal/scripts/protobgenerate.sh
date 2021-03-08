#!/bin/bash

protoc api/grpc/v1/routes/core_service.proto --go_opt=paths=source_relative --go_out=plugins=grpc:./ --experimental_allow_proto3_optional