#!/bin/sh
cp -r WebPage WebPageBuild
a=$(awk -F "\""  '/http(.+?)/{print $2}' WebPageBuild/index.html)
wget -q $a -P WebPageBuild
sed -i 's/https[a-zA-Z0-9\:\/\.@\-]*\///g' WebPageBuild/index.html
go get github.com/go-bindata/go-bindata/...
go get github.com/elazarl/go-bindata-assetfs/...
go-bindata -fs -prefix "WebPageBuild/" WebPageBuild/
go-bindata-assetfs WebPageBuild/...
go build
#rm -rf WebPageBuild
