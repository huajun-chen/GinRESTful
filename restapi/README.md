## 后端文档

### 项目简介

使用`Gin`框架搭建`RESTful`风格的基础脚手架，开箱即用，快速实现项目的业务功能开发

后端项目基本情况如下：

| 技术栈    | 概述                     |
| --------- | ------------------------ |
| Viper     | 配置管理                 |
| Zap       | 日志管理                 |
| Validator | 参数字段校验             |
| GORM      | 数据库ORM                |
| Redis     | 缓存数据库               |
| JWT       | 生成Token，登录验证Token |
| Docker    | 服务基于Docker部署       |


### 项目目录

目录结构：

<div align="center"><img src="http://tva1.sinaimg.cn/large/0079DIvogy1h8s6vz185wj30ho0rs42c.jpg" alt="image.png" style="zoom:50%;" /></div>

目录说明：

**备注**：<font color=Red>GitHub仓库缺少2个目录（logs、static），需要自己手动添加</font>

| 文件/目录        | 概述                                 | 文件概述                                                     |
| ---------------- | ------------------------------------ | ------------------------------------------------------------ |
| config           | 配置文件对应的结构体定义             | config.go：配置的对应的struct                                |
| controller       | 业务层                               | catpcha.go：生成图片验证码<br>user.go：用户模块控制层面相关代码 |
| dao              | 操作数据库，给业务controller提供数据 | user.go：用户模块数据库操作                                  |
| forms            | 字段验证的struct                     | user.go：用户模块参数对应的struct定义                        |
| global           | 定义全局变量                         | globalvar.go：定义后端项目的全局变量                         |
| initialize       | 服务初始化                           | config.go：使用Viper初始化获取配置文件<br>logger.go：使用zap初始化项目日志<br>mysql.go：使用GORM初始化MYSQL数据库<br>redis.go：初始化Redis缓存数据库<br>router.go：初始化项目的路由<br>validator.go：使用Validator初始化参数校验，参数校验信息中英文翻译 |
| logs             | 日志存储                             | 存储每天的日志文件                                           |
| middlewares      | 中间件                               | admin.go：权限相关的中间件<br>cors.go：跨域中间件<br>jwt.go：JWT验证中间件<br>logger.go：日志中间件 |
| models           | 数据库字段定义                       | user.go：用户模块的数据库字段                                |
| response         | 统一封装response                     | response.go：对后端返回的数据格式进行统一封装                |
| router           | 路由                                 | base.go：项目的基础路由<br>user.go：用户模块路由             |
| static           | 资源文件夹                           | 存放静态资源的目录                                           |
| utils            | 工具                                 | createtoken.go：生成Token<br>migration.go：执行main启动项目时对数据库表新建或迁移<br>page.go：与页数，每页的数量相关的代码封装<br>validator.go：参数校验出现错误时代码统一封装 |
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







---

