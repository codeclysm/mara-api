package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var JWT = JWTSecurity("jwt", func() {
	Header("Authorization")
})

var _ = Resource("calendar", func() {
	BasePath("/appointments")
	DefaultMedia(AppointmentMedia)

	Security(JWT, func() {
	})

	Action("list", func() {
		Description("show a list of appointments for the selected week")
		Routing(GET(""))
		Response(OK, CollectionOf(AppointmentMedia))
	})
	Action("create", func() {
		Description("create a new appointment")
		Routing(PUT(""))
		Payload(AppointmentPayload)
		Response(OK)
		Response(BadRequest)
		Response(InternalServerError)
	})
	Action("show", func() {
		Description("show the details of a single appointment")
		Routing(GET("/:id"))
		Response(OK)
		Response(NotFound)
	})
	Action("edit", func() {
		Description("change properties of a single appointment")
		Routing(POST("/:id"))
		Payload(AppointmentPayload)
		Response(OK)
		Response(BadRequest)
	})
	Action("delete", func() {
		Description("deletes a single appointment")
		Routing(DELETE("/:id"))
		Response(OK)
		Response(BadRequest)
	})
})

var AppointmentPayload = Type("Appointment", func() {
	Attribute("who", String, "Who is requesting the appointment", func() {
		Example("Pippo Pippis")
	})
	Attribute("email", String, "The email of who's requesting it", func() {
		Example("user@example.com")
		Format("email")
	})
	Attribute("send_email", Boolean, "Option to send email to the client", func() {
		Default(false)
	})
	Attribute("phone", String, "The phone of the client", func() {
		Example("3347662313")
	})
	Attribute("send_sms", Boolean, "Option to send sms to the client", func() {
		Default(false)
	})
	Attribute("what", String, "What is the purpose of the appointment", func() {
		Example("Convergenza")
	})
	Attribute("when", DateTime, "When is the appointment scheduled")
	Attribute("where", String, "Where is the appointment scheduled", func() {
		Example("Budrio")
	})
	Attribute("notes", String, "Internal notes about the Appointment")
	Attribute("urgent", Boolean, "Option specifying that the work is to be done now", func() {
		Default(false)
	})
	Attribute("problematic", Boolean, "Option specifying that the work is problematic for some reason", func() {
		Default(false)
	})
	Attribute("status", String, "The status of the work. Can be one of todo/doing/done", func() {
		Example("todo")
		Pattern("todo|doing|done")
	})
	Required("who", "what", "when", "where")
})

var AppointmentMedia = MediaType("application/vnd.mara.appointment", func() {
	Attributes(func() {
		Attribute("id", UUID, "The id of the appointment")
		Attribute("href", String, "The url of the appointment")
		Attribute("who", String, "Who is requesting the appointment", func() {
			Example("Pippo Pippis")
		})
		Attribute("email", String, "The email of who's requesting it", func() {
			Example("user@example.com")
			Format("email")
		})
		Attribute("send_email", Boolean, "Option to send email to the client", func() {
			Default(false)
		})
		Attribute("phone", String, "The phone of the client", func() {
			Example("3347662313")
		})
		Attribute("send_sms", Boolean, "Option to send sms to the client", func() {
			Default(false)
		})
		Attribute("what", String, "What is the purpose of the appointment", func() {
			Example("Convergenza")
		})
		Attribute("when", DateTime, "When is the appointment scheduled")
		Attribute("where", String, "Where is the appointment scheduled", func() {
			Example("Budrio")
		})
		Attribute("notes", String, "Internal notes about the Appointment")
		Attribute("urgent", Boolean, "Option specifying that the work is to be done now", func() {
			Default(false)
		})
		Attribute("problematic", Boolean, "Option specifying that the work is problematic for some reason", func() {
			Default(false)
		})
		Attribute("status", String, "The status of the work. Can be one of todo/doing/done", func() {
			Example("todo")
			Pattern("todo|doing|done")
		})
	})
	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("who")
		Attribute("email")
		Attribute("send_email")
		Attribute("phone")
		Attribute("send_sms")
		Attribute("what")
		Attribute("when")
		Attribute("where")
		Attribute("notes")
		Attribute("urgent")
		Attribute("problematic")
		Attribute("status")
	})
})
