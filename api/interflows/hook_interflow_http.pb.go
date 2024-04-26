// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v3.19.4
// source: interflows/hook_interflow.proto

package interflows

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

const OperationHookInterflowReceive = "/api.interflows.HookInterflow/Receive"

type HookInterflowHTTPServer interface {
	Receive(context.Context, *ReceiveRequest) (*ReceiveResponse, error)
}

func RegisterHookInterflowHTTPServer(s *http.Server, srv HookInterflowHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/interflows/receive", _HookInterflow_Receive0_HTTP_Handler(srv))
}

func _HookInterflow_Receive0_HTTP_Handler(srv HookInterflowHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReceiveRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationHookInterflowReceive)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Receive(ctx, req.(*ReceiveRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ReceiveResponse)
		return ctx.Result(200, reply)
	}
}

type HookInterflowHTTPClient interface {
	Receive(ctx context.Context, req *ReceiveRequest, opts ...http.CallOption) (rsp *ReceiveResponse, err error)
}

type HookInterflowHTTPClientImpl struct {
	cc *http.Client
}

func NewHookInterflowHTTPClient(client *http.Client) HookInterflowHTTPClient {
	return &HookInterflowHTTPClientImpl{client}
}

func (c *HookInterflowHTTPClientImpl) Receive(ctx context.Context, in *ReceiveRequest, opts ...http.CallOption) (*ReceiveResponse, error) {
	var out ReceiveResponse
	pattern := "/api/v1/interflows/receive"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationHookInterflowReceive))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
