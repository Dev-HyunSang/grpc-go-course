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

## Using Evans
### Installation
```shell
$ brew tap ktr0731/evans
$ brew install evans
```

### Command in gRPC
```shell
> pacakge calculator
> show message
+---------------+
|    MESSAGE    |
+---------------+
| AvgRequest    |
| AvgResponse   |
| MaxRequest    |
| MaxResponse   |
| PrimeRequest  |
| PrimeResponse |
| SqrtRequest   |
| SqrtResponse  |
| SumRequest    |
| SumResponse   |
+---------------+
> show service
+-------------------+--------+--------------+---------------+
|      SERVICE      |  RPC   | REQUEST TYPE | RESPONSE TYPE |
+-------------------+--------+--------------+---------------+
| CalculatorService | Sum    | SumRequest   | SumResponse   |
| CalculatorService | Primes | PrimeRequest | PrimeResponse |
| CalculatorService | Avg    | AvgRequest   | AvgResponse   |
| CalculatorService | Max    | MaxRequest   | MaxResponse   |
| CalculatorService | Sqrt   | SqrtRequest  | SqrtResponse  |
+-------------------+--------+--------------+---------------+

> service CalculatorService
> call Sum
first_number (TYPE_INT32) => 1
second_number (TYPE_INT32) => 1
{
  "result": 2
}
> call Primes
number (TYPE_INT64) => 42
{
  "result": "2"
}
{
  "result": "3"
}
{
  "result": "7"
}

> call Avg
number (TYPE_INT32) => 1
number (TYPE_INT32) => 2
number (TYPE_INT32) => 3
number (TYPE_INT32) => 4
number (TYPE_INT32) => 5
number (TYPE_INT32) => 6
number (TYPE_INT32) => 7
number (TYPE_INT32) => 8
number (TYPE_INT32) => 9
number (TYPE_INT32) => 10
number (TYPE_INT32) => 
{
  "result": 5.5
}
```


## 참고한 글 및 문서
- **protoc-gen-go-grpc: program not found or is not executable**
  - [[Protocol Buffer/gRPC] protoc-gen-go-grpc: program not found or is not executable--go-grpc_out: protoc-gen-go-grpc: Plugin failed with status code 1.](https://darkstart.tistory.com/160)
  - [Error "protoc-gen-go: program not found or is not executable"](https://stackoverflow.com/questions/57700860/error-protoc-gen-go-program-not-found-or-is-not-executable)