// Code generated by go-swagger; DO NOT EDIT.

package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new identity API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for identity API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAccount(params *CreateAccountParams, opts ...ClientOption) (*CreateAccountCreated, error)

	Disable(params *DisableParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DisableOK, error)

	Enable(params *EnableParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EnableCreated, error)

	Login(params *LoginParams, opts ...ClientOption) (*LoginOK, error)

	Verify(params *VerifyParams, opts ...ClientOption) (*VerifyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateAccount create account API
*/
func (a *Client) CreateAccount(params *CreateAccountParams, opts ...ClientOption) (*CreateAccountCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createAccount",
		Method:             "POST",
		PathPattern:        "/account",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateAccountCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAccount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Disable disable API
*/
func (a *Client) Disable(params *DisableParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DisableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisableParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "disable",
		Method:             "POST",
		PathPattern:        "/disable",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DisableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for disable: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Enable enable API
*/
func (a *Client) Enable(params *EnableParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EnableCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnableParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "enable",
		Method:             "POST",
		PathPattern:        "/enable",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EnableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EnableCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enable: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Login login API
*/
func (a *Client) Login(params *LoginParams, opts ...ClientOption) (*LoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoginParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "login",
		Method:             "POST",
		PathPattern:        "/login",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LoginReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LoginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for login: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Verify verify API
*/
func (a *Client) Verify(params *VerifyParams, opts ...ClientOption) (*VerifyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVerifyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "verify",
		Method:             "GET",
		PathPattern:        "/verify",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VerifyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VerifyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for verify: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
