## go zero
[go-zero api语法](https://github.com/Junhiee/lebron/blob/main/docs/api.md)

[go-zero rpc服务](https://github.com/Junhiee/lebron/blob/main/docs/protoc.md)

```shell
├── etc
├── internal
│   ├── config
│   ├── logic # 业务目录，所有业务编码文件都存放在这个目录下面，logic 为固定后缀
│   ├── server
│   └── svc # 依赖注入目录，所有 logic 层需要用到的依赖都要在这里进行显式注入
├── order
└── orderclient
```