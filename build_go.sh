#!/bin/sh
cp main.go main.go.org
sed -i 's/http.Dir(\".\/frontend\/dist\/\")/assetFS()/g' main.go
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go get github.com/go-bindata/go-bindata/...
go get github.com/elazarl/go-bindata-assetfs/...
go-bindata-assetfs -prefix frontend frontend/dist/...
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o Robot_Monitor_Web_linux
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -o Robot_Monitor_Web_mac
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o Robot_Monitor_Web_windows.exe
upx -q -9 Robot_Monitor_Web_linux Robot_Monitor_Web_mac Robot_Monitor_Web_windows.exe
zip Robot_Monitor_Web_linux.zip Robot_Monitor_Web_linux
zip Robot_Monitor_Web_mac.zip Robot_Monitor_Web_mac
zip Robot_Monitor_Web_windows.zip Robot_Monitor_Web_windows.exe
rm -f main.go
mv main.go.org main.go
