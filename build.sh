#!/bin/sh
cp -r WebPage WebPageBuild
cp main.go main.go.org
sed -i 's/http.Dir(\".\/WebPage\/\")/assetFS()/g' main.go
wget -q $(awk -F "\""  '/http(.+?)/{print $2}' WebPageBuild/index.html) -P WebPageBuild
sed -i 's/https[a-zA-Z0-9\:\/\.@\-]*\///g' WebPageBuild/index.html
go get github.com/go-bindata/go-bindata/...
go get github.com/elazarl/go-bindata-assetfs/...
go-bindata -fs -prefix "WebPageBuild/" WebPageBuild/
go-bindata-assetfs WebPageBuild/...
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o Robot_Monitor_Web_linux
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -o Robot_Monitor_Web_mac
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o Robot_Monitor_Web_windows.exe
zip Robot_Monitor_Web_linux.zip Robot_Monitor_Web_linux
zip Robot_Monitor_Web_mac.zip Robot_Monitor_Web_mac
zip Robot_Monitor_Web_windows.zip Robot_Monitor_Web_windows.exe
rm -f main.go
mv main.go.org main.go
#rm -rf WebPageBuild
