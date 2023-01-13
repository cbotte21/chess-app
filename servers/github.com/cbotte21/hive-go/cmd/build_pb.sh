#!/bin/bash

protoc --go_out=$SERVER_DIR --proto_path=$PROJ_DIR/proto/ hive.proto
