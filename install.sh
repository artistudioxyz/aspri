#!/bin/bash
go build main.go
chmod 777 main
yes | sudo cp main $1
rm -f main