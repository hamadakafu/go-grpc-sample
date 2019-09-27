#!/bin/bash

protoc --proto_path proto --go_out=plugins=grpc:./pb sample.proto
