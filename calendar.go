package main

import (
	"github.com/codeclysm/mara-api/app"
	"github.com/codeclysm/mara-api/calendar"
	"github.com/goadesign/goa"
	"github.com/juju/errors"
)

// CalendarController implements the calendar resource.
type CalendarController struct {
	*goa.Controller
}

// NewCalendarController creates a calendar controller.
func NewCalendarController(service *goa.Service) *CalendarController {
	return &CalendarController{Controller: service.NewController("CalendarController")}
}

// Create runs the create action.
func (c *CalendarController) Create(ctx *app.CreateCalendarContext) error {
	a := fromPayload(ctx.Payload)
	err := Calendar.Save(a)
	if err != nil {
		goa.LogError(ctx, err.Error())
		return ctx.InternalServerError()
	}

	return ctx.OK(toMedia(a))
}

// Delete runs the delete action.
func (c *CalendarController) Delete(ctx *app.DeleteCalendarContext) error {
	a, err := Calendar.Get(ctx.ID)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctx.NotFound()
		}
		goa.LogError(ctx, err.Error())
		return ctx.InternalServerError()
	}

	err = Calendar.Delete(a)
	if err != nil {
		goa.LogError(ctx, err.Error())
		return ctx.InternalServerError()
	}

	return ctx.OK(nil)
}

// Edit runs the edit action.
func (c *CalendarController) Edit(ctx *app.EditCalendarContext) error {
	old, err := Calendar.Get(ctx.ID)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctx.NotFound()
		}
		goa.LogError(ctx, err.Error())
		return ctx.InternalServerError()
	}

	new := fromPayload(ctx.Payload)

	old.Merge(new)
	err = Calendar.Save(old)
	if err != nil {
		goa.LogError(ctx, err.Error())
		return ctx.InternalServerError()
	}
	return ctx.OK(toMedia(old))
}

// List runs the list action.
func (c *CalendarController) List(ctx *app.ListCalendarContext) error {
	// CalendarController_List: start_implement

	// Put your logic here

	// CalendarController_List: end_implement
	res := app.MaraAppointmentCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *CalendarController) Show(ctx *app.ShowCalendarContext) error {
	a, err := Calendar.Get(ctx.ID)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctx.NotFound()
		}
		goa.LogError(ctx, err.Error())
		return ctx.InternalServerError()
	}

	return ctx.OK(toMedia(a))
}

func fromPayload(a *app.Appointment) *calendar.Appointment {
	ap := &calendar.Appointment{
		Who:         a.Who,
		What:        a.What,
		When:        a.When,
		Where:       a.Where,
		Problematic: a.Problematic,
		SendEmail:   a.SendEmail,
		SendSMS:     a.SendSms,
		Urgent:      a.Urgent,
	}
	if a.Email != nil {
		ap.Email = *a.Email
	}
	if a.Notes != nil {
		ap.Notes = *a.Notes
	}
	if a.Phone != nil {
		ap.Phone = *a.Phone
	}
	if a.Status != nil {
		ap.Status = *a.Status
	}
	return ap
}

func toMedia(a *calendar.Appointment) *app.MaraAppointment {
	href := app.CalendarHref(a.ID)
	return &app.MaraAppointment{
		ID:     &a.ID,
		Href:   &href,
		Who:    &a.Who,
		What:   &a.What,
		When:   &a.When,
		Where:  &a.Where,
		Status: &a.Status,

		Email:       &a.Email,
		Notes:       &a.Notes,
		Phone:       &a.Phone,
		Problematic: a.Problematic,
		SendEmail:   a.SendEmail,
		SendSms:     a.SendSMS,
		Urgent:      a.Urgent,
	}
}
