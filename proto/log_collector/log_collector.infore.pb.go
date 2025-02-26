// Code generated by protoc-gen-vkit. DO NOT EDIT.
// versions:
// - protoc-gen-vkit v1.0.0
// - protoc             v3.20.3
// source: log_collector/log_collector.proto

package log_collector

import (
	context "context"
	grpcx "github.com/vison888/go-vkit/grpcx"
	grpc "google.golang.org/grpc"
)

var _ = new(context.Context)
var _ = new(grpc.CallOption)
var _ = new(grpcx.Client)

type LogServiceClient struct {
	name string
	cc   grpcx.Client
}

func (c *LogServiceClient) One(ctx context.Context, in *CollectOneReq, opts ...grpc.CallOption) (*CollectOneResp, error) {
	out := new(CollectOneResp)
	err := c.cc.Invoke(ctx, c.name, "LogService.One", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *LogServiceClient) Batch(ctx context.Context, in *CollectBatchReq, opts ...grpc.CallOption) (*CollectBatchResp, error) {
	out := new(CollectBatchResp)
	err := c.cc.Invoke(ctx, c.name, "LogService.Batch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func NewLogServiceClient(name string, cc grpcx.Client) *LogServiceClient {
	return &LogServiceClient{name, cc}
}
