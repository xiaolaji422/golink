#!/bin/sh
protoc --go_out=plugins=grpc:. send_msg.proto