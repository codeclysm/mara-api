package main

import (
	"github.com/codeclysm/mara-api/app"
	"github.com/goadesign/goa"
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
	// CalendarController_Create: start_implement

	// Put your logic here

	// CalendarController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *CalendarController) Delete(ctx *app.DeleteCalendarContext) error {
	// CalendarController_Delete: start_implement

	// Put your logic here

	// CalendarController_Delete: end_implement
	res := &app.MaraAppointment{}
	return ctx.OK(res)
}

// Edit runs the edit action.
func (c *CalendarController) Edit(ctx *app.EditCalendarContext) error {
	// CalendarController_Edit: start_implement

	// Put your logic here

	// CalendarController_Edit: end_implement
	res := &app.MaraAppointment{}
	return ctx.OK(res)
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
	// CalendarController_Show: start_implement

	// Put your logic here

	// CalendarController_Show: end_implement
	res := &app.MaraAppointment{}
	return ctx.OK(res)
}
