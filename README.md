# apiGateway

## 概述

- Gateway项目：[SchrodingerwithCat/apiGateway: 2023年7月，CloudWeGo 课程项目API网关项目 (github.com)](https://github.com/SchrodingerwithCat/apiGateway)
- 编程语言：go语言
- API-Gateway服务：hertz框架
- RPC服务：Kitex Service框架
- Registry Center：etcd
- RPC 协议：thrift

### 项目结构

```
.
├── http
│   ├── biz
│   │   ├── clientprovider
│   │   │   └── client_provider.go
│   │   ├── handler
│   │   │   └── demo
│   │   │       ├── student_service.go
│   │   │       └── teacher_service.go
│   │   ├── idl
│   │   │   ├── item.thrift
│   │   │   ├── student.thrift
│   │   │   └── teacher.thrift
│   │   ├── model
│   │   │   └── demo
│   │   │       └── item.go
│   │   └── router
│   │       ├── demo
│   │       │   ├── item.go
│   │       │   └── middleware.go
│   │       └── register.go
│   ├── build.sh
│   ├── go.mod
│   ├── go.sum
│   ├── http
│   ├── main.go
│   ├── router_gen.go
│   ├── router.go
│   └── script
│       └── bootstrap.sh
└── rpc
    ├── student_service
    │   ├── build.sh
    │   ├── go.mod
    │   ├── go.sum
    │   ├── handler.go
    │   ├── idl
    │   │   └── student.thrift
    │   ├── kitex_gen
    │   │   └── demo
    │   │       ├── k-consts.go
    │   │       ├── k-student.go
    │   │       ├── student.go
    │   │       ├── student_item.go
    │   │       └── studentservice
    │   │           ├── client.go
    │   │           ├── invoker.go
    │   │           ├── server.go
    │   │           └── studentservice.go
    │   ├── kitex_info.yaml
    │   ├── main.go
    │   ├── script
    │   │   └── bootstrap.sh
    │   └── student.db
    └── teacher_service
        ├── build.sh
        ├── go.mod
        ├── go.sum
        ├── handler.go
        ├── idl
        │   └── teacher.thrift
        ├── kitex_gen
        │   └── demo
        │       ├── k-consts.go
        │       ├── k-teacher.go
        │       ├── teacher.go
        │       ├── teacher_item.go
        │       └── teacherservice
        │           ├── client.go
        │           ├── invoker.go
        │           ├── server.go
        │           └── teacherservice.go
        ├── kitex_info.yaml
        ├── main.go
        ├── script
        │   └── bootstrap.sh
        └── teacher.db

```

## 部署

1. 打开etcd注册中心

```shell
etcd --log-level debug
```
2. 打开api网关服务
   - 从根目录打开http文件夹。


```shell
## 构建可执行文件
go build

## 运行可执行文件
./http
```

3. 运行`student_service`和`teacher_service`服务
   - 进入./rpc/student_service文件夹
   
     ```shell
     ## 构建可执行文件
     go build
     
     ## 运行可执行文件
     ./student_service
     ```
   
   - 进入./rpc/teacher_service文件夹
   
     ```shell
     ## 构建可执行文件
     go build
     
     ## 运行可执行文件
     ./teacher_service
     ```

这样就完成了部署

### 检验部署完成

可以对网关发送HttpRequest来检查是否完成部署，如

```shell
curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8888/student/add-student-info -d '{"id": 100, "name":"Emma", "college": {"name": "software college", "address": "逸夫"}, "email": ["emma@nju.com"]}'
```

若收到

```shel
"{\"success\":true,\"message\":\"student_service: Student({Id:100 Name:Emma College:College({Name:software college Address:逸夫}) Email:[emma@nju.co}) 注册成功。\"}"j
```

则部署成功
