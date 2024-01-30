#!/bin/bash
cp .astric/config/server.env .env

sed -i 's/"ip de mysql"/'$DBHOST'/' .env
sed -i 's/"clave mysql"/'$DBCLAVE'/' .env
sed -i 's/"usuario mysql"/'$DBUSER'/' .env

