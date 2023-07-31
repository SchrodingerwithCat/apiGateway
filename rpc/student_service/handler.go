package main

import (
	"context"
	"errors"
	"fmt"
	demo "github.com/SchrodingerwithCat/apiGateway/rpc/student_service/kitex_gen/demo"
	"gorm.io/gorm"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct {
	db *gorm.DB
}

// StuRegister implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) StuRegister(ctx context.Context, student *demo.Student) (resp *demo.RegisterResp, err error) {
	// TODO: Your code here...
	var stuRes *demo.StudentItem
	result := s.db.Table("student_items").First(&stuRes, student.Id) // result.Error 是失败信息 , result.Error 是 gorm.ErrRecordNotFound 表示没有查询结果，eg: if errors.Is(result.Error, gorm.ErrRecordNotFound) { // 查询结果为空}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 为空，ID合法
		result = s.db.Table("student_items").Create(demo.Student2StudentItem(student))
		resp = &demo.RegisterResp{
			Success: true,
			Message: fmt.Sprintf("%s 注册成功。", student.String()),
		}
	} else {
		// 不为空，ID不合法
		resp = &demo.RegisterResp{
			Success: false,
			Message: fmt.Sprintf("ID %d 已经注册过了。", student.Id),
		}
	}
	return
}

// StuQuery implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) StuQuery(ctx context.Context, req *demo.QueryReq) (resp *demo.Student, err error) {
	// TODO: Your code here...
	var stuRes *demo.StudentItem
	result := s.db.Table("student_items").First(&stuRes, req.Id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		resp = &demo.Student{
			Id:      req.Id,
			Name:    "Not Found",
			College: &demo.College{},
			Email:   []string{"Not Found"},
		}
	} else {
		resp = demo.StudentItem2Student(stuRes)
	}
	return
}
