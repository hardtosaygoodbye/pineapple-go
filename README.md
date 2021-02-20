# gin-demo
用gin框架搭建的一个项目结构，方便快速开发项目。

### 特点

- 集成gorm，用户mysql存储层操作

- 集成go redis，用户操作缓存

- 集成uber/zap, zap是一个高效的日志组件

  项目中日志进行了分解:

  - request日志  

    记录http请求的request和response结果

  - init日志

    记录服务启动的日志

  - mysql日志 

    记录启动项目时，迁移数据库执行的sql情况日志

  - panic日志

    记录http请求产生的panic错误日志，方便快速定位错误问题

  - app日志

    记录开发过程中，我们记录的业务日志，这是我们最常用的日志，日志当中记录了requestId，方便快速根据请求查看当前请求的日志流,
    同时也有TraceId用于链路跟踪，分析用户行为，日志格式如下,包括了TraceId、requestId、file、keywords，以及我们记录的重要信息，体现在data里面:
    
 ```shell

    {"level":"error","time":"2020-09-25 13:47:53.578","keywords":"err","TraceId":"1acaba61-10e3-499f-9cb7-5cca47abfcc0","requestId":"c9f00ccb-f5f5-4c51-ad09-d3a29c46224a","file":"controller/home.go:19","data":"this is error"}
  
    {"level":"warn","time":"2020-09-25 13:47:54.246","keywords":"wa","TraceId":"1acaba61-10e3-499f-9cb7-5cca47abfcc0","requestId":"9d130d78-58a7-4271-914e-2c8de140d2be","file":"controller/home.go:17","data":"this is warn"}
  

 ```

- 集成jwt，一种流行的web身份认证方式，减轻服务端压力，将用户登录验证信息存储在可以端

- 集成gopkg.in/ini，用户解析我们的配置项

### 目录结构

- config -配置
- constant -常量
- controller -业务控制器
- logs -日志目录
- middleware -中间件目录
- model -表模型目录
- router -路由配置目录
- service -服务存放目录
- repository -数据库操作
- core -定义了log、db、redis等基础组件的封装、封装
- util -其它工具目录
- main.go -程序入库文件

### 其它

- 日志按天进行分割
- 根目录的.env.local进行本地配置,.env作为线上配置，避免误将本地配置提交到git仓库
- 业务日志可以根据requestId查看当次请求的所有日志，根据TraceId链路追踪用户行为，方便定位问题，分析逻辑行为
