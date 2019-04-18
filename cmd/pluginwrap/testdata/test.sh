#!/bin/sh
echo "build pluginscanner ..."
go build -o pluginwrap ..
echo "scanning ./foo/"
./pluginwrap ./foo/
echo "run 'go vet -v' ..."
go vet -v
echo "build plugin ..."
go build -o foo.so -buildmode=plugin ./foo
echo "run go test ..."
#go test -bench=.
go test
