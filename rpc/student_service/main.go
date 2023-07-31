package main

import (
	demo "github.com/SchrodingerwithCat/apiGateway/rpc/student_service/kitex_gen/demo/studentservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"

	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":9999")
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})

	if err != nil {
		log.Fatal(err)
	}

	handler := new(StudentServiceImpl)
	svr := demo.NewServer(handler, server.WithRegistry(r), server.WithServiceAddr(addr), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "student_service",
		Tags:        map[string]string{"Cluster": "student"},
	}))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
