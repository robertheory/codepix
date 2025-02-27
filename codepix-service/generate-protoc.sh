#!/bin/bash

protoc --go_out=app/grpc/pb --go_opt=paths=source_relative --go-grpc_out=app/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=app/grpc/protofiles app/grpc/protofiles/*.proto
