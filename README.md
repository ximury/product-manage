# 微服务 product-manage

![golang](https://img.shields.io/badge/Golang-v1.19-green.svg)

---

<font size=6 color=red face="华文新魏">【目录】</font>

[服务拆分](#div1)

&nbsp;&nbsp;[按业务服务拆分](#div1.1)

&nbsp;&nbsp;[按调用方式拆分](#div1.2)

[项目框架](#div2)

[搭建步骤](#div3)

&nbsp;&nbsp;[user 服务](#div3.1)

&nbsp;&nbsp;[product 服务](#div3.2)


## 服务拆分
<a name="div1" style=" position: relative;top: -180px;display: block;height: 0;overflow: hidden;"></a>

### 按业务服务拆分
<a name="div1.1" style=" position: relative;top: -180px;display: block;height: 0;overflow: hidden;"></a>

  - 用户服务（user）
  - 订单服务（order）
  - 产品服务（product）
  - 支付服务（pay）

### 按调用方式拆分
<a name="div1.2" style=" position: relative;top: -180px;display: block;height: 0;overflow: hidden;"></a>

| 区别                             | API 服务                                                                                    | RPC 服务                                                      |
|--------------------------------|-------------------------------------------------------------------------------------------|-------------------------------------------------------------|
| 传输协议                           | 基于 HTTP 协议                                                                                | 可以基于 HTTP 协议，也可以基于 TCP 协议                                   |
| 传输效率                           | 如果是基于 http1.1 的协议，请求中会包含很多无用的内容，如果是基于 HTTP2.0，那么简单的封装下可以作为一个 RPC 来使用，这时标准的 RPC 框架更多的是服务治理 | 使用自定义的 TCP 协议，可以让请求报文体积更小，或者使用 HTTP2 协议，也可以很好的减小报文体积，提高传输效率 |
| 性能消耗                           | 大部分是基于 json 实现的，字节大小和序列化耗时都比 thrift 要更消耗性能                                                | 可以基于 thrift 实现高效的二进制传输                                      |
| 负载均衡                           | 需要配置 Nginx、HAProxy 配置                                                                     | 基本自带了负载均衡策略                                                 |
| 服务治理：（下游服务新增，重启，下线时如何不影响上游调用者） | 需要事先通知，如修改 NGINX 配置                                                                       | 能做到自动通知，不影响上游                                               |

## 项目框架
<a name="div2" style=" position: relative;top: -180px;display: block;height: 0;overflow: hidden;"></a>

```text
├── common           # 通用库
├── service          # 服务
│   ├── order
│   │   ├── api      # order api 服务
│   │   ├── model    # order 数据模型
│   │   └── rpc      # order rpc 服务
│   ├── pay
│   │   ├── api      # pay api 服务
│   │   ├── model    # pay 数据模型
│   │   └── rpc      # pay rpc 服务
│   ├── product
│   │   ├── api      # product api 服务
│   │   ├── model    # product 数据模型
│   │   └── rpc      # product rpc 服务
│   └── user
│       ├── api      # user api 服务
│       ├── model    # user 数据模型
│       └── rpc      # user rpc 服务
└── go.mod
```

## 搭建步骤
<a name="div3" style=" position: relative;top: -180px;display: block;height: 0;overflow: hidden;"></a>

### user 服务
<a name="div3.1" style=" position: relative;top: -180px;display: block;height: 0;overflow: hidden;"></a>

1. 生成 user model 模型
   ```shell
   cd user
   # 添加sql，创建表信息
   vim model/user.sql
   goctl model mysql ddl -src ./model/user.sql -dir ./model -c
   ```
2. 创建 api 文件
   ```shell
   vim api/user.api
    ```
3. 根据api文件生成服务
   ```shell
   goctl api go -api ./api/user.api -dir ./api/
    ```
4. 创建 proto 文件
   ```shell
   vim rpc/product.proto
    ```
5. 根据 proto 文件生成代码
   ```shell
   goctl rpc protoc product.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.
    ```
6. 编写 user rpc 服务 - 修改 user.yaml 配置文件
   ```shell
   vim rpc/etc/user.yaml
    ```
7. 添加 user model 依赖
   ```shell
   vim rpc/internal/config/config.go
    ```
8. 注册服务上下文 user model 的依赖
   ```shell
   vim rpc/internal/svc/servicecontext.go
    ```
9. 添加用户注册逻辑配置

   i. 添加密码加密工具
   ```shell
   cd common
   mkdir cryptx && cd cryptx
   vim crypt.go
   ```

   ii. 添加密码加密 Salt 配置
   ```shell
   vim rpc/etc/user.yaml
   vim rpc/internal/config/config.go
   ```
10. 添加用户注册逻辑
   ```shell
   vim rpc/internal/logic/registerlogic.go
   ```
   > 注意：此文件会有 int64 与 uint64 类型的转换问题
11. 添加用户登录逻辑
   ```shell
   vim rpc/internal/logic/loginlogic.go
   ```
12. 添加用户信息逻辑
   ```shell
   vim rpc/internal/logic/userinfologic.go
   ```
13. 编写 user api 服务 - 修改 user.yaml 配置文件
   ```shell
   vim api/etc/user.yaml
   ```
14. 添加 user rpc 依赖
   ```shell
   # 添加 user rpc 服务配置
   vim api/etc/user.yaml
   # 添加 user rpc 服务配置的实例化
   vim api/internal/config/config.go
   # 注册服务上下文 user rpc 的依赖
   vim api/internal/svc/servicecontext.go
   ```
15. 添加用户注册逻辑
   ```shell
   vim api/internal/logic/registerlogic.go
   ```
16. 添加用户登录逻辑
   ```shell
   # 添加 JWT 工具
   vim common/jwtx/jwt.go
   # 添加用户登录逻辑
   vim api/internal/logic/loginlogic.go
   ```
17. 添加用户信息逻辑
   ```shell
   vim api/internal/logic/userinfologic.go
   ```
18. 启动 user rpc 服务
   ```shell
   $ cd service/user/rpc
   $ go run user.go -f etc/user.yaml
   ```
   > Starting rpc server at 127.0.0.1:9000...
19. 启动 user api 服务
   ```shell
   $ cd service/user/api
   $ go run user.go -f etc/user.yaml
   ```
   > Starting server at 0.0.0.0:8000...

### product 服务
<a name="div3.2" style=" position: relative;top: -180px;display: block;height: 0;overflow: hidden;"></a>

1. 生成 product model 模型
   ```shell
   cd product
   # 添加sql，创建表信息
   vim model/product.sql
   goctl model mysql ddl -src ./model/product.sql -dir ./model -c
   ```
2. 创建 api 文件并生成 api 服务
   ```shell
   vim api/product.api
   goctl api go -api ./api/product.api -dir ./api
   ```
3. 创建 proto 文件并生成 rpc 服务
   ```shell
   vim rpc/product.proto
   cd rpc
   goctl rpc protoc product.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.
   ```
4. 编写 product rpc 服务
   ```shell
   # 修改 product.yaml 配置文件
   vim rpc/etc/product.yaml
   # 添加 product model 依赖
   # 1.添加 Mysql 服务配置，CacheRedis 服务配置的实例化
   vim rpc/internal/config/config.go
   # 2.注册服务上下文 product model 的依赖
   vim rpc/internal/svc/servicecontext.go
   ```
5. 添加产品创建逻辑 Create
   ```shell
   vim rpc/internal/logic/createlogic.go
   ```
6. 添加产品详情逻辑 Detail
   ```shell
   vim rpc/internal/logic/detaillogic.go
   ```
7. 添加产品更新逻辑 Update
   ```shell
   vim rpc/internal/logic/updatelogic.go
   ```
8. 添加产品删除逻辑 Remove
   ```shell
   vim rpc/internal/logic/removelogic.go
   ```
9. 编写 product api 服务
   ```shell
   # 修改 product.yaml 文件，添加 product rpc 依赖
   # 1.添加 product rpc 服务配置
   vim api/etc/product.yaml
   # 2.添加 product rpc 服务配置的实例化
   vim api/internal/config/config.go
   # 3.注册服务上下文 product rpc 的依赖
   vim api/internal/svc/servicecontext.go
   ```
10. 添加产品创建逻辑 Create
   ```shell
   vim api/internal/logic/createlogic.go
   ```
11. 添加产品详情逻辑 Detail
   ```shell
   vim api/internal/logic/detaillogic.go
   ```
12. 添加产品更新逻辑 Update
   ```shell
   vim api/internal/logic/updatelogic.go
   ```
13. 添加产品删除逻辑 Remove
   ```shell
   vim api/internal/logic/removelogic.go
   ```
14. 启动 product rpc 服务
   ```shell
   $ cd service/product/rpc
   $ go run product.go -f etc/product.yaml
   ```
   > Starting rpc server at 127.0.0.1:9001...
15. 启动 product api 服务
   ```shell
   $ cd service/product/api
   $ go run product.go -f etc/product.yaml
   ```
   > Starting server at 0.0.0.0:8001...
