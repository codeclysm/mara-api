package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// LoginAuthPayload is the auth login action payload.
type LoginAuthPayload struct {
	// The password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// The username
	User *string `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
}

// LoginAuthPath computes a request path to the login action of auth.
func LoginAuthPath() string {
	return fmt.Sprintf("/auth/login")
}

// Login with username and password and obtain a token
func (c *Client) LoginAuth(ctx context.Context, path string, payload *LoginAuthPayload, contentType string) (*http.Response, error) {
	req, err := c.NewLoginAuthRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewLoginAuthRequest create the request corresponding to the login action endpoint of the auth resource.
func (c *Client) NewLoginAuthRequest(ctx context.Context, path string, payload *LoginAuthPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// ResetAuthPayload is the auth reset action payload.
type ResetAuthPayload struct {
	// The email that will receive the new password
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
}

// ResetAuthPath computes a request path to the reset action of auth.
func ResetAuthPath() string {
	return fmt.Sprintf("/auth/reset")
}

// Send a new password to the email specified
func (c *Client) ResetAuth(ctx context.Context, path string, payload *ResetAuthPayload, contentType string) (*http.Response, error) {
	req, err := c.NewResetAuthRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewResetAuthRequest create the request corresponding to the reset action endpoint of the auth resource.
func (c *Client) NewResetAuthRequest(ctx context.Context, path string, payload *ResetAuthPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
