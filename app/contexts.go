//************************************************************************//
// API "mara": Application Contexts
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/codeclysm/mara-api/design
// --out=$(GOPATH)/src/github.com/codeclysm/mara-api
// --version=v0.2.dev
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
	Payload *LoginAuthPayload
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

// loginAuthPayload is the auth login action payload.
type loginAuthPayload struct {
	// The password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// The username
	User *string `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
}

// Publicize creates LoginAuthPayload from loginAuthPayload
func (payload *loginAuthPayload) Publicize() *LoginAuthPayload {
	var pub LoginAuthPayload
	if payload.Password != nil {
		pub.Password = payload.Password
	}
	if payload.User != nil {
		pub.User = payload.User
	}
	return &pub
}

// LoginAuthPayload is the auth login action payload.
type LoginAuthPayload struct {
	// The password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// The username
	User *string `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
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
	Payload *ResetAuthPayload
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

// resetAuthPayload is the auth reset action payload.
type resetAuthPayload struct {
	// The email that will receive the new password
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *resetAuthPayload) Validate() (err error) {
	if payload.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *payload.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`raw.email`, *payload.Email, goa.FormatEmail, err2))
		}
	}
	return
}

// Publicize creates ResetAuthPayload from resetAuthPayload
func (payload *resetAuthPayload) Publicize() *ResetAuthPayload {
	var pub ResetAuthPayload
	if payload.Email != nil {
		pub.Email = payload.Email
	}
	return &pub
}

// ResetAuthPayload is the auth reset action payload.
type ResetAuthPayload struct {
	// The email that will receive the new password
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *ResetAuthPayload) Validate() (err error) {
	if payload.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *payload.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`raw.email`, *payload.Email, goa.FormatEmail, err2))
		}
	}
	return
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
	Payload *CreateCalendarPayload
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

// createCalendarPayload is the calendar create action payload.
type createCalendarPayload struct {
	// The email of who's requesting it
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Internal notes about the Appointment
	Notes *string `form:"notes,omitempty" json:"notes,omitempty" xml:"notes,omitempty"`
	// The phone of the client
	Phone *string `form:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty"`
	// Option specifying that the work is problematic for some reason
	Problematic *bool `form:"problematic,omitempty" json:"problematic,omitempty" xml:"problematic,omitempty"`
	// Option to send email to the client
	SendEmail *bool `form:"send_email,omitempty" json:"send_email,omitempty" xml:"send_email,omitempty"`
	// Option to send sms to the client
	SendSms *bool `form:"send_sms,omitempty" json:"send_sms,omitempty" xml:"send_sms,omitempty"`
	// The status of the work. Can be one of todo/doing/done
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Option specifying that the work is to be done now
	Urgent *bool `form:"urgent,omitempty" json:"urgent,omitempty" xml:"urgent,omitempty"`
	// What is the purpose of the appointment
	What *string `form:"what,omitempty" json:"what,omitempty" xml:"what,omitempty"`
	// When is the appointment scheduled
	When *time.Time `form:"when,omitempty" json:"when,omitempty" xml:"when,omitempty"`
	// Where is the appointment scheduled
	Where *string `form:"where,omitempty" json:"where,omitempty" xml:"where,omitempty"`
	// Who is requesting the appointment
	Who *string `form:"who,omitempty" json:"who,omitempty" xml:"who,omitempty"`
}

// Finalize sets the default values defined in the design.
func (payload *createCalendarPayload) Finalize() {
	var defaultProblematic = false
	if payload.Problematic == nil {
		payload.Problematic = &defaultProblematic
	}
	var defaultSendEmail = false
	if payload.SendEmail == nil {
		payload.SendEmail = &defaultSendEmail
	}
	var defaultSendSms = false
	if payload.SendSms == nil {
		payload.SendSms = &defaultSendSms
	}
	var defaultUrgent = false
	if payload.Urgent == nil {
		payload.Urgent = &defaultUrgent
	}
}

// Validate runs the validation rules defined in the design.
func (payload *createCalendarPayload) Validate() (err error) {
	if payload.Who == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "who"))
	}
	if payload.What == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "what"))
	}
	if payload.When == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "when"))
	}
	if payload.Where == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "where"))
	}

	if payload.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *payload.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`raw.email`, *payload.Email, goa.FormatEmail, err2))
		}
	}
	if payload.Status != nil {
		if ok := goa.ValidatePattern(`todo|doing|done`, *payload.Status); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`raw.status`, *payload.Status, `todo|doing|done`))
		}
	}
	return
}

// Publicize creates CreateCalendarPayload from createCalendarPayload
func (payload *createCalendarPayload) Publicize() *CreateCalendarPayload {
	var pub CreateCalendarPayload
	if payload.Email != nil {
		pub.Email = payload.Email
	}
	if payload.Notes != nil {
		pub.Notes = payload.Notes
	}
	if payload.Phone != nil {
		pub.Phone = payload.Phone
	}
	if payload.Problematic != nil {
		pub.Problematic = *payload.Problematic
	}
	if payload.SendEmail != nil {
		pub.SendEmail = *payload.SendEmail
	}
	if payload.SendSms != nil {
		pub.SendSms = *payload.SendSms
	}
	if payload.Status != nil {
		pub.Status = payload.Status
	}
	if payload.Urgent != nil {
		pub.Urgent = *payload.Urgent
	}
	if payload.What != nil {
		pub.What = *payload.What
	}
	if payload.When != nil {
		pub.When = *payload.When
	}
	if payload.Where != nil {
		pub.Where = *payload.Where
	}
	if payload.Who != nil {
		pub.Who = *payload.Who
	}
	return &pub
}

// CreateCalendarPayload is the calendar create action payload.
type CreateCalendarPayload struct {
	// The email of who's requesting it
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Internal notes about the Appointment
	Notes *string `form:"notes,omitempty" json:"notes,omitempty" xml:"notes,omitempty"`
	// The phone of the client
	Phone *string `form:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty"`
	// Option specifying that the work is problematic for some reason
	Problematic bool `form:"problematic" json:"problematic" xml:"problematic"`
	// Option to send email to the client
	SendEmail bool `form:"send_email" json:"send_email" xml:"send_email"`
	// Option to send sms to the client
	SendSms bool `form:"send_sms" json:"send_sms" xml:"send_sms"`
	// The status of the work. Can be one of todo/doing/done
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Option specifying that the work is to be done now
	Urgent bool `form:"urgent" json:"urgent" xml:"urgent"`
	// What is the purpose of the appointment
	What string `form:"what" json:"what" xml:"what"`
	// When is the appointment scheduled
	When time.Time `form:"when" json:"when" xml:"when"`
	// Where is the appointment scheduled
	Where string `form:"where" json:"where" xml:"where"`
	// Who is requesting the appointment
	Who string `form:"who" json:"who" xml:"who"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateCalendarPayload) Validate() (err error) {
	if payload.Who == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "who"))
	}
	if payload.What == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "what"))
	}
	if payload.Where == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "where"))
	}

	if payload.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *payload.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`raw.email`, *payload.Email, goa.FormatEmail, err2))
		}
	}
	if payload.Status != nil {
		if ok := goa.ValidatePattern(`todo|doing|done`, *payload.Status); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`raw.status`, *payload.Status, `todo|doing|done`))
		}
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateCalendarContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateCalendarContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
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

// BadRequest sends a HTTP response with status code 400.
func (ctx *DeleteCalendarContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// EditCalendarContext provides the calendar edit action context.
type EditCalendarContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      string
	Payload *EditCalendarPayload
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

// editCalendarPayload is the calendar edit action payload.
type editCalendarPayload struct {
	// The email of who's requesting it
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Internal notes about the Appointment
	Notes *string `form:"notes,omitempty" json:"notes,omitempty" xml:"notes,omitempty"`
	// The phone of the client
	Phone *string `form:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty"`
	// Option specifying that the work is problematic for some reason
	Problematic *bool `form:"problematic,omitempty" json:"problematic,omitempty" xml:"problematic,omitempty"`
	// Option to send email to the client
	SendEmail *bool `form:"send_email,omitempty" json:"send_email,omitempty" xml:"send_email,omitempty"`
	// Option to send sms to the client
	SendSms *bool `form:"send_sms,omitempty" json:"send_sms,omitempty" xml:"send_sms,omitempty"`
	// The status of the work. Can be one of todo/doing/done
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Option specifying that the work is to be done now
	Urgent *bool `form:"urgent,omitempty" json:"urgent,omitempty" xml:"urgent,omitempty"`
	// What is the purpose of the appointment
	What *string `form:"what,omitempty" json:"what,omitempty" xml:"what,omitempty"`
	// When is the appointment scheduled
	When *time.Time `form:"when,omitempty" json:"when,omitempty" xml:"when,omitempty"`
	// Where is the appointment scheduled
	Where *string `form:"where,omitempty" json:"where,omitempty" xml:"where,omitempty"`
	// Who is requesting the appointment
	Who *string `form:"who,omitempty" json:"who,omitempty" xml:"who,omitempty"`
}

// Finalize sets the default values defined in the design.
func (payload *editCalendarPayload) Finalize() {
	var defaultProblematic = false
	if payload.Problematic == nil {
		payload.Problematic = &defaultProblematic
	}
	var defaultSendEmail = false
	if payload.SendEmail == nil {
		payload.SendEmail = &defaultSendEmail
	}
	var defaultSendSms = false
	if payload.SendSms == nil {
		payload.SendSms = &defaultSendSms
	}
	var defaultUrgent = false
	if payload.Urgent == nil {
		payload.Urgent = &defaultUrgent
	}
}

// Validate runs the validation rules defined in the design.
func (payload *editCalendarPayload) Validate() (err error) {
	if payload.Who == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "who"))
	}
	if payload.What == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "what"))
	}
	if payload.When == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "when"))
	}
	if payload.Where == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "where"))
	}

	if payload.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *payload.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`raw.email`, *payload.Email, goa.FormatEmail, err2))
		}
	}
	if payload.Status != nil {
		if ok := goa.ValidatePattern(`todo|doing|done`, *payload.Status); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`raw.status`, *payload.Status, `todo|doing|done`))
		}
	}
	return
}

// Publicize creates EditCalendarPayload from editCalendarPayload
func (payload *editCalendarPayload) Publicize() *EditCalendarPayload {
	var pub EditCalendarPayload
	if payload.Email != nil {
		pub.Email = payload.Email
	}
	if payload.Notes != nil {
		pub.Notes = payload.Notes
	}
	if payload.Phone != nil {
		pub.Phone = payload.Phone
	}
	if payload.Problematic != nil {
		pub.Problematic = *payload.Problematic
	}
	if payload.SendEmail != nil {
		pub.SendEmail = *payload.SendEmail
	}
	if payload.SendSms != nil {
		pub.SendSms = *payload.SendSms
	}
	if payload.Status != nil {
		pub.Status = payload.Status
	}
	if payload.Urgent != nil {
		pub.Urgent = *payload.Urgent
	}
	if payload.What != nil {
		pub.What = *payload.What
	}
	if payload.When != nil {
		pub.When = *payload.When
	}
	if payload.Where != nil {
		pub.Where = *payload.Where
	}
	if payload.Who != nil {
		pub.Who = *payload.Who
	}
	return &pub
}

// EditCalendarPayload is the calendar edit action payload.
type EditCalendarPayload struct {
	// The email of who's requesting it
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Internal notes about the Appointment
	Notes *string `form:"notes,omitempty" json:"notes,omitempty" xml:"notes,omitempty"`
	// The phone of the client
	Phone *string `form:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty"`
	// Option specifying that the work is problematic for some reason
	Problematic bool `form:"problematic" json:"problematic" xml:"problematic"`
	// Option to send email to the client
	SendEmail bool `form:"send_email" json:"send_email" xml:"send_email"`
	// Option to send sms to the client
	SendSms bool `form:"send_sms" json:"send_sms" xml:"send_sms"`
	// The status of the work. Can be one of todo/doing/done
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Option specifying that the work is to be done now
	Urgent bool `form:"urgent" json:"urgent" xml:"urgent"`
	// What is the purpose of the appointment
	What string `form:"what" json:"what" xml:"what"`
	// When is the appointment scheduled
	When time.Time `form:"when" json:"when" xml:"when"`
	// Where is the appointment scheduled
	Where string `form:"where" json:"where" xml:"where"`
	// Who is requesting the appointment
	Who string `form:"who" json:"who" xml:"who"`
}

// Validate runs the validation rules defined in the design.
func (payload *EditCalendarPayload) Validate() (err error) {
	if payload.Who == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "who"))
	}
	if payload.What == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "what"))
	}
	if payload.Where == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "where"))
	}

	if payload.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *payload.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`raw.email`, *payload.Email, goa.FormatEmail, err2))
		}
	}
	if payload.Status != nil {
		if ok := goa.ValidatePattern(`todo|doing|done`, *payload.Status); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`raw.status`, *payload.Status, `todo|doing|done`))
		}
	}
	return
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

// ListCalendarContext provides the calendar list action context.
type ListCalendarContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListCalendarContext parses the incoming request URL and body, performs validations and creates the
// context used by the calendar controller list action.
func NewListCalendarContext(ctx context.Context, service *goa.Service) (*ListCalendarContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListCalendarContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListCalendarContext) OK(r MaraAppointmentCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.mara.appointment; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
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
