package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("auth", func() {
	BasePath("/auth")
	Action("login", func() {
		Description("Login with username and password and obtain a token")
		Routing(POST("/login"))
		Payload(LoginPayload)
		Response(OK, TokenMedia)
		Response(BadRequest)
	})
	Action("reset", func() {
		Description("Send a new password to the email specified")
		Routing(POST("/reset"))
		Payload(ResetPayload)
		Response(OK)
		Response(BadRequest)
	})
})

var LoginPayload = Type("Login", func() {
	Attribute("user", String, "The username", func() {
		Example("user")
	})
	Attribute("password", String, "The password", func() {
		Example("password")
	})
	Required("user", "password")
})

var ResetPayload = Type("Reset", func() {
	Attribute("user", String, "The user that will receive the new password", func() {
		Example("user")
	})
	Required("user")
})

var TokenMedia = MediaType("application/vnd.mara.token", func() {
	Attributes(func() {
		Attribute("token", String, "The token to use in subsequent api calls", func() {
			Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
		})
	})
	View("default", func() {
		Attribute("token")
	})
})
