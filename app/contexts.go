//************************************************************************//
// API "mara": Application Contexts
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/codeclysm/mara-api/design
// --out=$(GOPATH)/src/github.com/codeclysm/mara-api
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"time"
)

// LoginAuthContext provides the auth login action context.
type LoginAuthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *Login
}

// NewLoginAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller login action.
func NewLoginAuthContext(ctx context.Context, service *goa.Service) (*LoginAuthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := LoginAuthContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *LoginAuthContext) OK(r *MaraToken) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.mara.token")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *LoginAuthContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// ResetAuthContext provides the auth reset action context.
type ResetAuthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *Reset
}

// NewResetAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller reset action.
func NewResetAuthContext(ctx context.Context, service *goa.Service) (*ResetAuthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ResetAuthContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ResetAuthContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ResetAuthContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// CreateCalendarContext provides the calendar create action context.
type CreateCalendarContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *Appointment
}

// NewCreateCalendarContext parses the incoming request URL and body, performs validations and creates the
// context used by the calendar controller create action.
func NewCreateCalendarContext(ctx context.Context, service *goa.Service) (*CreateCalendarContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateCalendarContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateCalendarContext) OK(r *MaraAppointment) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.mara.appointment")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateCalendarContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateCalendarContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// DeleteCalendarContext provides the calendar delete action context.
type DeleteCalendarContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewDeleteCalendarContext parses the incoming request URL and body, performs validations and creates the
// context used by the calendar controller delete action.
func NewDeleteCalendarContext(ctx context.Context, service *goa.Service) (*DeleteCalendarContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := DeleteCalendarContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DeleteCalendarContext) OK(r *MaraAppointment) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.mara.appointment")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteCalendarContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *DeleteCalendarContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// EditCalendarContext provides the calendar edit action context.
type EditCalendarContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      string
	Payload *Appointment
}

// NewEditCalendarContext parses the incoming request URL and body, performs validations and creates the
// context used by the calendar controller edit action.
func NewEditCalendarContext(ctx context.Context, service *goa.Service) (*EditCalendarContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := EditCalendarContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *EditCalendarContext) OK(r *MaraAppointment) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.mara.appointment")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *EditCalendarContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *EditCalendarContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *EditCalendarContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListCalendarContext provides the calendar list action context.
type ListCalendarContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	End   *time.Time
	Start *time.Time
	Where string
}

// NewListCalendarContext parses the incoming request URL and body, performs validations and creates the
// context used by the calendar controller list action.
func NewListCalendarContext(ctx context.Context, service *goa.Service) (*ListCalendarContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListCalendarContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramEnd := req.Params["end"]
	if len(paramEnd) > 0 {
		rawEnd := paramEnd[0]
		if end, err2 := time.Parse(time.RFC3339, rawEnd); err2 == nil {
			tmp1 := &end
			rctx.End = tmp1
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("end", rawEnd, "datetime"))
		}
	}
	paramStart := req.Params["start"]
	if len(paramStart) > 0 {
		rawStart := paramStart[0]
		if start, err2 := time.Parse(time.RFC3339, rawStart); err2 == nil {
			tmp2 := &start
			rctx.Start = tmp2
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("start", rawStart, "datetime"))
		}
	}
	paramWhere := req.Params["where"]
	if len(paramWhere) > 0 {
		rawWhere := paramWhere[0]
		rctx.Where = rawWhere
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListCalendarContext) OK(r MaraAppointmentCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.mara.appointment; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListCalendarContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowCalendarContext provides the calendar show action context.
type ShowCalendarContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewShowCalendarContext parses the incoming request URL and body, performs validations and creates the
// context used by the calendar controller show action.
func NewShowCalendarContext(ctx context.Context, service *goa.Service) (*ShowCalendarContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowCalendarContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowCalendarContext) OK(r *MaraAppointment) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.mara.appointment")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowCalendarContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowCalendarContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}
