package main

import (
	"context"
	"fmt"
	demo "github.com/SchrodingerwithCat/apiGateway/rpc/teacher_service/kitex_gen/demo"
)

// TeacherServiceImpl implements the last service interface defined in the IDL.
type TeacherServiceImpl struct{}

var teachers = make(map[int32]*demo.Teacher)

// TeacherRegister implements the TeacherServiceImpl interface.
func (s *TeacherServiceImpl) TeacherRegister(ctx context.Context, teacher *demo.Teacher) (resp *demo.RegisterResp, err error) {
	// TODO: Your code here...
	if _, ok := teachers[teacher.Id]; ok {
		return &demo.RegisterResp{
			Success: false,
			Message: fmt.Sprintf("ID %d 已经注册过了。", teacher.Id),
		}, nil
	} else {
		teachers[teacher.Id] = &demo.Teacher{
			Id:      teacher.Id,
			Name:    teacher.Name,
			College: teacher.College,
			Email:   teacher.Email,
		}
		return &demo.RegisterResp{
			Success: true,
			Message: fmt.Sprintf("%s 注册成功。", teacher.String()),
		}, nil
	}
}

// TeacherQuery implements the TeacherServiceImpl interface.
func (s *TeacherServiceImpl) TeacherQuery(ctx context.Context, req *demo.QueryReq) (resp *demo.Teacher, err error) {
	// TODO: Your code here...
	resp = teachers[req.Id]
	if resp == nil {
		resp = &demo.Teacher{
			Id:      req.Id,
			Name:    "Not Found",
			College: &demo.College{},
			Email:   []string{"Not Found"},
		}
	}
	return
}
