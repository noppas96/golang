# GO 1.19

GO ENV

- GO111MODULE = "[auto,on,off]" default is not set = auto
- GOROOT =  store binary and basic package
- GOPATH = default store external package import from anywhere, Dir under gopath can't use go mod 

Create GO Module
- go mod init github.com/gitUser/gitRepo
    - go.mod similar to package.json nodejs

Dependencies MGMT
- go mod vendor 
    - vendor similar to node_modules nodejs

Go Packages
- [CSI](https://pkg.go.dev/github.com/container-storage-interface/spec@v1.6.0/lib/go/csi#ControllerGetCapabilitiesRequest)

# Protobuf

```
syntax = "proto3";

package pingpong;

option go_package = "./;pb";

message Ping {
  // Field number 1-15 use 1 byte, while field 16th - 2047th use 2 bytes
  // So, the first 15 fields should be reserved for fields that are used oftenly
  int32 id = 1;
  string message = 2;
}

message Pong {
  int32 id = 1;
  string message = 2;
}

service PingPong {

  rpc StartPing (Ping) returns (Pong);
}
```
## Example Generate Code

```
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
         routeguide/route_guide.proto
```

```
protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb
```
Remark
```
brew install protoc-gen-go
brew install protoc-gen-go-grpc
```
# gRPC
# Reference
-   [proto-gen-go issue](https://github.com/techschool/pcbook-go/issues/3)

