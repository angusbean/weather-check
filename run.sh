#!/bin/bash
redis-server &
go build -o weather-check *.go && ./weather-check