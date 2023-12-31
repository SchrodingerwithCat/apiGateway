// Code generated by Kitex v0.6.1. DO NOT EDIT.

package teacherservice

import (
	"context"
	demo "github.com/SchrodingerwithCat/apiGateway/rpc/teacher_service/kitex_gen/demo"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return teacherServiceServiceInfo
}

var teacherServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "TeacherService"
	handlerType := (*demo.TeacherService)(nil)
	methods := map[string]kitex.MethodInfo{
		"TeacherRegister": kitex.NewMethodInfo(teacherRegisterHandler, newTeacherServiceTeacherRegisterArgs, newTeacherServiceTeacherRegisterResult, false),
		"TeacherQuery":    kitex.NewMethodInfo(teacherQueryHandler, newTeacherServiceTeacherQueryArgs, newTeacherServiceTeacherQueryResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "demo",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func teacherRegisterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*demo.TeacherServiceTeacherRegisterArgs)
	realResult := result.(*demo.TeacherServiceTeacherRegisterResult)
	success, err := handler.(demo.TeacherService).TeacherRegister(ctx, realArg.Teacher)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTeacherServiceTeacherRegisterArgs() interface{} {
	return demo.NewTeacherServiceTeacherRegisterArgs()
}

func newTeacherServiceTeacherRegisterResult() interface{} {
	return demo.NewTeacherServiceTeacherRegisterResult()
}

func teacherQueryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*demo.TeacherServiceTeacherQueryArgs)
	realResult := result.(*demo.TeacherServiceTeacherQueryResult)
	success, err := handler.(demo.TeacherService).TeacherQuery(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTeacherServiceTeacherQueryArgs() interface{} {
	return demo.NewTeacherServiceTeacherQueryArgs()
}

func newTeacherServiceTeacherQueryResult() interface{} {
	return demo.NewTeacherServiceTeacherQueryResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) TeacherRegister(ctx context.Context, teacher *demo.Teacher) (r *demo.RegisterResp, err error) {
	var _args demo.TeacherServiceTeacherRegisterArgs
	_args.Teacher = teacher
	var _result demo.TeacherServiceTeacherRegisterResult
	if err = p.c.Call(ctx, "TeacherRegister", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) TeacherQuery(ctx context.Context, req *demo.QueryReq) (r *demo.Teacher, err error) {
	var _args demo.TeacherServiceTeacherQueryArgs
	_args.Req = req
	var _result demo.TeacherServiceTeacherQueryResult
	if err = p.c.Call(ctx, "TeacherQuery", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
