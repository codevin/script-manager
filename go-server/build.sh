#!/bin/bash

# Compile for x86, normal compilation.
echo "Building cmdline"
env GOOS=linux go build -ldflags="-extldflags=-static -X server.minversion=`date -u +.%Y%m%d.%H%M%S`" -o cmdline cmdline.go 

echo "Linking update to cmdline. It is same app, but used for sending update rather than new widget request."
ln -sf cmdline update

echo "Building script-manager"
env GOOS=linux go build -ldflags="-extldflags=-static -X server.minversion=`date -u +.%Y%m%d.%H%M%S`" -o script-manager main.go

if [ x$1 == x"run" ]
then
echo "Running script-manager"
./script-manager

else
echo "Invoke \"./build.sh run\" to run script-manager"

fi


