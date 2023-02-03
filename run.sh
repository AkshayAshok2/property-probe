#!/bin/bash

cd client
npm run start &

cd ../api
go run httpd/main.go