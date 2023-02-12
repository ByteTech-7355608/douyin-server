# douyin-server
## 项目部署
### 1.基础依赖
```shell
go mod tidy
docker-compose up
```

### 2.启动base服务
```shell
cd cmd/main
export ServiceName=douyin.base
go run .
```

### 3.启动interaction服务
```shell
cd cmd/main
export ServiceName=douyin.interaction
go run .
```

### 4.启动api服务
```shell
cd cmd/main
export ServiceName=douyin.api
go run .
```
