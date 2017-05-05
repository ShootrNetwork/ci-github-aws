#!/bin/bash
rm ci-github-aws
env GOOS=linux GOARCH=amd64 go build -v
md5 ci-github-aws
