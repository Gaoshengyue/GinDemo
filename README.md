# Gin Demo例子

## 说明
用于对Gin框架的基础学习，整体合理框架架构搭建

## 项目依赖

```shell
go install
```

## 项目结构

```
gindemo
    ├───config   //配置
    ├───controllers  //控制器
    ├───db    //数据库操作
    ├───log    //日志存储
    ├───middlewares  //中间件
    ├───repositories //数据操作实体
    ├───schema   //相关实体定义
    ├───services  //服务相关
    └───utils  //工具
```

.env内容:

```
#mysql
MYSQL_HOST=
MYSQL_PORT=
MYSQL_USER=
MYSQL_PASSWORD=
MYSQL_DB=

#redis
REDIS_HOST=
REDIS_PORT=
REDIS_DB=
REDIS_PASSWORD=

```

需要在目录下创建env加载配置项

## 项目启动
推荐
```
docker-compose up -d
```

本地
```
go run main.go -p 8100
```

打包
```
docker build -t {服务}:{版本} .

```