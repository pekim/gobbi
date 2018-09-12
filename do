#!/bin/bash
##!/usr/bin/env bash
set -e
set -u

exampleDir=internal/example
tags="glib_2.10"

usage () {
    echo "Usage:"
    echo "  do -h                       Show this help"
    echo "  do build                    Generate libraries, and run tests"
    echo "  do generate                 Generate libraries"
    echo "  do test                     Run tests"
    echo "  do examples                 List the available examples"
    echo "  do example <example-name>   Run an example"

    exit
}

example () {
    exampleName=$1
    go run -tags $tags $exampleDir/$exampleName/main.go
}

build () {
    generate && test
}

generate () {
    go run internal/cmd/generate/main.go
}

test () {
    go test -tags $tags ./...
}

examples () {
  for example in $(find $exampleDir -name main.go -printf '%h\n') ; do
    basename $example
  done
}

processOptions () {
  while getopts ":h" opt; do
    case ${opt} in
      h )
        usage
        ;;
      \? )
        echo "Invalid Option: -$OPTARG" 1>&2
        echo ""
        usage
        ;;
    esac
  done
  shift $((OPTIND -1))

  if [ "$#" -lt 1 ]; then
    usage
  fi
}

processSubcommand () {
  subCommand=$1
  shift

  case "$subCommand" in
    build)
      build
      ;;
    generate)
      generate
      ;;
    test)
      test
      ;;
    examples)
      examples
      ;;
    example)
      example $@
      ;;
    *)
      echo "Invalid subcommand: $subCommand" 1>&2
      echo ""
      usage
      ;;
  esac
}

processOptions "$@"
processSubcommand "$@"
