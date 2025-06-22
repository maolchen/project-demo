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
集群的增删改查  
K8S ns pod deployment sts cronjob 的增删改查  
...

### 目录结构 
```bash
├── api
│   ├── cluster.go  #cluster接口
│   ├── create.go   #面向对象方式的namespace create接口，练习使用，实际未使用
│   ├── delete.go   #面向对象方式的namespace delete接口，练习使用，实际未使用
│   ├── get.go      #面向对象方式的namespace get接口，练习使用，实际未使用
│   ├── list.go     #面向对象方式的namespace list接口，练习使用，实际未使用
│   ├── resource    #接口方式的K8S resource增删改查api 接口
│   │   ├── Create.go
│   │   ├── Delete.go
│   │   ├── Get.go
│   │   ├── List.go
│   │   └── Update.go
│   ├── update.go   #面向对象方式的namespace update接口，练习使用，实际未使用
│   └── user.go
├── conf   #测试config.yaml配置文件，启动服务--config指定
│   └── config1.yaml  
├── config #config配置相关
│   └── config.go
├── constants #常量
│   ├── code.go
│   ├── common.go
│   └── message.go
├── database  #数据库
│   └── sqlite.go
├── go.mod
├── go.sum
├── initializa  #初始化
│   ├── initCluster.go
│   ├── initConfig.go
│   ├── initDB.go
│   ├── initPrint.go
│   ├── initRouter.go
│   └── initZap.go
├── main.go
├── middlewares  #中间件，jwt和trace_id日志封装
│   ├── jwt.go
│   └── logger.go
├── models      #数据模型
│   ├── basicInfo.go
│   ├── cluster.go
│   ├── db.go
│   └── user.go
├── README.md
├── routers   #路由
│   ├── cluster.go
│   ├── login.go
│   ├── namespace.go  #面向对象方式的练习路由，实际未使用
│   ├── resource.go 
│   └── user.go
├── service   #服务层
│   ├── base.go   #k8s基础通用方法
│   ├── clusters  #K8S多集群相关逻辑
│   │   ├── cluster.go
│   │   ├── clusterInfo.go
│   │   ├── utils.go
│   │   └── wrapper.go
│   ├── common   #通用函数
│   │   └── newClientSet.go
│   ├── factory  
│   │   └── factory.go
│   ├── interfaces.go  #定义接口
│   ├── login.go    #登录登出逻辑
│   ├── namespace   #练习的namespace的curd逻辑
│   │   ├── create.go
│   │   ├── delete.go
│   │   ├── get.go
│   │   ├── list.go
│   │   └── update.go
│   ├── resource  #K8S资源curd
│   │   ├── cronjob.go
│   │   ├── daemonset.go
│   │   ├── deployment.go
│   │   ├── namespace.go
│   │   ├── pod.go
│   │   └── statefulset.go
│   └── user.go   #用户相关逻辑
├── utils      #通用工具包
│   ├── dirExists.go
│   ├── jwtutils.go
│   ├── kubeconfig_validator.go
│   ├── logutils.go
│   ├── request.go
│   ├── response.go
│   ├── strings.go
│   └── struct2Map.go
└── validator  #用户校验
    └── user_validator.go
```



### 环境变量 
除了使用配置文件，还可以使用环境变量配置相关启动参数
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
	"JWT_EXPIRES", 30  #token过期时间，分钟
```
	
### TODO
练习项目，多次修改结构，结构可能也还有不合理的地方，早期写的用户api和登录登出api的返回数据需要优化，status返回的是string类型，和cluster和resource的返回不一致。
