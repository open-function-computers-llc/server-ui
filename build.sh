#!/bin/bash

# let's build the web server!
cd httpd

# this command will build a staticly linked binary for 64 bit linux systems
# and place it in the dist folder
echo "Building linux binary..."
go build -o ../dist/ofco-web-ui
echo "done!"


export DB_TYPE=sqlite
export DB_DATABASE=test

export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=dbuser
export DB_PASSWORD=dbpass
export APP_PORT=9099
export ROUTE_PREFIX=tlioqwtisjdiauegliavaw4gjh

cd ../dist
killall ofco-web-ui || true
./ofco-web-ui &
