cd plugin1
go build -buildmode=plugin
cd ..
go build
./pluginloader1 ./plugin1/plugin1.so
