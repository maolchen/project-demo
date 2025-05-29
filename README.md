# go语言练习后端脚手架
### 基于gin + gorm + sqlite的一个后端脚手架练习
### 功能
具有通过环境变量或者yaml配置文件两种配置方式  
基于用户结构体自动创建sqlite表  
创建用户   
用户登录接口  
打印接口请求耗时日志中间件  
jwt认证   
zap全局日志  

### 目录结构 
```bash
D:.
│  .gitignore
│  app.db
│  app.log
│  go.mod
│  go.sum
│  main.go
│  README.md
├─api    #api接口
│      user.go
├─conf  #配置文件示例
│      config1.yaml
├─config  #全局config
│      config.go
├─constants    #常量
│      code.go
│      common.go
│      message.go
├─database  #数据库
│      sqlite.go
├─initializa   #初始化
│      initConfig.go  #初始化配置
│      initDB.go      #初始化数据库
│      initPrint.go   #初始打印的一些东西
│      initRouter.go  #初始化路由
│      initZap.go     #初始化日志
├─middlewares      #中间件
│      jwt.go
│      requestTime.go
├─models    #数据模型
│      user.go
├─routers   #路由
│      login.go
│      user.go
├─service   #service  路由->api->service
│      login.go
│      user.go
├─utils   #工具类
│      dirExists.go
│      jwtutils.go
│      logutils.go
│      strings.go
└─validator  #用户信息校验
        user_validator.go
```



### 环境变量 
```shell
    "ADDRESS", ":8000"
    "DB_PATH", ".\\data\\app.db"
	"SECRET", "default_secret_key"
	"LOG_CONF_LOG_FILE", ""
	"LOG_CONF_MAX_AGE", 7
	"LOG_CONF_MAX_BACKUPS", 5
	"LOG_CONF_MAX_SIZE", 10
	"LOG_CONF_COMPRESS", true
	"LOG_CONF_LOG_LEVEL", "debug"
	"LOG_CONF_LOG_TYPE", "text"
	"JWT_EXPIRES", 30
```
	

