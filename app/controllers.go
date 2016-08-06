//************************************************************************//
// API "mara": Application Controllers
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
	"github.com/goadesign/goa/cors"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// AuthController is the controller interface for the Auth actions.
type AuthController interface {
	goa.Muxer
	Login(*LoginAuthContext) error
	Reset(*ResetAuthContext) error
}

// MountAuthController "mounts" a Auth resource controller on the given service.
func MountAuthController(service *goa.Service, ctrl AuthController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/auth/login", ctrl.MuxHandler("preflight", handleAuthOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/auth/reset", ctrl.MuxHandler("preflight", handleAuthOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewLoginAuthContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*Login)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Login(rctx)
	}
	h = handleAuthOrigin(h)
	service.Mux.Handle("POST", "/auth/login", ctrl.MuxHandler("Login", h, unmarshalLoginAuthPayload))
	service.LogInfo("mount", "ctrl", "Auth", "action", "Login", "route", "POST /auth/login")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewResetAuthContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*Reset)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Reset(rctx)
	}
	h = handleAuthOrigin(h)
	service.Mux.Handle("POST", "/auth/reset", ctrl.MuxHandler("Reset", h, unmarshalResetAuthPayload))
	service.LogInfo("mount", "ctrl", "Auth", "action", "Reset", "route", "POST /auth/reset")
}

// handleAuthOrigin applies the CORS response headers corresponding to the origin.
func handleAuthOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, X-Requested-With, Content-Type, Accept")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalLoginAuthPayload unmarshals the request body into the context request data Payload field.
func unmarshalLoginAuthPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &login{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalResetAuthPayload unmarshals the request body into the context request data Payload field.
func unmarshalResetAuthPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &reset{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// CalendarController is the controller interface for the Calendar actions.
type CalendarController interface {
	goa.Muxer
	Create(*CreateCalendarContext) error
	Delete(*DeleteCalendarContext) error
	Edit(*EditCalendarContext) error
	List(*ListCalendarContext) error
	Show(*ShowCalendarContext) error
}

// MountCalendarController "mounts" a Calendar resource controller on the given service.
func MountCalendarController(service *goa.Service, ctrl CalendarController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/appointments", ctrl.MuxHandler("preflight", handleCalendarOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/appointments/:id", ctrl.MuxHandler("preflight", handleCalendarOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateCalendarContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*Appointment)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleCalendarOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("PUT", "/appointments", ctrl.MuxHandler("Create", h, unmarshalCreateCalendarPayload))
	service.LogInfo("mount", "ctrl", "Calendar", "action", "Create", "route", "PUT /appointments", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteCalendarContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleCalendarOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("DELETE", "/appointments/:id", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Calendar", "action", "Delete", "route", "DELETE /appointments/:id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewEditCalendarContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*Appointment)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Edit(rctx)
	}
	h = handleCalendarOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("POST", "/appointments/:id", ctrl.MuxHandler("Edit", h, unmarshalEditCalendarPayload))
	service.LogInfo("mount", "ctrl", "Calendar", "action", "Edit", "route", "POST /appointments/:id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListCalendarContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleCalendarOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("GET", "/appointments", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Calendar", "action", "List", "route", "GET /appointments", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowCalendarContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleCalendarOrigin(h)
	h = handleSecurity("jwt", h)
	service.Mux.Handle("GET", "/appointments/:id", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Calendar", "action", "Show", "route", "GET /appointments/:id", "security", "jwt")
}

// handleCalendarOrigin applies the CORS response headers corresponding to the origin.
func handleCalendarOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, X-Requested-With, Content-Type, Accept")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateCalendarPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateCalendarPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &appointment{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalEditCalendarPayload unmarshals the request body into the context request data Payload field.
func unmarshalEditCalendarPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &appointment{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// PublicController is the controller interface for the Public actions.
type PublicController interface {
	goa.Muxer
	goa.FileServer
}

// MountPublicController "mounts" a Public resource controller on the given service.
func MountPublicController(service *goa.Service, ctrl PublicController) {
	initService(service)
	var h goa.Handler

	h = ctrl.FileHandler("/builder/v1/swagger.json", "swagger/swagger.json")
	h = handlePublicOrigin(h)
	service.Mux.Handle("GET", "/builder/v1/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "swagger/swagger.json", "route", "GET /builder/v1/swagger.json")
}

// handlePublicOrigin applies the CORS response headers corresponding to the origin.
func handlePublicOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, X-Requested-With, Content-Type, Accept")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
