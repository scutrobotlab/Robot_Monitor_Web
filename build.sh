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
go build
rm -f main.go
mv main.go.org main.go
#rm -rf WebPageBuild
