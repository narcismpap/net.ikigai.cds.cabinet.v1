### Compiling protobuf

    protoc --proto_path=./cabinet --go_out=plugins=grpc:./rpc cabinet/*.proto
    
### fdb

To build packages which use the go bindings, you will need to
set the following environment variables:

       CGO_CPPFLAGS="-I/Users/michael/Documents/Code/ikigai/net.ikigai.cds.cabinet.v1/src/github.com/apple/foundationdb/bindings/c"
       CGO_CFLAGS="-g -O2"
       CGO_LDFLAGS="-L/usr/local/lib"