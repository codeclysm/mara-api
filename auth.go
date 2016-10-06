package main

import (
	"log"
	"math/rand"
	"time"

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
		User:     ctx.Payload.User,
		Password: ctx.Payload.Password,
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
	user, err := Auth.ByID(ctx.Payload.User)
	if err != nil {
		log.Println(err)
		return ctx.BadRequest()
	}
	password := random(16)
	user.SetPassword(password)
	log.Println("new password", password)
	err = Auth.Save(user)
	if err != nil {
		log.Println(err)
		return ctx.BadRequest()
	}
	return nil
}

// random generates a password of length n
func random(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const letterIdxBits = 6                    // 6 bits to represent a letter index
	const letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	const letterIdxMax = 63 / letterIdxBits    // # of letter indices fitting in 63 bits

	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
