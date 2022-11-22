# Go Basic

## Variables Declaration

Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type. Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

### Basic go data type

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

//Variables declared without an explicit initial value are given their zero value.

//The zero value is:

// 0     for numeric types,
// false for the boolean type, and
// ""    (the empty string) for strings.
```
Remark : The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.

### Constants
Constants are declared like variables, but with the const keyword.
Constants can be character, string, boolean, or numeric values.
Constants cannot be declared using the := syntax.

### Type conversions

The expression T(v) converts the value v to the type T.
When the right hand side of the declaration is typed, the new variable is of that same type, But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

i := 42
f := float64(i)
u := uint(f)
```

## Loop
the init statement: executed before the first iteration
the condition expression: evaluated before every iteration
the post statement: executed at the end of every iteration
```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
    fmt.Println(i)   
	}
	fmt.Println(sum)
}

```




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

