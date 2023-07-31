package main

import (
	"context"
	"errors"
	"fmt"
	demo "github.com/SchrodingerwithCat/apiGateway/rpc/teacher_service/kitex_gen/demo"
	"gorm.io/gorm"
)

// TeacherServiceImpl implements the last service interface defined in the IDL.
type TeacherServiceImpl struct {
	db *gorm.DB
}

// TeacherRegister implements the TeacherServiceImpl interface.
func (s *TeacherServiceImpl) TeacherRegister(ctx context.Context, teacher *demo.Teacher) (resp *demo.RegisterResp, err error) {
	// TODO: Your code here...
	var teacherRes *demo.TeacherItem
	result := s.db.Table("teacher_items").First(&teacherRes, teacher.Id) // result.Error 是失败信息 , result.Error 是 gorm.ErrRecordNotFound 表示没有查询结果，eg: if errors.Is(result.Error, gorm.ErrRecordNotFound) { // 查询结果为空}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 为空，ID合法
		result = s.db.Table("teacher_items").Create(demo.Teacher2TeacherItem(teacher))
		resp = &demo.RegisterResp{
			Success: true,
			Message: fmt.Sprintf("teacher_service: %s 注册成功。", teacher.String()),
		}
	} else {
		// 不为空，ID不合法
		resp = &demo.RegisterResp{
			Success: false,
			Message: fmt.Sprintf("teacher_service: ID %d 已经注册过了。", teacher.Id),
		}
	}
	return
}

// TeacherQuery implements the TeacherServiceImpl interface.
func (s *TeacherServiceImpl) TeacherQuery(ctx context.Context, req *demo.QueryReq) (resp *demo.Teacher, err error) {
	// TODO: Your code here...
	var teacherRes *demo.TeacherItem
	result := s.db.Table("teacher_items").First(&teacherRes, req.Id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		resp = &demo.Teacher{
			Id:      req.Id,
			Name:    "Not Found",
			College: &demo.College{},
			Email:   []string{"Not Found"},
		}
	} else {
		resp = demo.TeacherItem2Teacher(teacherRes)
	}
	return
}
