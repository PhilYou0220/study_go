`conf`:开发环境配置

`conf_online` 生成环境配置

`conf_qa` 测试环境配置

`controller` 1 参数校验

`dao`数据库查数据相关

`inner`内部调用

`log`日志文件存放处

`logger`zap双核心日志 一个是info级别日志一个是error级别日志

`middleware`中间件重写了 gin.default()使的zap作用于全局

`pkg`可以被外部访问的库

`router`路由

`service`业务层 2 业务逻辑 3返回数据

`setting` viper读取conf

`main.go` 入口函数，加载配置，初始化，启动程序

启动时 输入 module名加配置文件如
    
    gin_zap conf/conf.yaml
