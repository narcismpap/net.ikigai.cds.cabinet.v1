### Compiling protobuf

    protoc -I cabinet/ cabinet/*.proto --go_out=plugins=grpc:cabinet