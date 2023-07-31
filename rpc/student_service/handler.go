package main

import (
	"context"
	"fmt"
	demo "github.com/SchrodingerwithCat/apiGateway/rpc/student_service/kitex_gen/demo"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct{}

var students = make(map[int32]*demo.Student)

// StuRegister implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) StuRegister(ctx context.Context, student *demo.Student) (resp *demo.RegisterResp, err error) {
	// TODO: Your code here...
	if _, ok := students[student.Id]; ok {
		return &demo.RegisterResp{
			Success: false,
			Message: fmt.Sprintf("ID %d 已经注册过了。", student.Id),
		}, nil
	} else {
		students[student.Id] = &demo.Student{
			Id:      student.Id,
			Name:    student.Name,
			College: student.College,
			Email:   student.Email,
		}
		return &demo.RegisterResp{
			Success: true,
			Message: fmt.Sprintf("%s 注册成功。", student.String()),
		}, nil
	}
}

// StuQuery implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) StuQuery(ctx context.Context, req *demo.QueryReq) (resp *demo.Student, err error) {
	// TODO: Your code here...
	resp = students[req.Id]
	if resp == nil {
		resp = &demo.Student{
			Id:      req.Id,
			Name:    "Not Found",
			College: &demo.College{},
			Email:   []string{"Not Found"},
		}
	}
	return
}
