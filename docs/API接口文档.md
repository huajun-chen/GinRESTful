## API接口文档

接口文档示例：

- 描述

  这是一个`API`接口文档的描述

- 请求

  ```json
  GET /api/v1/example
  ```

- 参数列表

  | 参数名称                         | 类型                | 取值范围       | 是否必须 | 默认值           | 备注     |
  | -------------------------------- | ------------------- | -------------- | -------- | ---------------- | -------- |
  | 参数名称，英文单词或英文单词组合 | string/int/bool/... | 参数的取值范围 | yes/no   | 参数是否有默认值 | 备注信息 |

- 返回示例

  ```json
  {
      "code": 200,
      "msg": "",
      "data": {}
  }
  ```

  - code：返回自定义的状态码，200代表请求成功，超过200代表请求失败
  - msg：返回给前端的提示信息，正确的提示信息或错误的提示信息，可选项
  - data：返回的数据对象，可选项

---

### 1. 基础接口

#### 1.1 验证码

- 描述

  用户注册和登录时需要进行验证码验证，防止恶意注册和登录，返回验证码图片的`base64`码

- 请求

  ```json
  GET /api/v1/captcha
  ```

- 参数列表

  无参数

- 返回示例

  ```json
  {
      "code": 200,
      "data": {
          "captcha_id": "SEAQfQ3nb2ysBNy2QkuJ",
          "captcha_path": "data:image/png;base64,iVB..."
      }
  }
  ```

### 2. 用户模块

#### 2.1 用户注册

- 描述

  用户注册账户

- 请求

  ```json
  POST /api/v1/register
  ```

- 参数列表

  | 参数名称   | 类型   | 取值范围 | 是否必须 | 默认值 | 备注       |
  | ---------- | ------ | -------- | -------- | ------ | ---------- |
  | user_name  | string |          | yes      |        | 用户名     |
  | password   | string |          | yes      |        | 密码       |
  | password2  | string |          | yes      |        | 重复密码   |
  | captcha    | string |          | yes      |        | 验证码结果 |
  | captcha_id | string |          | yes      |        | 验证码ID   |

- 返回示例

  ```json
  {
      "code": 200,
      "msg": "注册成功",
      "data": {
          "id": 1,
          "name": "zhangsan",
          "token": "token信息"
      }
  }
  ```
  

#### 2.2 用户登录

- 描述

  用户登录

- 请求

  ```json
  POST /api/v1/user/login
  ```

- 参数列表

  | 参数名称 | 类型 | 取值范围 | 是否必须 | 默认值 | 备注 |
  | -------- | ---- | -------- | -------- | ------ | ---- |
  | user_name | string |          | yes      |        | 用户名 |
  | password2  | string |          | yes      |        | 重复密码   |
  | captcha    | string |          | yes      |        | 验证码结果 |
  | captcha_id | string |          | yes      |        | 验证码ID   |

- 返回示例

  ```json
  {
      "code": 200,
      "msg": "登录成功",
      "data": {
          "id": 1,
          "name": "zhangsan",
          "token": "token信息"
      }
  }
  ```

#### 2.3 查看用户信息

- 描述

  用户查看自己的信息

- 请求

  ```json
  GET /api/v1/user/info/1
  ```

- 参数列表

  | 参数名称 | 类型 | 取值范围  | 是否必须 | 默认值 | 备注   |
  | -------- | ---- | --------- | -------- | ------ | ------ |
  | id       | int  | 大于等于1 | yes      |        | 用户ID |

- 返回示例

  ```json
  {
      "code": 200,
      "data": {
          "id": 1,
          "created_at": "2022-12-06",
          "user_name": "zhangsan",
          "gender": "3",
          "desc": "这个人很懒，什么都没留下...",
          "role": "2",
          "mobile": "",
          "email": ""
      }
  }
  ```

#### 2.4 修改用户信息

- 描述

  用户修改自己的信息

- 请求

  ```json
  PUT /api/v1/user/info/1
  ```

- 参数列表

  | 参数名称     | 类型   | 取值范围       | 是否必须 | 默认值 | 备注       |
  | ------------ | ------ | -------------- | -------- | ------ | ---------- |
  | id           | int    | 大于等于1      | yes      |        | 用户ID     |
  | password_old | string |                | no       |        | 旧密码     |
  | password     | string |                | no       |        | 新密码     |
  | password2    | string |                | no       |        | 重复新密码 |
  | gender       | int    | [1,2,3]        | no       |        | 性别       |
  | desc         | string |                | no       |        | 描述       |
  | mobile       | string | 合法的电话号码 | no       |        | 电话       |
  | email        | string | 合法的邮箱     | no       |        | 邮箱       |

- 返回示例

  ```json
  {
      "code": 200,
      "msg": "数据修改成功"
  }
  ```
  

#### 2.5 删除用户信息

- 描述

  管理员删除某个用户的信息，假删除，数据实际依然存在，默认对外不可见

- 请求

  ```json
  DELETE /api/v1/user/info/1
  ```

- 参数列表

  | 参数名称 | 类型 | 取值范围  | 是否必须 | 默认值 | 备注   |
  | -------- | ---- | --------- | -------- | ------ | ------ |
  | id       | int  | 大于等于1 | yes      |        | 用户ID |

- 返回示例

  ```json
  {
      "code": 200,
      "msg": "数据删除成功"
  }
  ```

#### 2.6 查看用户列表

- 描述

  管理员查看全部的用户列表信息

- 请求

  ```json
  GET /api/v1/user/list
  ```

- 参数示例

  | 参数名称  | 类型 | 取值范围  | 是否必须 | 默认值 | 备注        |
  | --------- | ---- | --------- | -------- | ------ | ----------- |
  | page      | int  | [1~10000] | no       |        | 页数/第几页 |
  | page_size | int  | [1~10000] | no       |        | 每页的数量  |

- 返回示例

  ```json
  {
      "code": 200,
      "data": {
          "total": 6,
          "values": [
              {
                  "id": 2,
                  "created_at": "2022-12-06",
                  "user_name": "user1",
                  "gender": "3",
                  "desc": "这个人很懒，什么都没留下...",
                  "role": "2",
                  "mobile": "",
                  "email": ""
              },
              {
                  "id": 1,
                  "created_at": "2022-12-05",
                  "user_name": "admin",
                  "gender": "1",
                  "desc": "这个人很懒，什么都没留下...",
                  "role": "1",
                  "mobile": "",
                  "email": ""
              }
          ]
      }
  }
  ```
  



---















