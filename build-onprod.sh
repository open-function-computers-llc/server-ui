#!/bin/bash

# let's build the web server!

export DB_TYPE=sqlite
export DB_DATABASE=test
export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=dbuser
export DB_PASSWORD=dbpass
export APP_PORT=9099
export ROUTE_PREFIX=tlioqwtisjdiauegliavaw4gjh
export TOOLSBASEDIR=/root/ofco-deploy/modx-apache/

dist/ofco-web-ui
