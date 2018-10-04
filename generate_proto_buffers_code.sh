#!/usr/bin/env bash
protoc pb/problems.proto --go_out=plugins=grpc:.