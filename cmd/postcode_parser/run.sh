#!/bin/bash

exec go build main.go parser.go &
wait
exec ./main
