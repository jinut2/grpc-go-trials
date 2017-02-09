# grpc-go-trials [![Go Report Card](https://goreportcard.com/badge/github.com/jinut2/grpc-go-trials)](https://goreportcard.com/report/github.com/jinut2/grpc-go-trials)

To generate go files for the protobuf
```
protoc --go_out=plugins=grpc:<dest> <location of proto file>
eg. protoc --go_out=plugins=grpc:. ./protos/user.proto  
```
