# gin-3at

3at golang gin 后端

## Quick Start 

### 环境变量设置

```bash
cp .env.example .env
```

然后修改其中的配置项。

- 数据库连接配置

### 文档生成

```bash
swag init -g cmd/server/main.go
```

### 启动热重载服务

```bash
air -c .air.toml
```

### 执行测试脚本

```bash
godotenv -f .env go test -count=1 ./...
```

> `-count=1` 参数用于确保每次测试都从头开始执行，而不是使用缓存的结果。在执行 dao 测试需要。

## 依赖的组件

### 库

- gin
- db
    - gorm.io/gorm
    - driver/mysql
- 文档工具：swaggo
- 配置管理：caarlos0/env/v11

### 命令行

- 环境变量管理：joho/godotenv
- 文档生成工具：swaggo
- 热重载：cosmtrek/air 
