package main

import (
	teacherservice "github.com/SchrodingerwithCat/apiGateway/rpc/teacher_service/kitex_gen/demo/teacherservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	demo "github.com/SchrodingerwithCat/apiGateway/rpc/teacher_service/kitex_gen/demo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":9998")
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})

	if err != nil {
		log.Fatal(err)
	}

	handler := new(TeacherServiceImpl)
	handler.InitDB()
	svr := teacherservice.NewServer(handler, server.WithRegistry(r), server.WithServiceAddr(addr), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "teacher_service",
		Tags:        map[string]string{"Cluster": "teacher"},
	}))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

func (s *TeacherServiceImpl) InitDB() {
	db, err := gorm.Open(sqlite.Open("teacher.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// drop table
	db.Migrator().DropTable(demo.TeacherItem{})
	// create table
	err = db.Migrator().CreateTable(demo.TeacherItem{})
	if err != nil {
		panic(err)
	}
	s.db = db
}
