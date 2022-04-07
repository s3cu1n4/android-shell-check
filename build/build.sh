#!/bin/bash

project_path=$(cd `dirname $0`; pwd)
project_name="${project_path##*/}"

cd $project_path

# macos 
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -o ../bin/android-shell-check ../main.go


# windows 
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a  -o ../bin/android-shell-check.exe ../main.go
echo "build successful"


