# VulOj_BE

- 这里是VulOj项目的后端部分

## 项目框架

- 本项目使用gin框架，基于gin脚手架（[create-gin-app](https://gin-gonic.com/)）搭建

## 技术依赖

- [gin, gin-Context](ttps://gin-gonic.com/)  
    gin框架的官网
- [gorm](https://github.com/go-gorm/gorm)  
    gorm数据库框架
- [mysql](https://www.mysql.com/cn/)  
    mysql数据库用于静态数据存储
- [redis](https://redis.io/)  
    redis用于缓存数据存储
- [grpc](https://grpc.io/)  
    grpc用于数据传输（邮件服务部分的分离）  
- [docker](https://www.docker.com/)  
    使用docker将靶机分离成不同虚拟环境并动态创建  
## 项目结构

- **目前本项目的开发工作均在1.0.0分支进行，若需要源码编译等操作，请切换到1.0.0分支**

- 本项目结构如下：
    ```
    VulOj_BE\
        directory\            // VulOj靶场存储目录
        models\               // VulOj项目数据库表和基本数据类型
        public\               // 存放全局资源
        pkg\                  //项目源码
            consts\           //静态设置
            middleware\       //中间件
            rouer\            //路由树和前后交互api相关函数
            services\         //后端逻辑函数和数据库交互函数
            utils\            //格式序列化函数
    ```

## 安装

- **本项目基于golang 1.15.3，请确保你已经在官网安装尽量新版本的[golang](https://go.dev/)**

### 拉取项目源码

- 请使用git版本管理工具拉取仓库
    ``` sh
    git clone git@github.com:VulOj/VulOJ_BE.git
    cd VulOJ_BE
    ```
- 如果在拉取本项目时还未完工，main分支内容可能不完整，请手动切换到最新（1.0.0）分支
    ``` sh
    git checkout 1.0.0
    ```

### 获取项目依赖

- 获取项目相关依赖。该操作已在项目中配置完成，只需要执行如下命令
    ``` sh
    go mod init VBS
    go mod tidy/go mod download
    ```
- **注：本操作及之后操作请确保在 `\VulOJ_BE` 目录下进行**

### 本地运行

- 本地运行项目，该操作会在本地的8080端口开启虚拟服务，可以在[http://localhost:8080/](http://localhost:8080/)实时预览项目。若出现端口占用等需要修改默认端口的情况,可以直接在main.go中修改端口，后续版本会推出用户初始化配置安装流程，方便置顶后端响应端口。
    ``` sh
    go run main.go
    ```
- 该命令执行后，可以使用本地浏览器访问8080端口[http://localhost:8080/](http://localhost:8080/)来验证是否成功运行后端服务


## 授权

- 本项目最终解释权归VulOj团队所有
