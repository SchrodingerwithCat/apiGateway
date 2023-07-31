package main

import (
	studentservice "github.com/SchrodingerwithCat/apiGateway/rpc/student_service/kitex_gen/demo/studentservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net"

	demo "github.com/SchrodingerwithCat/apiGateway/rpc/student_service/kitex_gen/demo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":9999")
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})

	if err != nil {
		log.Fatal(err)
	}

	handler := new(StudentServiceImpl)
	handler.InitDB()
	svr := studentservice.NewServer(handler, server.WithRegistry(r), server.WithServiceAddr(addr), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "student_service",
		Tags:        map[string]string{"Cluster": "student"},
	}))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}

// eg: 初始化db，注意服务启动时初始化
func (s *StudentServiceImpl) InitDB() {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// drop table
	db.Migrator().DropTable(demo.StudentItem{})
	// create table
	err = db.Migrator().CreateTable(demo.StudentItem{})
	if err != nil {
		panic(err)
	}
	s.db = db
}
