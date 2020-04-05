# Mogo

Mogo: Simple Single Golang Web Service

Mogo是基于singo框架进行优化和新增功能的web服务开发框架

https://github.com/bydmm/singo

## 优化

1. 修复了由日志对象初始化失败引起的程序错误，并支持从环境变量设置日志级别
2. 完善了跨域设置，修复了原先因gin-modle设置失败导致上线后跨域失败的问题
3. 修复了一键部署时，GORM创建数据库表时编码错误导致乱码的问题
4. mogo-v1保留了原先以cookie实现的session来保存登录状态的代码，如果需要可以自行选用mogo版本
5. 对 Response 基础序列化器返回的不必要的json数据进行忽略。

## 新增

1. 实现了基于jwt-go的token用户登录状态验证，并封装了jwt的基本属性，可通过环境变量设置

## [mogo](https://github.com/Airmomo/mogo/tree/master) 对比  [mogo-v1](https://github.com/Airmomo/mogo/tree/v1)

<table>
        <tr>
            <th></th>
            <th>master</th>
            <th>v1</th>      
        </tr>
        <tr>
            <th>更新速度</th>
            <th>快</th>
            <th>较慢</th>
        </tr>
        <tr>
            <th>用户登录状态</th>
            <th>Json Web Token</th>
            <th>Cookies-Session</th>
        </tr>
        <tr>
            <th>Gin</th>
            <th>√</th>
            <th>√</th>
        </tr>
        <tr>
             <th>Gin-Session</th>
             <th>√</th>
             <th>√</th>
        </tr>
        <tr>
             <th>Gin-Cors</th>
             <th>√</th>
             <th>√</th>
        </tr>
        <tr>
             <th>GORM</th>
             <th>√</th>
             <th>√</th>
        </tr>
        <tr>
             <th>Go-Redis</th>
             <th>√</th>
             <th>√</th>
        </tr>
        <tr>
            <th>JWT-Go</th>
            <th>√</th>
            <th>×</th>
        </tr>
        <tr>
             <th>Go dot env</th>
             <th>√</th>
             <th>√</th>
        </tr>
        <tr>
            <th>Logger</th>
            <th>√</th>
            <th>√</th>
        </tr>
</table>

## 使用mogo开发的项目实例

https://github.com/Airmomo/jilijili

## 目的

本项目采用了一系列Golang中比较流行的组件，可以以本项目为基础快速搭建Restful Web API

## 特色

本项目已经整合了许多开发API所必要的组件：

1. [Gin](https://github.com/gin-gonic/gin): 轻量级Web框架，自称路由速度是golang最快的 
2. [GORM](http://gorm.io/docs/index.html): ORM工具。本项目需要配合Mysql使用 
3. [Gin-Session](https://github.com/gin-contrib/sessions): Gin框架提供的Session操作工具
4. [Go-Redis](https://github.com/go-redis/redis): Golang Redis客户端
5. [godotenv](https://github.com/joho/godotenv): 开发环境下的环境变量工具，方便使用环境变量
6. [Gin-Cors](https://github.com/gin-contrib/cors): Gin框架提供的跨域中间件
7. [jwt-go](https://github.com/dgrijalva/jwt-go): JSON WEB TOKEN 验证 (JWT)
8. 自行实现了国际化i18n的一些基本功能
9. 已实现的可实例化的Log日志对象,可用于打印日志,通过env全局变量配置日志级别
10. Branch master 使用基于jwt-go发放的jwt来保存用户的登录状态.

本项目已经预先实现了一些常用的代码方便参考和复用:

1. 创建了用户模型
2. 实现了```/api/v1/user/register```用户注册接口
3. 实现了```/api/v1/user/login```用户登录接口
4. 实现了```/api/v1/user/me```用户资料接口(需要登录后获取session)
5. 实现了```/api/v1/user/logout```用户登出接口(需要登录后获取session)

本项目已经预先创建了一系列文件夹划分出下列模块:

1. api文件夹就是MVC框架的controller，负责协调各部件完成任务
2. model文件夹负责存储数据库模型和数据库操作相关的代码
3. service负责处理比较复杂的业务，把业务代码模型化可以有效提高业务代码的质量（比如用户注册，充值，下单等）
4. serializer储存通用的json模型，把model得到的数据库模型转换成api需要的json对象
5. cache负责redis缓存相关的代码
6. auth权限控制文件夹
7. util一些通用的小工具
8. conf放一些静态存放的配置文件，其中locales内放置翻译相关的配置文件

## Godotenv

项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建.env文件设置环境变量便于使用(建议开发环境使用)

```shell
MYSQL_DSN="db_user:db_password@(localhost:3306)/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接地址

REDIS_ADDR="localhost:6379" # Redis端口和地址
REDIS_PW="" # Redis连接密码
REDIS_DB="0" # Redis库从0到10

#JWT属性设置
JWT_HEAD="Bearer" # 请求头中的标注
JWT_SECRET="youneedtoset" # JWT密钥，必须设置而且不要泄露
JWT_ISSUER="myProject" # JWT的签发人
JWT_SigningMethod="HS256" # JWT的声明加密算法，默认设置为HS256，

SESSION_SECRET="gin-session" # Seesion名称，必须设置
SESSION_SECRET="setOnProducation" # Seesion密钥，必须设置而且不要泄露

GIN_MODE="debug" # gin框架运行环境

LOG_LEVEL="debug" # 日志输出的级别

# OSS对象存储设置
OSS_END_POINT="oss-cn-hongkong.aliyuncs.com" 
OSS_ACCESS_KEY_ID="xxx"
OSS_ACCESS_KEY_SECRET="qqqq"
OSS_BUCKET="lalalal"
```

## Go Mod

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

```shell
go mod init go-crud
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go run main.go // 自动安装
```

## 运行

```shell
go run main.go
```

项目运行后启动在3000端口（可以修改，参考gin文档)