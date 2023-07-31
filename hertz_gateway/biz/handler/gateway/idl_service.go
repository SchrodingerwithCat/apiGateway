// Code generated by hertz generator.

package gateway

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	idlManager "hertz.demo/biz/idl"
	gateway "hertz.demo/biz/model/gateway"
)

// AddService .
// @router /add-service [POST]
func AddService(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.Service
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(gateway.SuccessResp)
	idlManager.AddService(req)

	resp = &gateway.SuccessResp{
		Success: true,
		Message: "Add " + req.ServiceName + " success",
	}
	c.JSON(consts.StatusOK, resp)
}

// DeleteService .
// @router /delete-service [POST]
func DeleteService(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.ServiceReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(gateway.SuccessResp)
	idlManager.DeleteService(req.ServiceName)

	resp = &gateway.SuccessResp{
		Success: true,
		Message: "Delete " + req.ServiceName + " success",
	}
	c.JSON(consts.StatusOK, resp)
}

// UpdateService .
// @router /update-service [POST]
func UpdateService(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.Service
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(gateway.SuccessResp)

	idlManager.UpdateService(req)

	resp = &gateway.SuccessResp{
		Success: true,
		Message: "Update " + req.ServiceName + " success",
	}

	c.JSON(consts.StatusOK, resp)
}

// GetService .
// @router /get-service [POST]
func GetService(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.ServiceReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(gateway.Service)

	resp = idlManager.GetService(req.ServiceName)

	c.JSON(consts.StatusOK, resp)
}

// ListService .
// @router /list-service [POST]
func ListService(ctx context.Context, c *app.RequestContext) {
	var err error
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new([]*gateway.Service)
	services := idlManager.GetAllService()
	resp = &services

	c.JSON(consts.StatusOK, resp)
}
