# apiGateway

# 1. 小组成员

| 成员   | 学号      |
| ------ | --------- |
| 周密   | 211098249 |
| 陈宇航 | 211850104 |
| 曹嘉琪 | 211250118 |

## 概述

- Gateway项目：[SchrodingerwithCat/apiGateway: 2023年7月，CloudWeGo 课程项目API网关项目 (github.com)](https://github.com/SchrodingerwithCat/apiGateway)
- 编程语言：go语言
- API-Gateway服务：hertz框架
- RPC服务：Kitex Service框架
- Registry Center：etcd
- RPC 协议：thrift

# 2. 目录结构

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

# 3.  部署说明

## 1. 环境配置

```
//安装thrift
go install github.com/cloudwego/thriftgo@latest
//安装hz
go install github.com/cloudwego/hertz/cmd/hz@latest
//安装kitex
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
```

## 2. 运行说明

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

# 4.测试部分

## 1.引言

测试文档由软件设计说明所驱动。测试用于验证模块单元实现了模块设计中定义的规格。一个完整的单元测试说明应该包含白盒测试和黑盒测试。测试验证程序应该执行的工作，测试验证程序不应该执行的工作。

1.1编写目的

通过测试尽可能的找出项目中的错误，并加以纠正。测试不仅是最后的复审，更是保证软件质量的关键。 简单地说就是想尽一切方法尝试“破坏”它，这样才能找出失败与不足之处，最终的任务就是构造更高质量的软件产品。

1.2参考文献

1. IEEE标准
2. 《软件工程与计算(卷二):软件开发的技术基础》刘钦、丁二玉著
3. MSCS软件需求规格文档

## 2.功能测试

```
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

// Student represents the student data structure.
type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	College struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	} `json:"college"`
	Email []string `json:"email"`
}

// aUrlRegister is the URL for registering a student.
const aUrlRegister = "http://127.0.0.1:8888/gateway/StudentService/Register"
const aUrlEcho = "http://localhost:8888/gateway/EchoService/Echo"

// generateRandomStudent generates a random student for testing.
func generateRandomStudent() Student {
	st := Student{
		ID:   rand.Intn(1000) + 1 + 32,
		Name: fmt.Sprintf("Student%d", rand.Intn(100)),
		College: struct {
			Name    string `json:"name"`
			Address string `json:"address"`
		}{
			Name:    fmt.Sprintf("College%d", rand.Intn(10)),
			Address: fmt.Sprintf("Address%d", rand.Intn(100)),
		},
	}
	emailCount := rand.Intn(5) + 1 // Generate a random number between 1 and 5
	st.Email = make([]string, emailCount)
	for i := 0; i < emailCount; i++ {
		st.Email[i] = fmt.Sprintf("email%d@test.com", rand.Intn(100))
	}

	return st
}

type RegisterResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// BenchmarkStudentServiceRegister tests the performance of the StudentService Register function.
func BenchmarkStudentServiceRegister(b *testing.B) {
	for i := 0; i < b.N; i++ {
		student := generateRandomStudent()

		bytesData, err := json.Marshal(student)
		if err != nil {
			b.Fatalf("Error marshalling data: %v", err)
		}

		reader := bytes.NewReader(bytesData)
		request, err := http.NewRequest("POST", aUrlRegister, reader)

		if err != nil {
			b.Fatalf("Error creating request: %v", err)
		}
		defer request.Body.Close()

		request.Header.Set("Content-Type", "application/json;charset=UTF-8")

		client := http.Client{}
		startTime := time.Now()
		resp, err := client.Do(request)
		if err != nil {
			b.Fatalf("Error sending request: %v", err)
		}
		elapsed := time.Since(startTime)

		defer resp.Body.Close()

		var response RegisterResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			b.Fatalf("Error decoding response: %v", err)
		}

		if !response.Success {
			b.Fatalf("Registration failed for student: %+v", student)
		}

		b.Logf("Student registration successful: %+v (Elapsed time: %v)", student, elapsed)
	}
}

type EchoRequest struct {
	Message string `json:"message"`
}

type EchoResponse struct {
	Message string `json:"message"`
}

// BenchmarkEchoService tests the performance of the EchoService Echo function.
func BenchmarkEchoService(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Create a random message for testing
		message := fmt.Sprintf("hello%d, Cloudwego!", rand.Intn(100))

		// Create the request payload
		request := EchoRequest{
			Message: message,
		}

		bytesData, err := json.Marshal(request)
		if err != nil {
			b.Fatalf("Error marshalling data: %v", err)
		}

		reader := bytes.NewReader(bytesData)
		request_2, err := http.NewRequest("POST", aUrlEcho, reader)
		if err != nil {
			b.Fatalf("Error creating request: %v", err)
		}
		defer request_2.Body.Close()

		request_2.Header.Set("Content-Type", "application/json;charset=UTF-8")

		client := http.Client{}
		startTime := time.Now()
		resp, err := client.Do(request_2)
		if err != nil {
			b.Fatalf("Error sending request: %v", err)
		}
		elapsed := time.Since(startTime)

		defer resp.Body.Close()

		var response EchoResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			b.Fatalf("Error decoding response: %v", err)
		}

		if response.Message != message {
			b.Fatalf("Unexpected response: expected %q, got %q", message, response.Message)
		}

		b.Logf("EchoService successful: Request=%q, Response=%q (Elapsed time: %v)", message, response.Message, elapsed)
	}
}

```



本次测试主要为验证网关对于服务基本功能是否成功支持，各模块功能是否符合预期，并且通过Golang性能测试与`Apache Benchmark`压力测试初步验证网关性能，并为后期优化提供参考。

测试环境

| 类型 | 说明                                                       |
| ---- | ---------------------------------------------------------- |
| OS   | wsl2                                                       |
| CPU  | AMD Ryzen 7 5800H with Radeon Graphics 3.20 GHz，8核16线程 |

### 测试步骤

#### 服务启动

##### 启动etcd

```
etcd --log-level debug
```

##### 启动HTTP Server

```
// 在 api-gateway/hertz-http-server 目录下
go run .
```

##### 启动响应微服务

```
// 在相应微服务文件目录下，如 api-gateway/microserviceaddition-service 下
go run .

```

结果表明网关成功提供了对服务的支持，能够正确接受与响应POST请求，并且根据请求路由确认目标服务和方法，并且可以根据相应IDL文件与微服务中的处理逻辑完成相应业务

## 3.Benchmark性能测试

运行

```
go test -bench=. echo_test.go
```

### 串行测试

![image-20230731212427271](C:\Users\Lenovo\AppData\Roaming\Typora\typora-user-images\image-20230731212427271.png)

![image-20230731213541522](C:\Users\Lenovo\AppData\Roaming\Typora\typora-user-images\image-20230731213541522.png)



测试报告解读：

总共用16个CPU核心，但是本次是串行测试，执行的测试是20097982， 每次耗时57.26

### 并行测试

![image-20230731213507294](C:\Users\Lenovo\AppData\Roaming\Typora\typora-user-images\image-20230731213507294.png)

总共用16个CPU核心，，执行的测试是74950425， 每次耗时16.66

## 5.Apache JMeter进行测试

下面使用测试工具， Apache JMeter 5.6.2 进行测试

测试配置

让10个线程循环运行70次

![img](file:///C:\Users\Lenovo\Documents\Tencent Files\3065718529\Image\Group2\BH\IM\BHIMH9F`LHCI$NCY96]6460.png)

运行结果

![image-20230731231109118](C:\Users\Lenovo\AppData\Roaming\Typora\typora-user-images\image-20230731231109118.png)

测试配置

让10个线程循环运行700次

![image-20230731231123942](C:\Users\Lenovo\AppData\Roaming\Typora\typora-user-images\image-20230731231123942.png)

![image-20230731231154112](C:\Users\Lenovo\AppData\Roaming\Typora\typora-user-images\image-20230731231154112.png)

测试使用命令：

- 进入JMeter目录
    - cd .\apache-jmeter-5.6.2\bin
- 生成html测试报告
    - .\jmeter -n -t ..\jmx\cloudwego.jmx -l ..\results\log -e -o ..\results\output

性能结果分析

1. 第一次实验中10个线程下， 吞吐量总体达到了35.2/min
2. 第二次实验后10个线程下，吞吐量总体值为5.8/sec