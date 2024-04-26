// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v3.19.4
// source: server/dashboard/dashboard.proto

package dashboard

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationDashboardCreateDashboard = "/api.server.dashboard.Dashboard/CreateDashboard"
const OperationDashboardDeleteDashboard = "/api.server.dashboard.Dashboard/DeleteDashboard"
const OperationDashboardGetDashboard = "/api.server.dashboard.Dashboard/GetDashboard"
const OperationDashboardListDashboard = "/api.server.dashboard.Dashboard/ListDashboard"
const OperationDashboardListDashboardSelect = "/api.server.dashboard.Dashboard/ListDashboardSelect"
const OperationDashboardUpdateDashboard = "/api.server.dashboard.Dashboard/UpdateDashboard"

type DashboardHTTPServer interface {
	// CreateDashboard 创建新的大盘
	CreateDashboard(context.Context, *CreateDashboardRequest) (*CreateDashboardReply, error)
	// DeleteDashboard 删除大盘
	DeleteDashboard(context.Context, *DeleteDashboardRequest) (*DeleteDashboardReply, error)
	// GetDashboard 大盘详情
	GetDashboard(context.Context, *GetDashboardRequest) (*GetDashboardReply, error)
	// ListDashboard 大盘列表
	ListDashboard(context.Context, *ListDashboardRequest) (*ListDashboardReply, error)
	// ListDashboardSelect 大盘下拉列表
	ListDashboardSelect(context.Context, *ListDashboardSelectRequest) (*ListDashboardSelectReply, error)
	// UpdateDashboard 更新大盘
	UpdateDashboard(context.Context, *UpdateDashboardRequest) (*UpdateDashboardReply, error)
}

func RegisterDashboardHTTPServer(s *http.Server, srv DashboardHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/dashboard/create", _Dashboard_CreateDashboard0_HTTP_Handler(srv))
	r.POST("/api/v1/dashboard/update", _Dashboard_UpdateDashboard0_HTTP_Handler(srv))
	r.POST("/api/v1/dashboard/delete", _Dashboard_DeleteDashboard0_HTTP_Handler(srv))
	r.POST("/api/v1/dashboard/detail", _Dashboard_GetDashboard0_HTTP_Handler(srv))
	r.POST("/api/v1/dashboard/list", _Dashboard_ListDashboard0_HTTP_Handler(srv))
	r.POST("/api/v1/dashboard/select", _Dashboard_ListDashboardSelect0_HTTP_Handler(srv))
}

func _Dashboard_CreateDashboard0_HTTP_Handler(srv DashboardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateDashboardRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDashboardCreateDashboard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateDashboard(ctx, req.(*CreateDashboardRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateDashboardReply)
		return ctx.Result(200, reply)
	}
}

func _Dashboard_UpdateDashboard0_HTTP_Handler(srv DashboardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateDashboardRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDashboardUpdateDashboard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateDashboard(ctx, req.(*UpdateDashboardRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateDashboardReply)
		return ctx.Result(200, reply)
	}
}

func _Dashboard_DeleteDashboard0_HTTP_Handler(srv DashboardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteDashboardRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDashboardDeleteDashboard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteDashboard(ctx, req.(*DeleteDashboardRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteDashboardReply)
		return ctx.Result(200, reply)
	}
}

func _Dashboard_GetDashboard0_HTTP_Handler(srv DashboardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDashboardRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDashboardGetDashboard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDashboard(ctx, req.(*GetDashboardRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetDashboardReply)
		return ctx.Result(200, reply)
	}
}

func _Dashboard_ListDashboard0_HTTP_Handler(srv DashboardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListDashboardRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDashboardListDashboard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDashboard(ctx, req.(*ListDashboardRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListDashboardReply)
		return ctx.Result(200, reply)
	}
}

func _Dashboard_ListDashboardSelect0_HTTP_Handler(srv DashboardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListDashboardSelectRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDashboardListDashboardSelect)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDashboardSelect(ctx, req.(*ListDashboardSelectRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListDashboardSelectReply)
		return ctx.Result(200, reply)
	}
}

type DashboardHTTPClient interface {
	CreateDashboard(ctx context.Context, req *CreateDashboardRequest, opts ...http.CallOption) (rsp *CreateDashboardReply, err error)
	DeleteDashboard(ctx context.Context, req *DeleteDashboardRequest, opts ...http.CallOption) (rsp *DeleteDashboardReply, err error)
	GetDashboard(ctx context.Context, req *GetDashboardRequest, opts ...http.CallOption) (rsp *GetDashboardReply, err error)
	ListDashboard(ctx context.Context, req *ListDashboardRequest, opts ...http.CallOption) (rsp *ListDashboardReply, err error)
	ListDashboardSelect(ctx context.Context, req *ListDashboardSelectRequest, opts ...http.CallOption) (rsp *ListDashboardSelectReply, err error)
	UpdateDashboard(ctx context.Context, req *UpdateDashboardRequest, opts ...http.CallOption) (rsp *UpdateDashboardReply, err error)
}

type DashboardHTTPClientImpl struct {
	cc *http.Client
}

func NewDashboardHTTPClient(client *http.Client) DashboardHTTPClient {
	return &DashboardHTTPClientImpl{client}
}

func (c *DashboardHTTPClientImpl) CreateDashboard(ctx context.Context, in *CreateDashboardRequest, opts ...http.CallOption) (*CreateDashboardReply, error) {
	var out CreateDashboardReply
	pattern := "/api/v1/dashboard/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDashboardCreateDashboard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DashboardHTTPClientImpl) DeleteDashboard(ctx context.Context, in *DeleteDashboardRequest, opts ...http.CallOption) (*DeleteDashboardReply, error) {
	var out DeleteDashboardReply
	pattern := "/api/v1/dashboard/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDashboardDeleteDashboard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DashboardHTTPClientImpl) GetDashboard(ctx context.Context, in *GetDashboardRequest, opts ...http.CallOption) (*GetDashboardReply, error) {
	var out GetDashboardReply
	pattern := "/api/v1/dashboard/detail"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDashboardGetDashboard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DashboardHTTPClientImpl) ListDashboard(ctx context.Context, in *ListDashboardRequest, opts ...http.CallOption) (*ListDashboardReply, error) {
	var out ListDashboardReply
	pattern := "/api/v1/dashboard/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDashboardListDashboard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DashboardHTTPClientImpl) ListDashboardSelect(ctx context.Context, in *ListDashboardSelectRequest, opts ...http.CallOption) (*ListDashboardSelectReply, error) {
	var out ListDashboardSelectReply
	pattern := "/api/v1/dashboard/select"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDashboardListDashboardSelect))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DashboardHTTPClientImpl) UpdateDashboard(ctx context.Context, in *UpdateDashboardRequest, opts ...http.CallOption) (*UpdateDashboardReply, error) {
	var out UpdateDashboardReply
	pattern := "/api/v1/dashboard/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDashboardUpdateDashboard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
