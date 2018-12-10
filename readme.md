### Compiling protobuf

    protoc -I cabinet/ cabinet/cabinet.proto --go_out=plugins=grpc:cabinet