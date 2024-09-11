#!/bin/bash

# Variables for only read
runtime="github.com/yael-castro/outbox/internal/runtime"
commit=$(git log --pretty=format:'%h' -n 1)

# Command arguments
subcommand="$1"
shift

function build() {
    cd "./cmd/$binary" || exit

    if ! go build \
      -o ../../build/ \
      -tags "$tags" \
      -ldflags "-X '$runtime.GitCommit=$commit'"
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
  CGO_ENABLED=1 build
fi

if [ "$subcommand" = "http" ]; then
  binary="outbox-http"
  tags="http"

  printf "\nBuilding API REST in \"build\" directory\n"
  CGO_ENABLED=0 build
fi

echo "Invalid subcommand: $subcommand"
exit 1