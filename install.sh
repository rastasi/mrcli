#!/bin/sh
git clone https://github.com/rastasi/mrcli
cd mrcli
go get all
go build
sudo mv ./mrcli /usr/local/bin
cd ..
rm -rf mrcli