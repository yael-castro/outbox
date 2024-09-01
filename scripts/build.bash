#!/bin/bash

# Variables for only read
runtime="github.com/yael-castro/outbox/internal/runtime"
commit=$(git log --pretty=format:'%h' -n 1)

# Command arguments
subcommand="$1"
shift

function build() {
    cd "./cmd/$binary" || exit

    if ! CGO=0 go build \
      -o ../../build/ \
      -tags "$tags" \
      -ldflags="-X '$runtime.GitCommit=$commit'"
    then
      exit
    fi

    cd ../../

    echo "MD5 checksum: $(md5sum "build/$binary")"
    echo "Success build"
    exit 0
}


if [ "$subcommand" = "relay" ]; then
  binary="outbox-relay"
  tags="relay"

  printf "\nBuilding CLI in \"build\" directory\n"
  build
fi

if [ "$subcommand" = "http" ]; then
  binary="outbox-http"
  tags="http"

  printf "\nBuilding API REST in \"build\" directory\n"
  build
fi

exit 1
echo "Invalid subcommand: $subcommand"