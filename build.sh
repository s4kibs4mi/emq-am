#!/usr/bin/env bash

if [ $# == 0 ]; then
   echo "Application version expected."
   exit
fi

dockerPath="sakibsami"
appName="emq-am"
appVersion=$1

## Change as per your requirement
export GOOS=linux
export GOARCH=amd64

echo "Getting dependencies..."
dep ensure -v
echo "Done."
echo "Generating application binary..."
go build -v -o $(pwd)/bin/$appName
echo "Done"
echo "Creating docker image..."
docker build -t $dockerPath/$appName:$appVersion .
echo "Application is ready for deployment !!!"
echo "Docker Image : " $dockerPath/$appName:$appVersion
