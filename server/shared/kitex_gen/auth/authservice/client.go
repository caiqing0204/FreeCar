// Code generated by Kitex v0.4.4. DO NOT EDIT.

package authservice

import (
	"context"
	auth "github.com/CyanAsterisk/FreeCar/server/shared/kitex_gen/auth"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Login(ctx context.Context, req *auth.LoginRequest, callOptions ...callopt.Option) (r *auth.LoginResponse, err error)
	UploadAvatar(ctx context.Context, req *auth.UploadAvatarRequset, callOptions ...callopt.Option) (r *auth.UploadAvatarResponse, err error)
	UpdateUser(ctx context.Context, req *auth.UpdateUserRequest, callOptions ...callopt.Option) (r *auth.UpdateUserResponse, err error)
	GetUser(ctx context.Context, req *auth.GetUserRequest, callOptions ...callopt.Option) (r *auth.UserInfo, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kAuthServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kAuthServiceClient struct {
	*kClient
}

func (p *kAuthServiceClient) Login(ctx context.Context, req *auth.LoginRequest, callOptions ...callopt.Option) (r *auth.LoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, req)
}

func (p *kAuthServiceClient) UploadAvatar(ctx context.Context, req *auth.UploadAvatarRequset, callOptions ...callopt.Option) (r *auth.UploadAvatarResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UploadAvatar(ctx, req)
}

func (p *kAuthServiceClient) UpdateUser(ctx context.Context, req *auth.UpdateUserRequest, callOptions ...callopt.Option) (r *auth.UpdateUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUser(ctx, req)
}

func (p *kAuthServiceClient) GetUser(ctx context.Context, req *auth.GetUserRequest, callOptions ...callopt.Option) (r *auth.UserInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUser(ctx, req)
}
