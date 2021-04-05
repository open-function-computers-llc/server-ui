#!/bin/bash

# let's build the web server!
cd httpd

# this command will build a staticly linked binary for 64 bit linux systems
# and place it in the dist folder
echo "Building linux binary..."
go build -o ../dist/ofco-web-ui
echo "done!"

cd ../
killall ofco-web-ui || true
dist/ofco-web-ui &
