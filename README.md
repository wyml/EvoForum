# Evo Forum

Evo Forum：基于 RESTful API 最佳实践的论坛程序，简约、轻快、Material You!

<br />

<p align="center">
  <a href="https://github.com/wyml/EvoForum/">
    LOGO
  </a>

<h3 align="center">EvoForum</h3>
  <p align="center">
    一个简约、轻快的轻论坛项目，使用 Material You Desing
    <br />
    <a href="https://github.com/wyml/EvoForum"><strong>探索本项目的文档 »</strong></a>
    <br />
    <br />
    <a href="https://github.com/wyml/EvoForum">查看Demo</a>
    ·
    <a href="https://github.com/wyml/EvoForum/issues">报告Bug</a>
    ·
    <a href="https://github.com/wyml/EvoForum/issues">提出新特性</a>
  </p>

</p>


本篇README.md面向用户以及开发者

## 目录

- [上手指南](#上手指南)
  - [开发前的配置要求](#开发前的配置要求)
  - [安装步骤](#安装步骤)
- [文件目录说明](#文件目录说明)
- [开发的架构](#开发的架构)
- [部署](#部署)
- [使用到的框架](#使用到的框架)
- [贡献者](#贡献者)
  - [如何参与开源项目](#如何参与开源项目)
- [版本控制](#版本控制)
- [作者](#作者)
- [鸣谢](#鸣谢)

### 上手指南

###### 开发前的配置要求

1. Go >= 1.7.0
2. Mysql >= 5.7.26
3. Redis >= 3.0.504

###### **安装步骤**

1. 克隆本仓库

```sh
git clone https://github.com/wyml/EvoForum.git
```

2. 复制 `.env.example` 重命名为 `.env`
3. 修改 `.env` 中的程序配置、数据库等相关信息
4. 根目录下执行
5. 暂未完成

### 文件目录说明
eg:

```
.├── app                            // 程序具体逻辑代码
│   ├── cmd                         // 命令
│   │   ├── cache.go                
│   │   ├── cmd.go
│   │   ├── key.go
│   │   ├── make                    // make 命令及子命令
│   │   │   ├── make.go
│   │   │   ├── make_apicontroller.go
│   │   │   ├── make_cmd.go
│   │   │   ├── make_factory.go
│   │   │   ├── make_migration.go
│   │   │   ├── make_model.go
│   │   │   ├── make_policy.go
│   │   │   ├── make_request.go
│   │   │   ├── make_seeder.go
│   │   │   └── stubs               // make 命令的模板
│   │   │       ├── apicontroller.stub
│   │   │       ├── cmd.stub
│   │   │       ├── factory.stub
│   │   │       ├── migration.stub
│   │   │       ├── model
│   │   │       │   ├── model.stub
│   │   │       │   ├── model_hooks.stub
│   │   │       │   └── model_util.stub
│   │   │       ├── policy.stub
│   │   │       ├── request.stub
│   │   │       └── seeder.stub
│   │   ├── migrate.go
│   │   ├── play.go
│   │   ├── seed.go
│   │   └── serve.go
│   ├── http                        // http 请求处理逻辑
│   │   ├── controllers             // 控制器，存放 API 和视图控制器
│   │   │   ├── api                 // API 控制器，支持多版本的 API 控制器
│   │   │   │   └── v1              // v1 版本的 API 控制器
│   │   │   │       ├── users_controller.go
│   │   │   │       └── ...
│   │   └── middlewares             // 中间件
│   │       ├── auth_jwt.go
│   │       ├── guest_jwt.go
│   │       ├── limit.go
│   │       ├── logger.go
│   │       └── recovery.go
│   ├── models                      // 数据模型
│   │   ├── user                    // 单独的模型目录
│   │   │   ├── user_hooks.go       // 模型钩子文件
│   │   │   ├── user_model.go       // 模型主文件
│   │   │   └── user_util.go        // 模型辅助方法
│   │   └── ...
│   ├── policies                    // 授权策略目录
│   │   ├── category_policy.go
│   │   └── ...
│   └── requests                    // 请求验证目录（支持表单、标头、Raw JSON、URL Query）
│       ├── validators              // 自定的验证规则
│       │   ├── custom_rules.go
│       │   └── custom_validators.go
│       ├── user_request.go
│       └── ...
├── bootstrap                       // 程序模块初始化目录
│   ├── app.go  
│   ├── cache.go
│   ├── database.go
│   ├── logger.go
│   ├── redis.go
│   └── route.go
├── config                          // 配置信息目录
│   ├── app.go
│   ├── captcha.go
│   ├── config.go
│   ├── database.go
│   ├── jwt.go
│   ├── log.go
│   ├── mail.go
│   ├── pagination.go
│   ├── redis.go
│   ├── sms.go
│   └── verifycode.go
├── database                        // 数据库相关目录
│   ├── database.db                 // sqlite 数据文件（加入到 .gitignore 中）
│   ├── factories                   // 模型工厂目录
│   │   ├── user_factory.go
│   │   └── ...
│   ├── migrations                  // 数据库迁移目录
│   │   ├── 2021_12_21_102259_create_users_table.go
│   │   ├── 2021_12_21_102340_create_categories_table.go
│   │   └── ...
│   └── seeders                     // 数据库填充目录
│       ├── users_seeder.go
│       ├── ...
├── pkg                             // 内置辅助包
│   ├── app
│   ├── auth
│   ├── cache
│   ├── captcha
│   ├── config
│   └── ...
├── public                          // 静态文件存放目录
│   ├── css
│   ├── js
│   └── uploads                     // 用户上传文件目录
│       └── avatars                 // 用户上传头像目录
├── routes                          // 路由
│   ├── api.go
│   └── web.go
├── storage                         // 内部存储目录
│   ├── app
│   └── logs                        // 日志存储目录
│       ├── 2021-12-28.log
│       ├── 2021-12-29.log
│       ├── 2021-12-30.log
│       └── logs.log
└── tmp                             // air 的工作目录
├── .env                            // 环境变量文件
├── .env.example                    // 环境变量示例文件
├── .gitignore                      // git 配置文件
├── .air.toml                       // air 配置文件
├── .editorconfig                   // editorconfig 配置文件
├── go.mod                          // Go Module 依赖配置文件
├── go.sum                          // Go Module 模块版本锁定文件
├── main.go                         // EvoForum 程序主入口
├── Makefile                        // 自动化命令文件
├── README.md                       // 项目 readme

```


### 开发的架构

请阅读[ARCHITECTURE.md](https://github.com/wyml/EvoForum/blob/master/ARCHITECTURE.md) 查阅为该项目的架构。

### 部署

暂无

### 使用到的框架

- [gin](https://gin-gonic.com/zh-cn/)
- [cobra](https://github.com/spf13/cobra)
- [viper](https://github.com/spf13/viper)
- [zap](https://github.com/uber-go/zap)
- [gorm](https://github.com/go-gorm/gorm)

### 贡献者

请阅读**CONTRIBUTING.md** 查阅为该项目做出贡献的开发者。

#### 如何参与开源项目

贡献使开源社区成为一个学习、激励和创造的绝佳场所。你所作的任何贡献都是**非常感谢**的。


1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



### 版本控制

该项目使用Git进行版本管理。您可以在repository参看当前可用版本。

### 作者

lanjiabin.self@qq.com


*您也可以在贡献者名单中参看所有参与该项目的开发者。*

### 版权说明

该项目签署了Apache 2.0 授权许可，详情请参阅 [LICENSE.txt](https://github.com/wyml/EvoForum/blob/master/LICENSE.txt)

### 鸣谢


- [gin](https://gin-gonic.com/zh-cn/)
- [cobra](https://github.com/spf13/cobra)
- [viper](https://github.com/spf13/viper)
- [zap](https://github.com/uber-go/zap)
- [gorm](https://github.com/go-gorm/gorm)

<!-- links -->
[your-project-path]:wyml/EvoForum
[contributors-shield]: https://img.shields.io/github/contributors/wyml/EvoForum.svg?style=flat-square
[contributors-url]: https://github.com/wyml/EvoForum/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/wyml/EvoForum.svg?style=flat-square
[forks-url]: https://github.com/wyml/EvoForum/network/members
[stars-shield]: https://img.shields.io/github/stars/wyml/EvoForum.svg?style=flat-square
[stars-url]: https://github.com/wyml/EvoForum/stargazers
[issues-shield]: https://img.shields.io/github/issues/wyml/EvoForum.svg?style=flat-square
[issues-url]: https://img.shields.io/github/issues/wyml/EvoForum.svg
[license-shield]: https://img.shields.io/github/license/wyml/EvoForum.svg?style=flat-square
[license-url]: https://github.com/wyml/EvoForum/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/shaojintian



