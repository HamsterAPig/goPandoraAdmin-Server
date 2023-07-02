# goPandroa-Admin

> 此部分程序为[goPandora](https://github.com/HamsterAPig/goPandora)项目的账号管理程序的后端服务部分，采用go语言配合gin框架写成



## 简介

该项目实现了一个基于 Gin 框架的 Web 应用程序，用于管理用户信息、自动登录信息和分享令牌。它提供了一组 API，可以执行用户信息的增删改查操作，管理自动登录信息以及处理分享令牌。

## 功能特性

- 用户信息管理：可以列出所有用户信息、添加新用户、删除用户以及修改用户信息。
- 自动登录信息管理：可以列出所有自动登录信息、添加新的自动登录信息、删除自动登录信息以及修改自动登录信息。
- 分享令牌管理：可以列出所有分享令牌、添加新的分享令牌、删除分享令牌以及修改分享令牌信息。

## 依赖项

- Gin 框架：用于构建 Web 应用程序的轻量级框架。
- Spf13/Viper：用于读取和解析配置文件。
- Gin-Contrib/Cors：用于处理跨域请求。

## 配置

可以通过配置文件或命令行参数来配置应用程序。以下是可配置的参数：

- `listen`：监听地址，默认为 `127.0.0.1:8080`。
- `proxy`：代理地址。
- `database`：数据库路径，默认为 `./goPandora.db`。
- `debug-level`：日志等级。
- `allow-cors`：是否允许跨域请求。
- `enable-uuid-uri`：是否在 API 路径中添加 UUID。



----

**注意**：本项目并没有使用API鉴权，故而需要用户自行添加授权信息或者是在内网使用

**强烈建议启用`enable-uuid-uri`选项避免后端API被嗅探**

## 安装与运行

1. 克隆项目代码：

   ```shell
   git clone https://github.com/HamsterAPig/goPandoraAdmin-Server
   ```

2. 进入项目目录：

   ```shell
   cd goPandoraAdmin-Server
   ```

3. 安装依赖项：

   ```shell
   go mod download
   ```

4. 运行应用程序：

   ```shell
   go run main.go
   ```

5. 应用程序将在指定的监听地址启动，并可以通过浏览器或 API 调用来访问。

6. 此外，本项目针对该后端API写了一个简陋的前端，见[goPandoraAdmin-Web](https://github.com/HamsterAPig/goPandoraAdmin-Web)

## API 文档

### 用户信息管理

- 列出所有用户信息：`GET /api/v1/users`
- 添加新用户：`POST /api/v1/users`
- 删除用户：`DELETE /api/v1/users/:userID`
- 修改用户信息：`PATCH /api/v1/users/:userID`

### 自动登录信息管理

- 列出所有自动登录信息：`GET /api/v1/auto-login-infos`
- 添加新的自动登录信息：`POST /api/v1/auto-login-infos`
- 删除自动登录信息：`DELETE /api/v1/auto-login-infos/:UUID`
- 修改自动登录信息：`PATCH /api/v1/auto-login-infos/:UUID`

### 分享令牌管理

- 列出所有分享令牌：`GET /api/v1/share-tokens`
- 添加新的分享令牌：`POST /api/v1/share-tokens`
- 删除分享令牌：`DELETE /api/v1/share-tokens/:fk`

- 修改分享令牌信息：`PATCH /api/v1/share-tokens/:fk`

可以使用这些 API 来与应用程序交互，执行各种操作，如管理用户信息、自动登录信息和分享令牌。

## 贡献

欢迎对该项目进行贡献！如果您想为该项目贡献代码，请先 Fork 该仓库，然后提交 Pull Request。我们非常感谢您的贡献！
