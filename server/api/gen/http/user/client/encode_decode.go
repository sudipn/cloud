// Code generated by goa v3.1.2, DO NOT EDIT.
//
// user HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	user "github.com/fieldkit/cloud/server/api/gen/user"
	userviews "github.com/fieldkit/cloud/server/api/gen/user/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildRolesRequest instantiates a HTTP request object with method and path
// set to call the "user" service "roles" endpoint
func (c *Client) BuildRolesRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RolesUserPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "roles", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeRolesRequest returns an encoder for requests sent to the user roles
// server.
func EncodeRolesRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.RolesPayload)
		if !ok {
			return goahttp.ErrInvalidType("user", "roles", "*user.RolesPayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeRolesResponse returns a decoder for responses returned by the user
// roles endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeRolesResponse may return the following errors:
//	- "bad-request" (type user.BadRequest): http.StatusBadRequest
//	- "forbidden" (type user.Forbidden): http.StatusForbidden
//	- "not-found" (type user.NotFound): http.StatusNotFound
//	- "unauthorized" (type user.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeRolesResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body RolesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "roles", err)
			}
			p := NewRolesAvailableRolesOK(&body)
			view := "default"
			vres := &userviews.AvailableRoles{Projected: p, View: view}
			if err = userviews.ValidateAvailableRoles(vres); err != nil {
				return nil, goahttp.ErrValidationError("user", "roles", err)
			}
			res := user.NewAvailableRoles(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body RolesBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "roles", err)
			}
			return nil, NewRolesBadRequest(body)
		case http.StatusForbidden:
			var (
				body RolesForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "roles", err)
			}
			return nil, NewRolesForbidden(body)
		case http.StatusNotFound:
			var (
				body RolesNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "roles", err)
			}
			return nil, NewRolesNotFound(body)
		case http.StatusUnauthorized:
			var (
				body RolesUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "roles", err)
			}
			return nil, NewRolesUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "roles", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteRequest instantiates a HTTP request object with method and path
// set to call the "user" service "delete" endpoint
func (c *Client) BuildDeleteRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		userID int32
	)
	{
		p, ok := v.(*user.DeletePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("user", "delete", "*user.DeletePayload", v)
		}
		userID = p.UserID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteUserPath(userID)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteRequest returns an encoder for requests sent to the user delete
// server.
func EncodeDeleteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.DeletePayload)
		if !ok {
			return goahttp.ErrInvalidType("user", "delete", "*user.DeletePayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeDeleteResponse returns a decoder for responses returned by the user
// delete endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeDeleteResponse may return the following errors:
//	- "bad-request" (type user.BadRequest): http.StatusBadRequest
//	- "forbidden" (type user.Forbidden): http.StatusForbidden
//	- "not-found" (type user.NotFound): http.StatusNotFound
//	- "unauthorized" (type user.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeDeleteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		case http.StatusBadRequest:
			var (
				body DeleteBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "delete", err)
			}
			return nil, NewDeleteBadRequest(body)
		case http.StatusForbidden:
			var (
				body DeleteForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "delete", err)
			}
			return nil, NewDeleteForbidden(body)
		case http.StatusNotFound:
			var (
				body DeleteNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "delete", err)
			}
			return nil, NewDeleteNotFound(body)
		case http.StatusUnauthorized:
			var (
				body DeleteUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "delete", err)
			}
			return nil, NewDeleteUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "delete", resp.StatusCode, string(body))
		}
	}
}

// unmarshalAvailableRoleResponseBodyToUserviewsAvailableRoleView builds a
// value of type *userviews.AvailableRoleView from a value of type
// *AvailableRoleResponseBody.
func unmarshalAvailableRoleResponseBodyToUserviewsAvailableRoleView(v *AvailableRoleResponseBody) *userviews.AvailableRoleView {
	res := &userviews.AvailableRoleView{
		ID:   v.ID,
		Name: v.Name,
	}

	return res
}
