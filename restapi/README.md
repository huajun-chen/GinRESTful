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


### 后端项目目录

目录结构：

<div align="center"><img src="http://tva1.sinaimg.cn/large/0079DIvogy1h8qp4wqunvj30hm0t2gpo.jpg" alt="image.png" style="zoom:45%;" /></div>

目录说明：

| 文件/目录        | 概述                                 |
| ---------------- | ------------------------------------ |
| config           | 配置文件对应的结构体定义             |
| controller       | 业务层                               |
| dao              | 操作数据库，给业务controller提供数据 |
| forms            | 字段验证的struct                     |
| global           | 定义全局变量                         |
| initialize       | 服务初始化                           |
| logs             | 日志存储                             |
| middlewares      | 中间件                               |
| models           | 数据库字段定义                       |
| response         | 统一封装response                     |
| static           | 资源文件夹                           |
| router           | 路由                                 |
| setting-dev.yaml | 配置文件                             |
| main.go          | 程序入口文件/主程序                  |

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

  ```json
  POST /api/v1/login
  参数如下：
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









---

