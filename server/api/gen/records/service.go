// Code generated by goa v3.1.2, DO NOT EDIT.
//
// records service
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package records

import (
	"context"

	"goa.design/goa/v3/security"
)

// Service is the records service interface.
type Service interface {
	// Data implements data.
	Data(context.Context, *DataPayload) (res *DataResult, err error)
	// Meta implements meta.
	Meta(context.Context, *MetaPayload) (res *MetaResult, err error)
	// Resolved implements resolved.
	Resolved(context.Context, *ResolvedPayload) (res *ResolvedResult, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "records"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"data", "meta", "resolved"}

// DataPayload is the payload type of the records service data method.
type DataPayload struct {
	Auth     *string
	RecordID *int64
}

// DataResult is the result type of the records service data method.
type DataResult struct {
	Object interface{}
}

// MetaPayload is the payload type of the records service meta method.
type MetaPayload struct {
	Auth     *string
	RecordID *int64
}

// MetaResult is the result type of the records service meta method.
type MetaResult struct {
	Object interface{}
}

// ResolvedPayload is the payload type of the records service resolved method.
type ResolvedPayload struct {
	Auth     *string
	RecordID *int64
}

// ResolvedResult is the result type of the records service resolved method.
type ResolvedResult struct {
	Object interface{}
}

// unauthorized
type Unauthorized string

// forbidden
type Forbidden string

// not-found
type NotFound string

// bad-request
type BadRequest string

// Error returns an error description.
func (e Unauthorized) Error() string {
	return "unauthorized"
}

// ErrorName returns "unauthorized".
func (e Unauthorized) ErrorName() string {
	return "unauthorized"
}

// Error returns an error description.
func (e Forbidden) Error() string {
	return "forbidden"
}

// ErrorName returns "forbidden".
func (e Forbidden) ErrorName() string {
	return "forbidden"
}

// Error returns an error description.
func (e NotFound) Error() string {
	return "not-found"
}

// ErrorName returns "not-found".
func (e NotFound) ErrorName() string {
	return "not-found"
}

// Error returns an error description.
func (e BadRequest) Error() string {
	return "bad-request"
}

// ErrorName returns "bad-request".
func (e BadRequest) ErrorName() string {
	return "bad-request"
}
