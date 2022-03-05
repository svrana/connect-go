// Copyright 2021-2022 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: connect/ping/v1test/ping.proto

package pingv1test

import (
	context "context"
	errors "errors"
	connect "github.com/bufbuild/connect"
	v1test "github.com/bufbuild/connect/internal/gen/proto/go/connect/ping/v1test"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant isn't defined, this code was generated
// with a version of connect newer than the one compiled into your binary. You can fix the problem
// by either regenerating this code with an older version of connect or updating the connect version
// compiled into your binary.
const _ = connect.IsAtLeastVersion0_0_1

// PingServiceClient is a client for the connect.ping.v1test.PingService service.
type PingServiceClient interface {
	// Ping sends a ping to the server to determine if it's reachable.
	Ping(context.Context, *connect.Envelope[v1test.PingRequest]) (*connect.Envelope[v1test.PingResponse], error)
	// Fail always fails.
	Fail(context.Context, *connect.Envelope[v1test.FailRequest]) (*connect.Envelope[v1test.FailResponse], error)
	// Sum calculates the sum of the numbers sent on the stream.
	Sum(context.Context) *connect.ClientStreamForClient[v1test.SumRequest, v1test.SumResponse]
	// CountUp returns a stream of the numbers up to the given request.
	CountUp(context.Context, *connect.Envelope[v1test.CountUpRequest]) (*connect.ServerStreamForClient[v1test.CountUpResponse], error)
	// CumSum determines the cumulative sum of all the numbers sent on the stream.
	CumSum(context.Context) *connect.BidiStreamForClient[v1test.CumSumRequest, v1test.CumSumResponse]
}

// NewPingServiceClient constructs a client for the connect.ping.v1test.PingService service. By
// default, it uses the HTTP/2 gRPC protocol and the binary protobuf Codec. It asks for gzipped
// responses and sends uncompressed requests.
//
// The URL supplied here should be the base URL for the gRPC server (e.g., https://api.acme.com or
// https://acme.com/grpc).
func NewPingServiceClient(baseURL string, doer connect.Doer, opts ...connect.ClientOption) (PingServiceClient, error) {
	baseURL = strings.TrimRight(baseURL, "/")
	pingClient, pingErr := connect.NewClient[v1test.PingRequest, v1test.PingResponse](
		baseURL,
		"connect.ping.v1test.PingService/Ping",
		doer,
		opts...,
	)
	if pingErr != nil {
		return nil, pingErr
	}
	failClient, failErr := connect.NewClient[v1test.FailRequest, v1test.FailResponse](
		baseURL,
		"connect.ping.v1test.PingService/Fail",
		doer,
		opts...,
	)
	if failErr != nil {
		return nil, failErr
	}
	sumClient, sumErr := connect.NewClient[v1test.SumRequest, v1test.SumResponse](
		baseURL,
		"connect.ping.v1test.PingService/Sum",
		doer,
		opts...,
	)
	if sumErr != nil {
		return nil, sumErr
	}
	countUpClient, countUpErr := connect.NewClient[v1test.CountUpRequest, v1test.CountUpResponse](
		baseURL,
		"connect.ping.v1test.PingService/CountUp",
		doer,
		opts...,
	)
	if countUpErr != nil {
		return nil, countUpErr
	}
	cumSumClient, cumSumErr := connect.NewClient[v1test.CumSumRequest, v1test.CumSumResponse](
		baseURL,
		"connect.ping.v1test.PingService/CumSum",
		doer,
		opts...,
	)
	if cumSumErr != nil {
		return nil, cumSumErr
	}
	return &pingServiceClient{
		ping:    pingClient,
		fail:    failClient,
		sum:     sumClient,
		countUp: countUpClient,
		cumSum:  cumSumClient,
	}, nil
}

// pingServiceClient implements PingServiceClient.
type pingServiceClient struct {
	ping    *connect.Client[v1test.PingRequest, v1test.PingResponse]
	fail    *connect.Client[v1test.FailRequest, v1test.FailResponse]
	sum     *connect.Client[v1test.SumRequest, v1test.SumResponse]
	countUp *connect.Client[v1test.CountUpRequest, v1test.CountUpResponse]
	cumSum  *connect.Client[v1test.CumSumRequest, v1test.CumSumResponse]
}

// Ping calls connect.ping.v1test.PingService.Ping.
func (c *pingServiceClient) Ping(ctx context.Context, req *connect.Envelope[v1test.PingRequest]) (*connect.Envelope[v1test.PingResponse], error) {
	return c.ping.CallUnary(ctx, req)
}

// Fail calls connect.ping.v1test.PingService.Fail.
func (c *pingServiceClient) Fail(ctx context.Context, req *connect.Envelope[v1test.FailRequest]) (*connect.Envelope[v1test.FailResponse], error) {
	return c.fail.CallUnary(ctx, req)
}

// Sum calls connect.ping.v1test.PingService.Sum.
func (c *pingServiceClient) Sum(ctx context.Context) *connect.ClientStreamForClient[v1test.SumRequest, v1test.SumResponse] {
	return c.sum.CallClientStream(ctx)
}

// CountUp calls connect.ping.v1test.PingService.CountUp.
func (c *pingServiceClient) CountUp(ctx context.Context, req *connect.Envelope[v1test.CountUpRequest]) (*connect.ServerStreamForClient[v1test.CountUpResponse], error) {
	return c.countUp.CallServerStream(ctx, req)
}

// CumSum calls connect.ping.v1test.PingService.CumSum.
func (c *pingServiceClient) CumSum(ctx context.Context) *connect.BidiStreamForClient[v1test.CumSumRequest, v1test.CumSumResponse] {
	return c.cumSum.CallBidiStream(ctx)
}

// PingServiceHandler is an implementation of the connect.ping.v1test.PingService service.
type PingServiceHandler interface {
	// Ping sends a ping to the server to determine if it's reachable.
	Ping(context.Context, *connect.Envelope[v1test.PingRequest]) (*connect.Envelope[v1test.PingResponse], error)
	// Fail always fails.
	Fail(context.Context, *connect.Envelope[v1test.FailRequest]) (*connect.Envelope[v1test.FailResponse], error)
	// Sum calculates the sum of the numbers sent on the stream.
	Sum(context.Context, *connect.ClientStream[v1test.SumRequest, v1test.SumResponse]) error
	// CountUp returns a stream of the numbers up to the given request.
	CountUp(context.Context, *connect.Envelope[v1test.CountUpRequest], *connect.ServerStream[v1test.CountUpResponse]) error
	// CumSum determines the cumulative sum of all the numbers sent on the stream.
	CumSum(context.Context, *connect.BidiStream[v1test.CumSumRequest, v1test.CumSumResponse]) error
}

// NewPingServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the gRPC and gRPC-Web protocols with the binary protobuf and JSON
// codecs.
func NewPingServiceHandler(svc PingServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/connect.ping.v1test.PingService/Ping", connect.NewUnaryHandler(
		"connect.ping.v1test.PingService/Ping", // procedure name
		"connect.ping.v1test.PingService",      // reflection name
		svc.Ping,
		opts...,
	))
	mux.Handle("/connect.ping.v1test.PingService/Fail", connect.NewUnaryHandler(
		"connect.ping.v1test.PingService/Fail", // procedure name
		"connect.ping.v1test.PingService",      // reflection name
		svc.Fail,
		opts...,
	))
	mux.Handle("/connect.ping.v1test.PingService/Sum", connect.NewClientStreamHandler(
		"connect.ping.v1test.PingService/Sum", // procedure name
		"connect.ping.v1test.PingService",     // reflection name
		svc.Sum,
		opts...,
	))
	mux.Handle("/connect.ping.v1test.PingService/CountUp", connect.NewServerStreamHandler(
		"connect.ping.v1test.PingService/CountUp", // procedure name
		"connect.ping.v1test.PingService",         // reflection name
		svc.CountUp,
		opts...,
	))
	mux.Handle("/connect.ping.v1test.PingService/CumSum", connect.NewBidiStreamHandler(
		"connect.ping.v1test.PingService/CumSum", // procedure name
		"connect.ping.v1test.PingService",        // reflection name
		svc.CumSum,
		opts...,
	))
	return "/connect.ping.v1test.PingService/", mux
}

// UnimplementedPingServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPingServiceHandler struct{}

func (UnimplementedPingServiceHandler) Ping(context.Context, *connect.Envelope[v1test.PingRequest]) (*connect.Envelope[v1test.PingResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("connect.ping.v1test.PingService.Ping isn't implemented"))
}

func (UnimplementedPingServiceHandler) Fail(context.Context, *connect.Envelope[v1test.FailRequest]) (*connect.Envelope[v1test.FailResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("connect.ping.v1test.PingService.Fail isn't implemented"))
}

func (UnimplementedPingServiceHandler) Sum(context.Context, *connect.ClientStream[v1test.SumRequest, v1test.SumResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("connect.ping.v1test.PingService.Sum isn't implemented"))
}

func (UnimplementedPingServiceHandler) CountUp(context.Context, *connect.Envelope[v1test.CountUpRequest], *connect.ServerStream[v1test.CountUpResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("connect.ping.v1test.PingService.CountUp isn't implemented"))
}

func (UnimplementedPingServiceHandler) CumSum(context.Context, *connect.BidiStream[v1test.CumSumRequest, v1test.CumSumResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("connect.ping.v1test.PingService.CumSum isn't implemented"))
}
