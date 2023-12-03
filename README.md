# golang-grpc
Go언어와 gRPC를 사용하며, [gRPC [Golang] Master Class: Build Modern API & Microservices](https://udemy.com/course/grpc-golang/)를 보면서 Go언어에서 gRPC를 사용하는 방법을 익힙니다.

## Installation

### gRPC
```shell
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin 
```

```shell
$ brew install protoc-gen-go
$ brew tap yoheimuta/protolint
$ brew install protolint
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

## Generate gRPC
```shell
$ protoc -Igreet/proto --go_out=. --go_opt=module=github.com/dev-hyunsang/grpc-go-course --go-grpc_out=. --go-grpc_opt=module=github.com/dev-hyunsang/grpc-go-course greet/proto/dummy.proto
```

## 참고한 글 및 문서
- **protoc-gen-go-grpc: program not found or is not executable**
  - [[Protocol Buffer/gRPC] protoc-gen-go-grpc: program not found or is not executable--go-grpc_out: protoc-gen-go-grpc: Plugin failed with status code 1.](https://darkstart.tistory.com/160)
  - [Error "protoc-gen-go: program not found or is not executable"](https://stackoverflow.com/questions/57700860/error-protoc-gen-go-program-not-found-or-is-not-executable)