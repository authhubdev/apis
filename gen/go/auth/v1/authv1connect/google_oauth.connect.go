// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: auth/v1/google_oauth.proto

package authv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/authhubdev/apis/gen/go/auth/v1"
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
	// AuthGoogleOAuthServiceName is the fully-qualified name of the AuthGoogleOAuthService service.
	AuthGoogleOAuthServiceName = "auth.v1.AuthGoogleOAuthService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuthGoogleOAuthServiceGetAuthURLProcedure is the fully-qualified name of the
	// AuthGoogleOAuthService's GetAuthURL RPC.
	AuthGoogleOAuthServiceGetAuthURLProcedure = "/auth.v1.AuthGoogleOAuthService/GetAuthURL"
	// AuthGoogleOAuthServiceExchangeAuthCodeProcedure is the fully-qualified name of the
	// AuthGoogleOAuthService's ExchangeAuthCode RPC.
	AuthGoogleOAuthServiceExchangeAuthCodeProcedure = "/auth.v1.AuthGoogleOAuthService/ExchangeAuthCode"
)

// AuthGoogleOAuthServiceClient is a client for the auth.v1.AuthGoogleOAuthService service.
type AuthGoogleOAuthServiceClient interface {
	// GetAuthURL generates a Google OAuth authorization URL.
	// This URL is used to redirect users to Google's authentication page.
	GetAuthURL(context.Context, *connect.Request[v1.GetAuthURLRequest]) (*connect.Response[v1.GetAuthURLResponse], error)
	// ExchangeAuthCode exchanges a Google OAuth authorization code for authentication tokens.
	// It validates the authorization code and issues authentication tokens upon success.
	ExchangeAuthCode(context.Context, *connect.Request[v1.ExchangeAuthCodeRequest]) (*connect.Response[v1.ExchangeAuthCodeResponse], error)
}

// NewAuthGoogleOAuthServiceClient constructs a client for the auth.v1.AuthGoogleOAuthService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthGoogleOAuthServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AuthGoogleOAuthServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	authGoogleOAuthServiceMethods := v1.File_auth_v1_google_oauth_proto.Services().ByName("AuthGoogleOAuthService").Methods()
	return &authGoogleOAuthServiceClient{
		getAuthURL: connect.NewClient[v1.GetAuthURLRequest, v1.GetAuthURLResponse](
			httpClient,
			baseURL+AuthGoogleOAuthServiceGetAuthURLProcedure,
			connect.WithSchema(authGoogleOAuthServiceMethods.ByName("GetAuthURL")),
			connect.WithClientOptions(opts...),
		),
		exchangeAuthCode: connect.NewClient[v1.ExchangeAuthCodeRequest, v1.ExchangeAuthCodeResponse](
			httpClient,
			baseURL+AuthGoogleOAuthServiceExchangeAuthCodeProcedure,
			connect.WithSchema(authGoogleOAuthServiceMethods.ByName("ExchangeAuthCode")),
			connect.WithClientOptions(opts...),
		),
	}
}

// authGoogleOAuthServiceClient implements AuthGoogleOAuthServiceClient.
type authGoogleOAuthServiceClient struct {
	getAuthURL       *connect.Client[v1.GetAuthURLRequest, v1.GetAuthURLResponse]
	exchangeAuthCode *connect.Client[v1.ExchangeAuthCodeRequest, v1.ExchangeAuthCodeResponse]
}

// GetAuthURL calls auth.v1.AuthGoogleOAuthService.GetAuthURL.
func (c *authGoogleOAuthServiceClient) GetAuthURL(ctx context.Context, req *connect.Request[v1.GetAuthURLRequest]) (*connect.Response[v1.GetAuthURLResponse], error) {
	return c.getAuthURL.CallUnary(ctx, req)
}

// ExchangeAuthCode calls auth.v1.AuthGoogleOAuthService.ExchangeAuthCode.
func (c *authGoogleOAuthServiceClient) ExchangeAuthCode(ctx context.Context, req *connect.Request[v1.ExchangeAuthCodeRequest]) (*connect.Response[v1.ExchangeAuthCodeResponse], error) {
	return c.exchangeAuthCode.CallUnary(ctx, req)
}

// AuthGoogleOAuthServiceHandler is an implementation of the auth.v1.AuthGoogleOAuthService service.
type AuthGoogleOAuthServiceHandler interface {
	// GetAuthURL generates a Google OAuth authorization URL.
	// This URL is used to redirect users to Google's authentication page.
	GetAuthURL(context.Context, *connect.Request[v1.GetAuthURLRequest]) (*connect.Response[v1.GetAuthURLResponse], error)
	// ExchangeAuthCode exchanges a Google OAuth authorization code for authentication tokens.
	// It validates the authorization code and issues authentication tokens upon success.
	ExchangeAuthCode(context.Context, *connect.Request[v1.ExchangeAuthCodeRequest]) (*connect.Response[v1.ExchangeAuthCodeResponse], error)
}

// NewAuthGoogleOAuthServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthGoogleOAuthServiceHandler(svc AuthGoogleOAuthServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	authGoogleOAuthServiceMethods := v1.File_auth_v1_google_oauth_proto.Services().ByName("AuthGoogleOAuthService").Methods()
	authGoogleOAuthServiceGetAuthURLHandler := connect.NewUnaryHandler(
		AuthGoogleOAuthServiceGetAuthURLProcedure,
		svc.GetAuthURL,
		connect.WithSchema(authGoogleOAuthServiceMethods.ByName("GetAuthURL")),
		connect.WithHandlerOptions(opts...),
	)
	authGoogleOAuthServiceExchangeAuthCodeHandler := connect.NewUnaryHandler(
		AuthGoogleOAuthServiceExchangeAuthCodeProcedure,
		svc.ExchangeAuthCode,
		connect.WithSchema(authGoogleOAuthServiceMethods.ByName("ExchangeAuthCode")),
		connect.WithHandlerOptions(opts...),
	)
	return "/auth.v1.AuthGoogleOAuthService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AuthGoogleOAuthServiceGetAuthURLProcedure:
			authGoogleOAuthServiceGetAuthURLHandler.ServeHTTP(w, r)
		case AuthGoogleOAuthServiceExchangeAuthCodeProcedure:
			authGoogleOAuthServiceExchangeAuthCodeHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuthGoogleOAuthServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthGoogleOAuthServiceHandler struct{}

func (UnimplementedAuthGoogleOAuthServiceHandler) GetAuthURL(context.Context, *connect.Request[v1.GetAuthURLRequest]) (*connect.Response[v1.GetAuthURLResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth.v1.AuthGoogleOAuthService.GetAuthURL is not implemented"))
}

func (UnimplementedAuthGoogleOAuthServiceHandler) ExchangeAuthCode(context.Context, *connect.Request[v1.ExchangeAuthCodeRequest]) (*connect.Response[v1.ExchangeAuthCodeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth.v1.AuthGoogleOAuthService.ExchangeAuthCode is not implemented"))
}
