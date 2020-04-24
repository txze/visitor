#!/usr/bin/env bash

set -e

SERVER_NAME=visitor

echo "--------- Start Publish Project ---------"

ROOT_DIR=$PUBLISH_DIR/$SERVER_NAME
rm -rf $ROOT_DIR
mkdir $ROOT_DIR

cp build/$SERVER_NAME.zip $ROOT_DIR/
cd $ROOT_DIR
unzip $SERVER_NAME.zip
