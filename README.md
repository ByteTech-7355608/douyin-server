# douyin-server
## 项目部署
### 1. 基础依赖
```shell
go mod tidy
docker-compose up
```

### 2. 启动base服务
```shell
cd script
sh run_base.sh
```

### 3. 启动interaction服务
```shell
cd script
sh run_interaction.sh
```

### 4. 启动social服务
```shell
cd script
sh run_social.sh
```

### 5. 启动api服务
```shell
cd script
sh run_api.sh
```

### 6.Jaeger
打开浏览器进入 `http://127.0.0.1:16686/`