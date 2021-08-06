PATH="${PATH}:${HOME}/go/bin" protoc --go-grpc_out=. *.proto

PATH="${PATH}:${HOME}/go/bin" protoc --go_out=. *.proto