# gin-rest
go语言，基于gin框架的rest简单模板

## 启动
```bash
2018/12/07 17:11:21 ====== 配置文件加载成功 ======
2018/12/07 17:11:21 ====== 数据库初始化成功 ======
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /v1/user/                 --> controllers.(*UserController).Index-fm (3 handlers)
[GIN-debug] GET    /v1/user/create           --> controllers.(*UserController).Create-fm (3 handlers)
[GIN-debug] GET    /v1/about/                --> controllers.(*AboutController).Index-fm (3 handlers)
[GIN-debug] Listening and serving HTTP on 0.0.0.0:8081
```
