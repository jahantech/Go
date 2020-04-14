# Service A GRPC


## Protoc Compiler
https://github.com/google/protobuf/releases

## GRPC
```go get -u google.golang.org/grpc```

## Protoc Go Plugin
```go get -u github.com/golang/protobuf/protoc-gen-go```

## Generator go.mod
``` go mod init service_a_grpc```

## Use protoc compiler to generate code files from proto file
```cd api/ && protoc -I accumulator  accumulator/accumulator.proto --go_out=plugins=grpc:accumulator```

