# Go bindings and CLI tool for the Skycoin hardware wallet

[![Build Status](https://travis-ci.com/skycoin/hardware-wallet-go.svg?branch=master)](https://travis-ci.com/skycoin/hardware-wallet-go)

## Installation

### Install golang

    https://github.com/golang/go/wiki/Ubuntu

### Install google protobuf

    sudo apt-get install protobuf-compiler python-protobuf golang-goprotobuf-dev
    go get -u github.com/golang/protobuf/proto/proto
    go get -u github.com/stretchr/testify/require

## Compile the protobuf project dependencies

    make -C vendor/nanopb/generator/proto/
    make -C protob/

## Usage

### Generate protobuf files

Only once each time the messages change:

    cd device-wallet/
    protoc -I ./protob  --go_out=./protob protob/messages.proto protob/types.proto protob/descriptor.proto 

### Run

    go test -run TestMain

## Wiki

More information in [the wiki](https://github.com/skycoin/hardware-wallet-go/wiki)
