#!/bin/bash

cd client
npm install
npm run start &

cd ../api
go run httpd/main.go
