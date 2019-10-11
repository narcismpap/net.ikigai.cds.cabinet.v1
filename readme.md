# cds.cabinet
This is a prototype database developed against Apple's NewSQL open source database at https://www.foundationdb.org/

While this v1 is an early prototype of the system, subsequent versions have been successfully running large data clusters with nearly infinite horizontal scalability for our customers. This repository is kept for educational purposes.

## Documentation
You can read the technology specs for this prototype over at [V1 Final Docs](net.ikigai.cds.cabinet.v1.pdf).

### Compiling protobuf

    protoc --proto_path=./cabinet --go_out=plugins=grpc:./rpc cabinet/*.proto
    
### fdb

To build packages which use the go bindings, you will need to
set the following environment variables:

       CGO_CPPFLAGS="-I/Users/michael/Documents/Code/ikigai/net.ikigai.cds.cabinet.v1/src/github.com/apple/foundationdb/bindings/c"
       CGO_CFLAGS="-g -O2"
       CGO_LDFLAGS="-L/usr/local/lib"


Author: Narcis M. Pap, Singapore.
