# gin-make
## a tool for making gin struct
### How to use?
`gm -g Your-Project-Name`

# 项目结构

`my-gin-project/
├── cmd/                    # 应用程序入口
│   └── main.go             # 主程序入口，初始化和启动服务
├── config/                 # 配置文件
│   ├── config.go           # 配置加载逻辑
│   └── config.yaml         # 配置文件（YAML、JSON 等格式）
├── internal/               # 内部模块（不对外暴露）
│   ├── app/                # 应用核心逻辑
│   │   ├── app.go          # Gin 应用初始化
│   │   └── routes.go       # 路由定义
│   ├── database/           # 数据库相关
│   │   ├── db.go           # 数据库连接初始化
│   │   └── migrations/     # 数据库迁移脚本
│   ├── models/             # 数据模型
│   │   └── user.go         # 用户模型定义
│   ├── handlers/           # HTTP 处理函数
│   │   ├── user.go         # 用户相关的处理逻辑
│   │   └── auth.go         # 认证相关的处理逻辑
│   ├── middleware/         # 中间件
│   │   ├── auth.go         # 认证中间件
│   │   └── logger.go       # 日志中间件
│   ├── repository/         # 数据访问层
│   │   └── user_repo.go    # 用户数据操作
│   └── services/           # 业务逻辑层
│       └── user_service.go # 用户业务逻辑
├── pkg/                    # 可复用的外部包（可选）
│   └── utils/              # 工具函数
│       └── helpers.go      # 通用辅助函数
├── api/                    # API 定义（可选）
│   └── swagger/            # Swagger 文档
├── tests/                  # 测试文件
│   ├── user_test.go        # 用户相关测试
│   └── auth_test.go        # 认证相关测试
├── go.mod                  # Go 模块定义
├── go.sum                  # 依赖校验文件
└── README.md               # 项目说明
`
