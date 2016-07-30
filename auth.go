package main

import (
	"log"

	"github.com/codeclysm/mara-api/app"
	"github.com/codeclysm/mara-api/auth"
	"github.com/goadesign/goa"
)

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service) *AuthController {
	return &AuthController{Controller: service.NewController("AuthController")}
}

// Login runs the login action.
func (c *AuthController) Login(ctx *app.LoginAuthContext) error {
	data := auth.LoginData{
		User:     *ctx.Payload.User,
		Password: *ctx.Payload.Password,
	}
	token, err := Auth.Login(&data)
	if err != nil {
		log.Println(err.Error())
		return ctx.BadRequest()
	}
	res := &app.MaraToken{Token: &token}
	return ctx.OK(res)
}

// Reset runs the reset action.
func (c *AuthController) Reset(ctx *app.ResetAuthContext) error {
	// AuthController_Reset: start_implement

	// Put your logic here

	// AuthController_Reset: end_implement
	return nil
}
