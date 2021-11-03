// Code generated by go-swagger; DO NOT EDIT.

package namespace_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new namespace services API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for namespace services API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateNamespaceService(params *CreateNamespaceServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateNamespaceServiceOK, error)

	DeleteNamespaceRevision(params *DeleteNamespaceRevisionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteNamespaceRevisionOK, error)

	DeleteNamespaceService(params *DeleteNamespaceServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteNamespaceServiceOK, error)

	GetNamespaceService(params *GetNamespaceServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetNamespaceServiceOK, error)

	GetNamespaceServiceList(params *GetNamespaceServiceListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetNamespaceServiceListOK, error)

	ListNamespaceServiceRevisionPods(params *ListNamespaceServiceRevisionPodsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNamespaceServiceRevisionPodsOK, error)

	UpdateNamespaceService(params *UpdateNamespaceServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateNamespaceServiceOK, error)

	UpdateNamespaceServiceTraffic(params *UpdateNamespaceServiceTrafficParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateNamespaceServiceTrafficOK, error)

	WatchNamespaceServiceRevision(params *WatchNamespaceServiceRevisionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WatchNamespaceServiceRevisionOK, error)

	WatchNamespaceServiceRevisionList(params *WatchNamespaceServiceRevisionListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WatchNamespaceServiceRevisionListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateNamespaceService creates namespace service

  Creates namespace scoped knative service.
Service Names are unique on a scope level.
These services can be used as functions in workflows, more about this can be read here:
https://docs.direktiv.io/docs/walkthrough/using-functions.html

*/
func (a *Client) CreateNamespaceService(params *CreateNamespaceServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateNamespaceServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNamespaceServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createNamespaceService",
		Method:             "POST",
		PathPattern:        "/api/functions/namespaces/{namespace}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateNamespaceServiceReader{formats: a.formats},
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
	success, ok := result.(*CreateNamespaceServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createNamespaceService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteNamespaceRevision deletes namespace service revision

  Delete a namespace scoped knative service revision.
The target revision generation is the number suffix on a revision.
Example: A revisions named 'namespace-direktiv-fast-request-00003' would have the revisionGeneration '00003'.
Note: Revisions with traffic cannot be deleted.

*/
func (a *Client) DeleteNamespaceRevision(params *DeleteNamespaceRevisionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteNamespaceRevisionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNamespaceRevisionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteNamespaceRevision",
		Method:             "DELETE",
		PathPattern:        "/api/functions/namespaces/{namespace}/function/{serviceName}/revisions/{revisionGeneration}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNamespaceRevisionReader{formats: a.formats},
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
	success, ok := result.(*DeleteNamespaceRevisionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteNamespaceRevision: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteNamespaceService deletes namespace service

  Deletes namespace scoped knative service and all its revisions.

*/
func (a *Client) DeleteNamespaceService(params *DeleteNamespaceServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteNamespaceServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNamespaceServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteNamespaceService",
		Method:             "DELETE",
		PathPattern:        "/api/functions/namespaces/{namespace}/function/{serviceName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNamespaceServiceReader{formats: a.formats},
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
	success, ok := result.(*DeleteNamespaceServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteNamespaceService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNamespaceService gets namespace service details

  Get details of a namespace scoped knative service.

*/
func (a *Client) GetNamespaceService(params *GetNamespaceServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetNamespaceServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNamespaceServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getNamespaceService",
		Method:             "GET",
		PathPattern:        "/api/functions/namespaces/{namespace}/function/{serviceName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNamespaceServiceReader{formats: a.formats},
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
	success, ok := result.(*GetNamespaceServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNamespaceService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNamespaceServiceList gets namespace services list

  Gets a list of namespace knative services.

*/
func (a *Client) GetNamespaceServiceList(params *GetNamespaceServiceListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetNamespaceServiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNamespaceServiceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getNamespaceServiceList",
		Method:             "GET",
		PathPattern:        "/api/functions/namespaces/{namespace}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNamespaceServiceListReader{formats: a.formats},
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
	success, ok := result.(*GetNamespaceServiceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNamespaceServiceList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListNamespaceServiceRevisionPods gets namespace service revision pods list

  List a revisions pods of a namespace scoped knative service.
The target revision generation is the number suffix on a revision.
Example: A revisions named 'namespace-direktiv-fast-request-00003' would have the revisionGeneration '00003'.

*/
func (a *Client) ListNamespaceServiceRevisionPods(params *ListNamespaceServiceRevisionPodsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNamespaceServiceRevisionPodsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNamespaceServiceRevisionPodsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listNamespaceServiceRevisionPods",
		Method:             "GET",
		PathPattern:        "/api/functions/namespaces/{namespace}/function/{serviceName}/revisions/{revisionGeneration}/pods",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListNamespaceServiceRevisionPodsReader{formats: a.formats},
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
	success, ok := result.(*ListNamespaceServiceRevisionPodsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listNamespaceServiceRevisionPods: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateNamespaceService creates namespace service revision

  Creates a new namespace scoped knative service revision.
Revisions are created with a traffic percentage. This percentage controls
how much traffic will be directed to this revision. Traffic can be set to 100
to direct all traffic.

*/
func (a *Client) UpdateNamespaceService(params *UpdateNamespaceServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateNamespaceServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNamespaceServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateNamespaceService",
		Method:             "POST",
		PathPattern:        "/api/functions/namespaces/{namespace}/function/{serviceName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateNamespaceServiceReader{formats: a.formats},
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
	success, ok := result.(*UpdateNamespaceServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateNamespaceService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateNamespaceServiceTraffic updates namespace service traffic

  Update Namespace Service traffic directed to each revision,
traffic can only be configured between two revisions. All other revisions
will bet set to 0 traffic.

*/
func (a *Client) UpdateNamespaceServiceTraffic(params *UpdateNamespaceServiceTrafficParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateNamespaceServiceTrafficOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNamespaceServiceTrafficParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateNamespaceServiceTraffic",
		Method:             "PATCH",
		PathPattern:        "/api/functions/namespaces/{namespace}/function/{serviceName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateNamespaceServiceTrafficReader{formats: a.formats},
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
	success, ok := result.(*UpdateNamespaceServiceTrafficOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateNamespaceServiceTraffic: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  WatchNamespaceServiceRevision watches namespace service revision

  Watch a namespace scoped knative service revision.
The target revision generation is the number suffix on a revision.
Example: A revisions named 'namespace-direktiv-fast-request-00003' would have the revisionGeneration '00003'.
Note: This is a Server-Sent-Event endpoint, and will not work with the default swagger client.

*/
func (a *Client) WatchNamespaceServiceRevision(params *WatchNamespaceServiceRevisionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WatchNamespaceServiceRevisionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWatchNamespaceServiceRevisionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "watchNamespaceServiceRevision",
		Method:             "GET",
		PathPattern:        "/api/functions/namespaces/{namespace}/function/{serviceName}/revisions/{revisionGeneration}",
		ProducesMediaTypes: []string{"text/event-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &WatchNamespaceServiceRevisionReader{formats: a.formats},
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
	success, ok := result.(*WatchNamespaceServiceRevisionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for watchNamespaceServiceRevision: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  WatchNamespaceServiceRevisionList watches namespace service revision list

  Watch the revision list of a namespace scoped knative service.
Note: This is a Server-Sent-Event endpoint, and will not work with the default swagger client.

*/
func (a *Client) WatchNamespaceServiceRevisionList(params *WatchNamespaceServiceRevisionListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WatchNamespaceServiceRevisionListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWatchNamespaceServiceRevisionListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "watchNamespaceServiceRevisionList",
		Method:             "GET",
		PathPattern:        "/api/functions/namespaces/{namespace}/function/{serviceName}/revisions",
		ProducesMediaTypes: []string{"text/event-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &WatchNamespaceServiceRevisionListReader{formats: a.formats},
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
	success, ok := result.(*WatchNamespaceServiceRevisionListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for watchNamespaceServiceRevisionList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}