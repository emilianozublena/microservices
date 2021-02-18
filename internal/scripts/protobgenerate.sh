#!/bin/bash

protoc pkg/customers/api/proto/v1/customers.proto --go_out=plugins=grpc:./pkg/grpc/v1