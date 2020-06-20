#!/bin/sh
cp main.go main.go.org
sed -i 's/http.Dir(\".\/frontend\/dist\/\")/assetFS()/g' main.go
go get github.com/go-bindata/go-bindata/...
go get github.com/elazarl/go-bindata-assetfs/...
go-bindata-assetfs -prefix frontend frontend/dist/...
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o Robot_Monitor_Web_linux
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -o Robot_Monitor_Web_mac
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o Robot_Monitor_Web_windows.exe
zip Robot_Monitor_Web_linux.zip Robot_Monitor_Web_linux
zip Robot_Monitor_Web_mac.zip Robot_Monitor_Web_mac
zip Robot_Monitor_Web_windows.zip Robot_Monitor_Web_windows.exe
rm -f main.go
mv main.go.org main.go
