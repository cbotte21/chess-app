#!/bin/bash

protoc --go_out=$SERVER_DIR --go_opt=paths=source_relative --go-grpc_out=$SERVER_DIR --go-grpc_opt=paths=source_relative --proto_path=$PROJ_DIR/proto/ hive.proto
