// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "fieldkit": Application Contexts
//
// Command:
// $ main

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"strconv"
)

// AddFirmwareContext provides the Firmware add action context.
type AddFirmwareContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *AddFirmwarePayload
}

// NewAddFirmwareContext parses the incoming request URL and body, performs validations and creates the
// context used by the Firmware controller add action.
func NewAddFirmwareContext(ctx context.Context, r *http.Request, service *goa.Service) (*AddFirmwareContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := AddFirmwareContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *AddFirmwareContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *AddFirmwareContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// DeleteFirmwareContext provides the Firmware delete action context.
type DeleteFirmwareContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	FirmwareID int
}

// NewDeleteFirmwareContext parses the incoming request URL and body, performs validations and creates the
// context used by the Firmware controller delete action.
func NewDeleteFirmwareContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteFirmwareContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteFirmwareContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramFirmwareID := req.Params["firmwareId"]
	if len(paramFirmwareID) > 0 {
		rawFirmwareID := paramFirmwareID[0]
		if firmwareID, err2 := strconv.Atoi(rawFirmwareID); err2 == nil {
			rctx.FirmwareID = firmwareID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("firmwareId", rawFirmwareID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DeleteFirmwareContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteFirmwareContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// DownloadFirmwareContext provides the Firmware download action context.
type DownloadFirmwareContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	FirmwareID int
}

// NewDownloadFirmwareContext parses the incoming request URL and body, performs validations and creates the
// context used by the Firmware controller download action.
func NewDownloadFirmwareContext(ctx context.Context, r *http.Request, service *goa.Service) (*DownloadFirmwareContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DownloadFirmwareContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramFirmwareID := req.Params["firmwareId"]
	if len(paramFirmwareID) > 0 {
		rawFirmwareID := paramFirmwareID[0]
		if firmwareID, err2 := strconv.Atoi(rawFirmwareID); err2 == nil {
			rctx.FirmwareID = firmwareID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("firmwareId", rawFirmwareID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DownloadFirmwareContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DownloadFirmwareContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListFirmwareContext provides the Firmware list action context.
type ListFirmwareContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Module   *string
	Page     *int
	PageSize *int
	Profile  *string
}

// NewListFirmwareContext parses the incoming request URL and body, performs validations and creates the
// context used by the Firmware controller list action.
func NewListFirmwareContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListFirmwareContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListFirmwareContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramModule := req.Params["module"]
	if len(paramModule) > 0 {
		rawModule := paramModule[0]
		rctx.Module = &rawModule
	}
	paramPage := req.Params["page"]
	if len(paramPage) > 0 {
		rawPage := paramPage[0]
		if page, err2 := strconv.Atoi(rawPage); err2 == nil {
			tmp4 := page
			tmp3 := &tmp4
			rctx.Page = tmp3
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("page", rawPage, "integer"))
		}
	}
	paramPageSize := req.Params["pageSize"]
	if len(paramPageSize) > 0 {
		rawPageSize := paramPageSize[0]
		if pageSize, err2 := strconv.Atoi(rawPageSize); err2 == nil {
			tmp6 := pageSize
			tmp5 := &tmp6
			rctx.PageSize = tmp5
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("pageSize", rawPageSize, "integer"))
		}
	}
	paramProfile := req.Params["profile"]
	if len(paramProfile) > 0 {
		rawProfile := paramProfile[0]
		rctx.Profile = &rawProfile
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListFirmwareContext) OK(r *Firmwares) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.app.firmwares+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// DeviceDataDataContext provides the data device data action context.
type DeviceDataDataContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	DeviceID   string
	FirstBlock *int
	LastBlock  *int
	Page       *int
	PageSize   *int
}

// NewDeviceDataDataContext parses the incoming request URL and body, performs validations and creates the
// context used by the data controller device data action.
func NewDeviceDataDataContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeviceDataDataContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeviceDataDataContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramDeviceID := req.Params["deviceId"]
	if len(paramDeviceID) > 0 {
		rawDeviceID := paramDeviceID[0]
		rctx.DeviceID = rawDeviceID
	}
	paramFirstBlock := req.Params["firstBlock"]
	if len(paramFirstBlock) > 0 {
		rawFirstBlock := paramFirstBlock[0]
		if firstBlock, err2 := strconv.Atoi(rawFirstBlock); err2 == nil {
			tmp8 := firstBlock
			tmp7 := &tmp8
			rctx.FirstBlock = tmp7
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("firstBlock", rawFirstBlock, "integer"))
		}
	}
	paramLastBlock := req.Params["lastBlock"]
	if len(paramLastBlock) > 0 {
		rawLastBlock := paramLastBlock[0]
		if lastBlock, err2 := strconv.Atoi(rawLastBlock); err2 == nil {
			tmp10 := lastBlock
			tmp9 := &tmp10
			rctx.LastBlock = tmp9
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("lastBlock", rawLastBlock, "integer"))
		}
	}
	paramPage := req.Params["page"]
	if len(paramPage) > 0 {
		rawPage := paramPage[0]
		if page, err2 := strconv.Atoi(rawPage); err2 == nil {
			tmp12 := page
			tmp11 := &tmp12
			rctx.Page = tmp11
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("page", rawPage, "integer"))
		}
	}
	paramPageSize := req.Params["pageSize"]
	if len(paramPageSize) > 0 {
		rawPageSize := paramPageSize[0]
		if pageSize, err2 := strconv.Atoi(rawPageSize); err2 == nil {
			tmp14 := pageSize
			tmp13 := &tmp14
			rctx.PageSize = tmp13
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("pageSize", rawPageSize, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DeviceDataDataContext) OK(r *DeviceDataRecordsResponse) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.app.device.data+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeviceDataDataContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// DeviceSummaryDataContext provides the data device summary action context.
type DeviceSummaryDataContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	DeviceID string
}

// NewDeviceSummaryDataContext parses the incoming request URL and body, performs validations and creates the
// context used by the data controller device summary action.
func NewDeviceSummaryDataContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeviceSummaryDataContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeviceSummaryDataContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramDeviceID := req.Params["deviceId"]
	if len(paramDeviceID) > 0 {
		rawDeviceID := paramDeviceID[0]
		rctx.DeviceID = rawDeviceID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DeviceSummaryDataContext) OK(r *DeviceDataSummaryResponse) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.app.device.summary+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeviceSummaryDataContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ProjectGetIDPictureContext provides the picture project get id action context.
type ProjectGetIDPictureContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ProjectID int
}

// NewProjectGetIDPictureContext parses the incoming request URL and body, performs validations and creates the
// context used by the picture controller project get id action.
func NewProjectGetIDPictureContext(ctx context.Context, r *http.Request, service *goa.Service) (*ProjectGetIDPictureContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ProjectGetIDPictureContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProjectID := req.Params["projectId"]
	if len(paramProjectID) > 0 {
		rawProjectID := paramProjectID[0]
		if projectID, err2 := strconv.Atoi(rawProjectID); err2 == nil {
			rctx.ProjectID = projectID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("projectId", rawProjectID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ProjectGetIDPictureContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "image/png")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ProjectGetIDPictureContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// UserGetIDPictureContext provides the picture user get id action context.
type UserGetIDPictureContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID int
}

// NewUserGetIDPictureContext parses the incoming request URL and body, performs validations and creates the
// context used by the picture controller user get id action.
func NewUserGetIDPictureContext(ctx context.Context, r *http.Request, service *goa.Service) (*UserGetIDPictureContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UserGetIDPictureContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUserID := req.Params["userId"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userId", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *UserGetIDPictureContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "image/png")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *UserGetIDPictureContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// UserSaveIDPictureContext provides the picture user save id action context.
type UserSaveIDPictureContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID int
}

// NewUserSaveIDPictureContext parses the incoming request URL and body, performs validations and creates the
// context used by the picture controller user save id action.
func NewUserSaveIDPictureContext(ctx context.Context, r *http.Request, service *goa.Service) (*UserSaveIDPictureContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UserSaveIDPictureContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUserID := req.Params["userId"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userId", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *UserSaveIDPictureContext) OK(r *MediaReferenceResponse) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.app.media+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *UserSaveIDPictureContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// AddUserContext provides the user add action context.
type AddUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *AddUserPayload
}

// NewAddUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller add action.
func NewAddUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*AddUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := AddUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *AddUserContext) OK(r *User) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.app.user+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *AddUserContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *AddUserContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// AdminDeleteUserContext provides the user admin delete action context.
type AdminDeleteUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *AdminDeleteUserPayload
}

// NewAdminDeleteUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller admin delete action.
func NewAdminDeleteUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*AdminDeleteUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := AdminDeleteUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// adminDeleteUserPayload is the user admin delete action payload.
type adminDeleteUserPayload struct {
	Email    *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	Password *string `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *adminDeleteUserPayload) Validate() (err error) {
	if payload.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}
	if payload.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	return
}

// Publicize creates AdminDeleteUserPayload from adminDeleteUserPayload
func (payload *adminDeleteUserPayload) Publicize() *AdminDeleteUserPayload {
	var pub AdminDeleteUserPayload
	if payload.Email != nil {
		pub.Email = *payload.Email
	}
	if payload.Password != nil {
		pub.Password = *payload.Password
	}
	return &pub
}

// AdminDeleteUserPayload is the user admin delete action payload.
type AdminDeleteUserPayload struct {
	Email    string `form:"email" json:"email" yaml:"email" xml:"email"`
	Password string `form:"password" json:"password" yaml:"password" xml:"password"`
}

// Validate runs the validation rules defined in the design.
func (payload *AdminDeleteUserPayload) Validate() (err error) {
	if payload.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}
	if payload.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *AdminDeleteUserContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *AdminDeleteUserContext) Unauthorized(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// Forbidden sends a HTTP response with status code 403.
func (ctx *AdminDeleteUserContext) Forbidden(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 403, r)
}

// ChangePasswordUserContext provides the user change password action context.
type ChangePasswordUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID  int
	Payload *UpdateUserPasswordPayload
}

// NewChangePasswordUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller change password action.
func NewChangePasswordUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*ChangePasswordUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ChangePasswordUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUserID := req.Params["userId"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userId", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ChangePasswordUserContext) OK(r *User) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.app.user+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ChangePasswordUserContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// Forbidden sends a HTTP response with status code 403.
func (ctx *ChangePasswordUserContext) Forbidden() error {
	ctx.ResponseData.WriteHeader(403)
	return nil
}

// GetCurrentUserContext provides the user get current action context.
type GetCurrentUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewGetCurrentUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller get current action.
func NewGetCurrentUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*GetCurrentUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GetCurrentUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetCurrentUserContext) OK(r *User) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.app.user+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ListByProjectUserContext provides the user list by project action context.
type ListByProjectUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ProjectID string
}

// NewListByProjectUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller list by project action.
func NewListByProjectUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListByProjectUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListByProjectUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramProjectID := req.Params["projectId"]
	if len(paramProjectID) > 0 {
		rawProjectID := paramProjectID[0]
		rctx.ProjectID = rawProjectID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListByProjectUserContext) OK(r *ProjectUsers) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.app.users+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// LoginUserContext provides the user login action context.
type LoginUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *LoginPayload
}

// NewLoginUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller login action.
func NewLoginUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*LoginUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := LoginUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *LoginUserContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *LoginUserContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *LoginUserContext) Unauthorized(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// LogoutUserContext provides the user logout action context.
type LogoutUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewLogoutUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller logout action.
func NewLogoutUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*LogoutUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := LogoutUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *LogoutUserContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *LogoutUserContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// ProjectRolesUserContext provides the user project roles action context.
type ProjectRolesUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewProjectRolesUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller project roles action.
func NewProjectRolesUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*ProjectRolesUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ProjectRolesUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ProjectRolesUserContext) OK(r ProjectRoleCollection) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.app.project.role+json; type=collection")
	}
	if r == nil {
		r = ProjectRoleCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// RecoveryUserContext provides the user recovery action context.
type RecoveryUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *RecoveryPayload
}

// NewRecoveryUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller recovery action.
func NewRecoveryUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*RecoveryUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := RecoveryUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *RecoveryUserContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *RecoveryUserContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// RecoveryLookupUserContext provides the user recovery lookup action context.
type RecoveryLookupUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *RecoveryLookupPayload
}

// NewRecoveryLookupUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller recovery lookup action.
func NewRecoveryLookupUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*RecoveryLookupUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := RecoveryLookupUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *RecoveryLookupUserContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *RecoveryLookupUserContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// RefreshUserContext provides the user refresh action context.
type RefreshUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *RefreshUserPayload
}

// NewRefreshUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller refresh action.
func NewRefreshUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*RefreshUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := RefreshUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// refreshUserPayload is the user refresh action payload.
type refreshUserPayload struct {
	RefreshToken *string `form:"refresh_token,omitempty" json:"refresh_token,omitempty" yaml:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *refreshUserPayload) Validate() (err error) {
	if payload.RefreshToken == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "refresh_token"))
	}
	return
}

// Publicize creates RefreshUserPayload from refreshUserPayload
func (payload *refreshUserPayload) Publicize() *RefreshUserPayload {
	var pub RefreshUserPayload
	if payload.RefreshToken != nil {
		pub.RefreshToken = *payload.RefreshToken
	}
	return &pub
}

// RefreshUserPayload is the user refresh action payload.
type RefreshUserPayload struct {
	RefreshToken string `form:"refresh_token" json:"refresh_token" yaml:"refresh_token" xml:"refresh_token"`
}

// Validate runs the validation rules defined in the design.
func (payload *RefreshUserPayload) Validate() (err error) {
	if payload.RefreshToken == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "refresh_token"))
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *RefreshUserContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *RefreshUserContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// SendValidationUserContext provides the user send validation action context.
type SendValidationUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID int
}

// NewSendValidationUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller send validation action.
func NewSendValidationUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*SendValidationUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := SendValidationUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUserID := req.Params["userId"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userId", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *SendValidationUserContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *SendValidationUserContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// TransmissionTokenUserContext provides the user transmission token action context.
type TransmissionTokenUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewTransmissionTokenUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller transmission token action.
func NewTransmissionTokenUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*TransmissionTokenUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := TransmissionTokenUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *TransmissionTokenUserContext) OK(r *TransmissionToken) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.app.user.transmission.token+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// UpdateUserContext provides the user update action context.
type UpdateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID  int
	Payload *UpdateUserPayload
}

// NewUpdateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller update action.
func NewUpdateUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*UpdateUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UpdateUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUserID := req.Params["userId"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userId", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateUserContext) OK(r *User) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.app.user+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ValidateUserContext provides the user validate action context.
type ValidateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Token string
}

// NewValidateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller validate action.
func NewValidateUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*ValidateUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ValidateUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramToken := req.Params["token"]
	if len(paramToken) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("token"))
	} else {
		rawToken := paramToken[0]
		rctx.Token = rawToken
	}
	return &rctx, err
}

// Found sends a HTTP response with status code 302.
func (ctx *ValidateUserContext) Found() error {
	ctx.ResponseData.WriteHeader(302)
	return nil
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ValidateUserContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}
