Protobuf

**what is protobuf**

Protobuf（Protocol Buffers）是一种由Google开发的、用于序列化数据的语言中立、平台无关的高效机制。它通过定义简单的 `.proto` 文件来描述数据结构，并通过编译器生成相应的代码，方便程序在不同语言中处理这些数据格式

```go
// 声明 proto 语法版本，固定值
syntax = "proto3";

// 包名
package proto;

// golang 代码后的包名
option go_package = "/pb";

// 请求体
message SayHelloRequest {
}


// 响应体
message SayHelloResponse {

}


/*
字段标签的作用：
    唯一标识字段：每个字段都必须有一个唯一的标签，Protobuf 使用这些标签来区分不同的字段，
    而不是字段名称。因此，在序列化数据时，Protobuf会根据这些数字来识别和组织数据。

    编码和解码：当 Protobuf 序列化时，数据会被压缩成一个紧凑的二进制格式，
    标签值成为数据包的一部分。解码时，Protobuf 通过标签找到对应的字段。

    版本兼容性：在更新 .proto 文件时，可以添加新字段，而不影响旧版代码的兼容性。
    只要为新字段指定唯一的标签，不同版本的消息就可以相互理解和解析。
*/
message Hello {
    // 字段类型 字段名称 = 字段标签;
    string name = 1;
    string country = 2;
}

message SayHello {
    string name = 1;
}

// 定义HelloService服务
service SaySomethingService {
    rpc SayHello (SayHelloRequest) returns (SayHelloResponse);
}
```

**Generate Code**

```shell
protoc --proto_path=proto --go_out=pb --go-grpc_out=. \
--go_opt=paths=source_relative proto/*.proto
```

**go-zero generate code**
[go-zero goctl rpc](https://go-zero.dev/docs/tutorials/cli/rpc)

```shell
goctl rpc protoc rpc.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```

## grpc-go

[grpc_go quick start](https://grpc.io/docs/languages/go/quickstart/)