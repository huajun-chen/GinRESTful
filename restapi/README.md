## 后端文档

### 项目简介

使用`Gin`框架搭建`RESTful`风格的基础脚手架，开箱即用，快速实现项目的业务功能开发

后端项目基本情况如下：

| 技术栈    | 概述                                                         |
| --------- | ------------------------------------------------------------ |
| Viper     | 配置管理                                                     |
| Zap       | 日志管理                                                     |
| Validator | 参数字段校验                                                 |
| GORM      | 数据库ORM                                                    |
| Redis     | 缓存数据库，Token黑名单，退出登录10秒后此Token不能再正常使用 |
| JWT       | 生成Token，登录验证Token                                     |
| Docker    | 服务基于Docker部署                                           |

### 实现功能

- 初始化`YAML`配置文件、初始化日志信息、初始化参数校验翻译、初始化`MySQL`、初始化`Redis`、初始化一个`admin`账户
- 实现用户注册，用户登录，修改用户信息、查看用户信息、登出、`admin`账户查看用户列表、`admin`用户删除用户信息
- 用户权限中间件、跨域中间件、`JWT`中间件、日志中间件
- 日志文件分割归档
- 用户注册、用户登录需要验证码
- 对用户/前端提交的参数进行校验
- `Token`黑名单，用户退出登录后10秒，将此`Token`加入`Redis`的黑名单中，此`Token`不能继续使用
- 后端响应返回的数据格式风格统一
- 服务优雅关机
- `i18n`国际化翻译
- 限制`IP`访问频率，每分钟只能访问100次
- 并发请求限流，通过开关控制是否开启此功能
- 获取系统健康信息，包括CPU、内存、磁盘等信息

### 项目目录

目录结构：

<div align="center"><img src="http://tva1.sinaimg.cn/large/0079DIvogy1h8s6vz185wj30ho0rs42c.jpg" alt="image.png" style="zoom:50%;" /></div>

目录说明：

**备注：GitHub仓库缺少个目录（logs），需要自己手动添加**

| 文件/目录        | 概述                                 | 文件概述                                                     |
| ---------------- | ------------------------------------ | ------------------------------------------------------------ |
| config           | 配置文件对应的结构体定义             | config.go：配置的对应的struct                                |
| controller       | 业务层                               | catpcha.go：生成图片验证码<br>health.go：获取系统CPU、内存、磁盘等信息<br>user.go：用户模块控制层面相关代码 |
| dao              | 操作数据库，给业务controller提供数据 | user.go：用户模块数据库操作                                  |
| forms            | 字段验证的struct                     | base.go：基础的参数对应的struct<br>user.go：用户模块参数，返回的数据结构对应的struct定义 |
| global           | 定义全局变量                         | globalvar.go：定义后端项目的全局变量                         |
| initialize       | 服务初始化                           | account.go：初始化一个admin账号<br>config.go：使用Viper初始化获取配置文件<br>logger.go：使用zap初始化项目日志<br>mysql.go：使用GORM初始化MYSQL数据库<br>redis.go：初始化Redis缓存数据库<br>router.go：初始化项目的路由<br>runserver.go：运行Gin服务，实现优雅关机<br>validator.go：使用Validator初始化参数校验，参数校验信息中英文翻译 |
| logs             | 日志存储                             | 存储每天的日志文件                                           |
| middlewares      | 中间件                               | admin.go：权限相关的中间件<br>cors.go：跨域中间件<br>i18n.go：i18n国际化中间件<br>ipfrequency.go：限制某个IP的访问频率<br>jwt.go：JWT验证中间件<br>logger.go：日志中间件<br>ratelimit.go：并发请求限流，通过开关控制 |
| models           | 数据库字段定义                       | user.go：用户模块的数据库字段                                |
| response         | 统一封装response                     | response.go：对后端返回的数据格式进行统一封装                |
| router           | 路由                                 | routerv1.go：V1版本路由                                      |
| static           | 静态资源文件夹                       | i18n：存放i18n国际化翻译json文件                             |
| utils            | 工具                                 | json.go：读取json文件，将其序列化为map<br>jwt.go：Token相关的函数/方法<br>md5.go：MD5计算<br>migration.go：执行main启动项目时对数据库表新建或迁移<br>page.go：与页数，每页的数量相关的代码封装<br>password.go：密码加密与密码校验<br>redis.go：与Redis操作相关的方法<br>systeminfo.go：获取系统CPU、内存、磁盘等信息<br>validator.go：参数校验出现错误时代码统一封装 |
| main.go          | 程序入口文件/主程序                  |                                                              |
| README.md        | 后端Readme文件                       |                                                              |
| setting-dev.yaml | 配置文件                             |                                                              |

### 参数

**主键使用路径变量，其他字段或其他查询参数用查询变量**

查询变量：

- `GET`方法通过`form`表单提交参数
- `POST`、`PUT`、`DELETE`方法通过`Body`的`json`格式提交参数

示例：

- 查看用户自己的信息

  ```sh
  GET /api/v1/user/666
  ```

- 查看用户名为xxx的信息

  ```sh
  GET /api/v1/user/info?name=张三
  ```

- 用户登录

  ```sh
  POST /api/v1/login
  ```

  参数如下：

  ```json
  {
      "user_name": "admin",
      "password": "admin12345"
  }
  ```

参数绑定：

参考文档：[点击跳转](https://cloud.tencent.com/developer/article/1689928)

`Gin`的`ShouldBind`可以绑定全部类型，在绑定前会对参数的类型做判断，此项目里默认不使用`ShouldBind`对参数进行绑定

- 使用`ShouldBindUri`绑定路径参数，`tag`标签：`uri:"id"`
- 使用`ShouldBindQuery`绑定`form`参数，`tag`标签：`form:"name"`
- 使用`ShouldBindJSON`绑定`json`参数，`tag`标签：`json:"name"`

### 响应

所有接口请求的响应系统状态码都返回`200`，在`code`字段自定义具体的状态码，自定义状态码为`200`表示请求成功，自定义状态码大于`200`表示请求失败

接口请求所返回的数据为`json`格式，如果字段没有值就不显示此字段，返回的字段如下：

- code：自定义响应状态码
- msg：返回的提示信息
- data：返回的数据

示例：

```json
{
    "code": 10000,
    "msg": "参数不正确",
    "data": {}
}
```

**备注：**<font color=Red>返回的`json`数据，只有`code`、`total`、`id`为数字，其他的全部字段都为字符串；项目中，返回字段的结构体与数据库字段的结构体不是同一个，返回的结构体单独定义，原因：返回的字段不一定全部都是数据库的字段，也有可能是数据库字段之间计算之后的值，所以返回的数据结构体单独定义</font>

### 日志记录

在`controller`控制层，生成验证码、获取系统资源信息记录`err`日志，其他模块里只记录与数据库操作有关的错误日志，不记录其他错误日志信息

### 编码规范

按照`Go编码规范`编写代码，风格统一

**函数或方法注释**

每个函数、方法都有注释说明，包括三个方面（顺序严格）：

- 函数或方法名：简要说明
- 参数列表：每行一个参数
- 返回值：每行一个返回值

示例：

```go
// DaoFindUserInfoToId 根据用户ID查询用户信息
// 参数：
//		userId：用户ID
// 返回值：
//		*models.User：用户信息的指针
//		bool：查询是否成功
func DaoFindUserInfoToId(userId uint) (*models.User, bool) {
  // 代码块
  ...
}
```

### API接口文档

详情见`docs`目录下`API接口文档`

### 运行项目

项目使用了`MySQL`数据库，`Redis`缓存数据库，`MySQL`和`Redis`在本地是基于`Docker`运行

需要保证本地已经有`Docker`环境

#### 1. Docker安装MySQL

`MySQL`的版本是5.7.30

拉取镜像：

```sh
docker pull mysql:5.7.30
```

运行容器：

```sh
docker run -itd --restart=always --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=1qaz@WSX -d mysql:5.7.30
```

参数说明：

- `-itd`：这个参数会启动一个交互式的、后台运行的`Docker`容器。它的作用是让`Docker`容器在后台运行，并且可以通过终端进行交互
- `--restart=always`：这个参数会设置Docker容器的重启策略为总是重启，如果`Docker`容器意外停止运行，`Docker`会自动重启它
- `--name mysql`：这个参数会指定`Docker`容器的名称。在这个例子中，`Docker`容器的名称为`mysql`
- `-p 3306:3306`：这个参数会指定端口映射。它的作用是将容器内的3306端口映射到宿主机的3306端口。这样，外部程序就可以通过宿主机的3306端口访问`Docker`容器内的`MySQL`服务
- `-e MYSQL_ROOT_PASSWORD=1qaz@WSX`：这个参数会指定`MySQL`的`root`用户的密码，此处密码为：`1qaz@WSX`。如果没有指定密码，外部程序就可以不用密码访问`MySQL`服务
- `-d mysql:5.7.30`：这个参数会指定要运行的`Docker`镜像。在这个例子中，我们指定要运行的是`MySQL 5.7.30`版本的镜像

上面的命令会运行一个名为`mysql`的`MySQL`容器，并将容器内的`MySQL`服务在宿主机的3306端口供外部程序访问。同时，该容器的重启策略设置为总是重启

`Redis`基础的命令：

```sh
Docker进入Redis：docker exec -it redis /bin/bash
启动客户端：redis-cli
输入密码：auth 1qaz@WSX
查看所有key：keys *
Set字符串：set testkey testvalues
Get字符串：get testkey
查看到期时间：ttl testkey
Del字符串：del testkey
```

#### 2. Docker安装Redis

`Redis`的版本是6.0

拉取镜像：

```sh
docker pull redis:6.0
```

运行容器：

```sh
docker run -itd --restart=always --name redis -p 6379:6379 redis:6.0 --requirepass "1qaz@WSX"
```

参数说明：

- `-itd`：这个参数会启动一个交互式的、后台运行的`Docker`容器。它的作用是让`Docker`容器在后台运行，并且可以通过终端进行交互
- `--restart=always`：这个参数会设置`Docker`容器的重启策略为总是重启，如果`Docker`容器意外停止运行，`Docker`会自动重启它
- `--name redis`：这个参数会指定`Docker`容器的名称。在这个例子中，`Docker`容器的名称为`redis`
- `-p 6379:6379`：这个参数会指定端口映射。它的作用是将容器内的6379端口映射到宿主机的6379端口。这样，外部程序就可以通过宿主机的6379端口访问`Docker`容器内的`Redis`服务

- `redis:6.0`：这个参数会指定要运行的`Docker`镜像。在这个例子中，我们指定要运行的是`Redis 6.0`版本的镜像
- `--requirepass "1qaz@WSX"`：这个参数会指定`Redis`的密码。如果没有指定密码，外部程序就可以不用密码访问`Redis`服务

上面的命令会运行一个名为`redis`的`Redis`容器，并将容器内的`Redis`服务在宿主机的6379端口供外部程序访问。同时，该容器的重启策略设置为总是重启

#### 3. 本地运行项目

如果是通过`git clone`将项目克隆到本地，需要进入到项目的`main.go`所在目录，在这个目录下，可以看到`main.go`，`go.mod`文件，在此目录下执行：`go mod download`安装依赖

依赖安装完成之后，在此目录下继续执行`go run main.go`来运行项目

**注意：** 运行项目之前需要保证`MySQL`和`Redis`都正常启动，配置文件`setting-dev.yaml`中`MySQL`和`Redis`的`host`都为：`127.0.0.1`

### 项目部署

#### 1. 本地Mac环境将Go Gin项目部署到Docker

**前提条件：确保项目能正常在本地运行，确保`Docker`的`MySQL`和`Redis`正常运行**

**尤为重要：需要将配置文件`setting-dev.yaml`中`Mysql`和`Redis`的`host`都改为：`docker.for.mac.localhost`**

```yaml
mysql:
  host: "docker.for.mac.localhost"

redis:
  host: "docker.for.mac.localhost"
```

在`restapi`目录里，此目录可以看到`main.go`、`go.mod`、`Dockerfile`文件，在此目录下执行命令构建镜像，`goginimage`是构建的镜像名

```sh
docker build . -t goginimage
```

运行容器：

```sh
docker run -d -it --link mysql:mysql --link redis:redis -p 8080:8080 -v /Users/chenhuajun/golangprojects/GinRESTful/restapi:/Users/chenhuajun/golangprojects/GinRESTful/restapi/app/ --add-host=ginrestfulapp-host:127.0.0.1 --name ginrestfulapp goginimage
```

参数说明：

- `-d`：该参数表示在容器运行时以后台模式运行。
- `-it`：该参数表示容器的标准输入、输出和错误将被绑定到宿主机的终端，并且容器运行时会交互式的运行。
- `--link`：该参数允许将此容器连接到另一个容器，并且在连接的容器中可以通过主机名访问另一个容器。在这里，可以通过名为`ginrestfulapp`的容器访问 `mysql` 和 `redis` 容器
- `-p`：该参数允许将容器中的端口映射到宿主机的端口。在这里，将容器中的 8080 端口映射到宿主机的 8080 端口
- `-v`：该参数允许挂载宿主机的目录或者文件到容器中。在这里，将宿主机的目录 `/Users/chenhuajun/golangprojects/GinRESTful/restapi` 挂载到容器的 `/Users/chenhuajun/golangprojects/GinRESTful/restapi/app/` 目录中
- `--add-host`：该参数允许在容器中添加一个自定义的主机名映射。在这里，将主机名 `ginrestfulapp-host` 映射到 IP 地址 `127.0.0.1`
- `--name`：该参数用于指定容器的名称，容器名称为：`ginrestfulapp,`，镜像名称为：`goginimage`

该容器会绑定宿主机的终端，并且容器的标准输入、输出和错误都将被绑定到宿主机的终端。容器运行时会交互式的运行。该容器连接到名为 `mysql` 和 `redis` 的容器，可以通过主机名 `mysql` 和 `redis` 访问连接的容器。将容器中的 8080 端口映射到宿主机的 8080 端口，这样宿主机就可以通过该端口访问该容器。将宿主机的目录 `/Users/chenhuajun/golangprojects/GinRESTful/restapi` 挂载到容器的 `/Users/chenhuajun/golangprojects/GinRESTful/restapi/app/` 目录中，这样宿主机中的文件就可以在容器中使用。在容器中添加一个自定义的主机名映射，将主机名 `ginrestfulapp-host` 映射到 IP 地址 `127.0.0.1`。该容器的名称为 `ginrestfulapp`

#### 2. 远程Linux服务器部署

TODO



---

