#!/usr/bin/env bash

set -e

SERVER_NAME=visitor

rm -rf build
mkdir build

echo "--------- Start Build Project ---------"
go build -o build/$SERVER_NAME

YamlFile="visitor.yaml"
if [[ -f "$YamlFile" ]]; then
  cp $YamlFile build/$YamlFile
fi

cd build

echo "--------- Start Package Project ---------"
zip $SERVER_NAME.zip -r ./*
