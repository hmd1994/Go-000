学习笔记
├── README.md
├── api             # rpc文件
│   └── hello
│          ├── hello.pb.go
│          └── hello.proto
├── cmd
│   └── hello      #项目的主干，文件夹名和可执行文件匹配
|				├── hello
│       ├── main.go
|
├── go.mod
├── go.sum
├── configs
│   ├── config.go       # 初始化配置
│    └── config.yaml     # 配置 yaml 文件
├── internal        #私有应用程序和库代码
│   ├── biz         # 业务组装层
│   ├── data        # 数据层
│   ├── pkg         # 公共库
│   │   └── errcode    # 自定义错误
│   │       ├── common.go  # 通用错误      
│   │       └── errors.go  # 错误定义
│   ├── dao         # dao层
│   └── service     # API实现层
│       └── hello.go
└── pkg            # 测试服务API
    └── model # 数据库
          └── init    # 初始化