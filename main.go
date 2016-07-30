//go:generate goagen bootstrap -d github.com/codeclysm/mara-api/design

package main

import (
	"github.com/codeclysm/mara-api/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("mara")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "auth" controller
	c := NewAuthController(service)
	app.MountAuthController(service, c)
	// Mount "calendar" controller
	c2 := NewCalendarController(service)
	app.MountCalendarController(service, c2)

	// Start service
	if err := service.ListenAndServe(":9000"); err != nil {
		service.LogError("startup", "err", err)
	}
}
