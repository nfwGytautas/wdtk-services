#/bin/bash

# This is a utility script for building various in-buit WDTK services

mkdir bin > /dev/null 2>&1

# Gateway
echo Building gateway

cd gateway/

echo Running go tidy
go mod tidy

echo Downloading dependencies
go get ./

echo Building
go build -o ../bin/APIGateway .

cd .. # return to root
