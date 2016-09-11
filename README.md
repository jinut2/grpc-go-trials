# grpc-go-trials

To generate go files for the protobuf
```
protoc --go_out=plugins=grpc:<dest> <location of proto file>
eg. protoc --go_out=plugins=grpc:. ./protos/user.proto  
```
