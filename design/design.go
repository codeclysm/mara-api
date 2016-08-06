package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("mara", func() {
	Title("An api to handle appointments")
	Scheme("http")
	Host("localhost:9000")
	Origin("*", func() {
		Methods("GET", "PUT", "POST", "DELETE")
		Headers("Authorization", "Origin", "X-Requested-With", "Content-Type", "Accept")
		Credentials()
	})
})

var _ = Resource("public", func() {

	Files("builder/v1/swagger.json", "swagger/swagger.json")
})
