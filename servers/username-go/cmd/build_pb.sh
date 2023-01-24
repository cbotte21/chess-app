#!/bin/bash

protoc --go_out=$USERNAME_DIR/pb --go_opt=paths=source_relative --go-grpc_out=$USERNAME_DIR/pb --go-grpc_opt=paths=source_relative --proto_path=$PROJ_DIR/proto/ username.proto
