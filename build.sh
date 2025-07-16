#!/bin/bash

GOOS=linux GOARCH=arm GOARM=5 go build -o bjornbot_arm
GOOS=linux GOARCH=amd64 go build -o bjornbot_amd64
