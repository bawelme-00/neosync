// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: mgmt/v1alpha1/anonymization.proto

package mgmtv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// AnonymizationServiceName is the fully-qualified name of the AnonymizationService service.
	AnonymizationServiceName = "mgmt.v1alpha1.AnonymizationService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AnonymizationServiceAnonymizeManyProcedure is the fully-qualified name of the
	// AnonymizationService's AnonymizeMany RPC.
	AnonymizationServiceAnonymizeManyProcedure = "/mgmt.v1alpha1.AnonymizationService/AnonymizeMany"
	// AnonymizationServiceAnonymizeSingleProcedure is the fully-qualified name of the
	// AnonymizationService's AnonymizeSingle RPC.
	AnonymizationServiceAnonymizeSingleProcedure = "/mgmt.v1alpha1.AnonymizationService/AnonymizeSingle"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	anonymizationServiceServiceDescriptor               = v1alpha1.File_mgmt_v1alpha1_anonymization_proto.Services().ByName("AnonymizationService")
	anonymizationServiceAnonymizeManyMethodDescriptor   = anonymizationServiceServiceDescriptor.Methods().ByName("AnonymizeMany")
	anonymizationServiceAnonymizeSingleMethodDescriptor = anonymizationServiceServiceDescriptor.Methods().ByName("AnonymizeSingle")
)

// AnonymizationServiceClient is a client for the mgmt.v1alpha1.AnonymizationService service.
type AnonymizationServiceClient interface {
	// Anonymizes many JSON strings by applying specified transformation mappings.
	AnonymizeMany(context.Context, *connect.Request[v1alpha1.AnonymizeManyRequest]) (*connect.Response[v1alpha1.AnonymizeManyResponse], error)
	// Anonymizes a single JSON strings by applying specified transformation mappings.
	AnonymizeSingle(context.Context, *connect.Request[v1alpha1.AnonymizeSingleRequest]) (*connect.Response[v1alpha1.AnonymizeSingleResponse], error)
}

// NewAnonymizationServiceClient constructs a client for the mgmt.v1alpha1.AnonymizationService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAnonymizationServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AnonymizationServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &anonymizationServiceClient{
		anonymizeMany: connect.NewClient[v1alpha1.AnonymizeManyRequest, v1alpha1.AnonymizeManyResponse](
			httpClient,
			baseURL+AnonymizationServiceAnonymizeManyProcedure,
			connect.WithSchema(anonymizationServiceAnonymizeManyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		anonymizeSingle: connect.NewClient[v1alpha1.AnonymizeSingleRequest, v1alpha1.AnonymizeSingleResponse](
			httpClient,
			baseURL+AnonymizationServiceAnonymizeSingleProcedure,
			connect.WithSchema(anonymizationServiceAnonymizeSingleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// anonymizationServiceClient implements AnonymizationServiceClient.
type anonymizationServiceClient struct {
	anonymizeMany   *connect.Client[v1alpha1.AnonymizeManyRequest, v1alpha1.AnonymizeManyResponse]
	anonymizeSingle *connect.Client[v1alpha1.AnonymizeSingleRequest, v1alpha1.AnonymizeSingleResponse]
}

// AnonymizeMany calls mgmt.v1alpha1.AnonymizationService.AnonymizeMany.
func (c *anonymizationServiceClient) AnonymizeMany(ctx context.Context, req *connect.Request[v1alpha1.AnonymizeManyRequest]) (*connect.Response[v1alpha1.AnonymizeManyResponse], error) {
	return c.anonymizeMany.CallUnary(ctx, req)
}

// AnonymizeSingle calls mgmt.v1alpha1.AnonymizationService.AnonymizeSingle.
func (c *anonymizationServiceClient) AnonymizeSingle(ctx context.Context, req *connect.Request[v1alpha1.AnonymizeSingleRequest]) (*connect.Response[v1alpha1.AnonymizeSingleResponse], error) {
	return c.anonymizeSingle.CallUnary(ctx, req)
}

// AnonymizationServiceHandler is an implementation of the mgmt.v1alpha1.AnonymizationService
// service.
type AnonymizationServiceHandler interface {
	// Anonymizes many JSON strings by applying specified transformation mappings.
	AnonymizeMany(context.Context, *connect.Request[v1alpha1.AnonymizeManyRequest]) (*connect.Response[v1alpha1.AnonymizeManyResponse], error)
	// Anonymizes a single JSON strings by applying specified transformation mappings.
	AnonymizeSingle(context.Context, *connect.Request[v1alpha1.AnonymizeSingleRequest]) (*connect.Response[v1alpha1.AnonymizeSingleResponse], error)
}

// NewAnonymizationServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAnonymizationServiceHandler(svc AnonymizationServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	anonymizationServiceAnonymizeManyHandler := connect.NewUnaryHandler(
		AnonymizationServiceAnonymizeManyProcedure,
		svc.AnonymizeMany,
		connect.WithSchema(anonymizationServiceAnonymizeManyMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	anonymizationServiceAnonymizeSingleHandler := connect.NewUnaryHandler(
		AnonymizationServiceAnonymizeSingleProcedure,
		svc.AnonymizeSingle,
		connect.WithSchema(anonymizationServiceAnonymizeSingleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/mgmt.v1alpha1.AnonymizationService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AnonymizationServiceAnonymizeManyProcedure:
			anonymizationServiceAnonymizeManyHandler.ServeHTTP(w, r)
		case AnonymizationServiceAnonymizeSingleProcedure:
			anonymizationServiceAnonymizeSingleHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAnonymizationServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAnonymizationServiceHandler struct{}

func (UnimplementedAnonymizationServiceHandler) AnonymizeMany(context.Context, *connect.Request[v1alpha1.AnonymizeManyRequest]) (*connect.Response[v1alpha1.AnonymizeManyResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.AnonymizationService.AnonymizeMany is not implemented"))
}

func (UnimplementedAnonymizationServiceHandler) AnonymizeSingle(context.Context, *connect.Request[v1alpha1.AnonymizeSingleRequest]) (*connect.Response[v1alpha1.AnonymizeSingleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.AnonymizationService.AnonymizeSingle is not implemented"))
}
