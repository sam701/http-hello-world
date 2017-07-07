#!/bin/bash

(
  cd ..

  docker run --name http-hello-world-build --rm \
    -v "$PWD":/go/src/github.com/sam701/http-hello-world -w /go/src/github.com/sam701/http-hello-world \
    sam701/golang-and-dependencies \
    sh -c 'go build'
    
  mv http-hello-world docker/
)

docker build -t sam701/http-hello-world .
