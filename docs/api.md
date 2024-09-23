# go-zero api语法

api 领域特性语言包含语法版本，info 块，结构体声明，服务描述等几大块语法组成


## info 块
```go
syntax = "v1"

info(
    title: "订单管理"
    desc: "订单管理"
    author: "HUI"
    email: "example@gmail.com"
    version: 1.0
)

```


## 结构体声明
```go
type Orders {
	Id          string  `json:"id"`          //订单id
	Userid      uint64  `json:"userid"`      //用户id
	Shoppingid  int64   `json:"shoppingid"`  //收货信息表id
	Payment     float64 `json:"payment"`     //实际付款金额,单位是元,保留两位小数
	Paymenttype int8    `json:"paymenttype"` //支付类型,1-在线支付
	Postage     int     `json:"postage"`     //运费,单位是元
	Status      int16   `json:"status"`      //订单状态:0-已取消-10-未付款，20-已付款，30-待发货 40-待收货，50-交易成功，60-交易关闭
	CreateTime  int64   `json:"create_time"` //创建时间
	UpdateTime  int64   `json:"update_time"` //更新时间
}


type (
	OrderAddReq {
		ReceiveAddressId int64 `json:"receiveAddressId"` //用户收货地址表id
		Postage          int64 `json:"postage"`          //运费,单位是元
		Productid        int64 `json:"productid"`        //商品id
		Quantity         int64 `json:"quantity"`         //商品数量
	}
	OrderAddResp {
		Id string `json:"id"` //订单id
	}
)
```


## 服务描述
```go
@server (
    jwt:        Auth // 对当前 Foo 语法块下的所有路由，开启 jwt 认证，不需要则请删除此行
    prefix:     /v1 // 对当前 Foo 语法块下的所有路由，新增 /v1 路由前缀，不需要则请删除此行
    group:      g1 // 对当前 Foo 语法块下的所有路由，路由归并到 g1 目录下，不需要则请删除此行
    timeout:    3s // 对当前 Foo 语法块下的所有路由进行超时配置，不需要则请删除此行
    middleware: AuthInterceptor // 对当前 Foo 语法块下的所有路由添加中间件，不需要则请删除此行
    maxBytes:   1048576 // 对当前 Foo 语法块下的所有路由添加请求体大小控制，单位为 byte,goctl 版本 >= 1.5.0 才支持
)

service Foo {
    // 定义没有请求体和响应体的接口，如 ping
    @handler ping
    get /ping

    // 定义只有请求体的接口，如更新信息
    @handler update
    post /update (UpdateReq)

    // 定义只有响应体的结构，如获取全部信息列表
    @handler list
    get /list returns ([]ListItem)

    // 定义有结构体和响应体的接口，如登录
    @handler login
    post /login (LoginReq) returns (LoginResp)

    // 定义表单请求
    @handler formExample
    post /form/example (FormExampleReq)

    // 定义 path 参数
    @handler pathExample
    get /path/example/:id (PathExampleReq) returns (PathExampleResp)
}
```

## 导入模块
```go
import (
	"apis/userReceiveAddress.api"
	"apis/user.api"
	"apis/usercollection.api"
	"apis/order.api"
)
```