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
		Headers("Origin", "X-Requested-With", "Content-Type", "Accept")
	})
})
