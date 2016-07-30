package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"time"
)

// CreateCalendarPayload is the calendar create action payload.
type CreateCalendarPayload struct {
	// The email of who's requesting it
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Internal notes about the Appointment
	Notes *string `form:"notes,omitempty" json:"notes,omitempty" xml:"notes,omitempty"`
	// The phone of the client
	Phone *string `form:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty"`
	// Option specifying that the work is problematic for some reason
	Problematic bool `form:"problematic" json:"problematic" xml:"problematic"`
	// Option to send email to the client
	SendEmail bool `form:"send_email" json:"send_email" xml:"send_email"`
	// Option to send sms to the client
	SendSms bool `form:"send_sms" json:"send_sms" xml:"send_sms"`
	// The status of the work. Can be one of todo/doing/done
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Option specifying that the work is to be done now
	Urgent bool `form:"urgent" json:"urgent" xml:"urgent"`
	// What is the purpose of the appointment
	What string `form:"what" json:"what" xml:"what"`
	// When is the appointment scheduled
	When time.Time `form:"when" json:"when" xml:"when"`
	// Where is the appointment scheduled
	Where string `form:"where" json:"where" xml:"where"`
	// Who is requesting the appointment
	Who string `form:"who" json:"who" xml:"who"`
}

// CreateCalendarPath computes a request path to the create action of calendar.
func CreateCalendarPath() string {
	return fmt.Sprintf("/appointments")
}

// create a new appointment
func (c *Client) CreateCalendar(ctx context.Context, path string, payload *CreateCalendarPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateCalendarRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateCalendarRequest create the request corresponding to the create action endpoint of the calendar resource.
func (c *Client) NewCreateCalendarRequest(ctx context.Context, path string, payload *CreateCalendarPayload, contentType string) (*http.Request, error) {
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
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// DeleteCalendarPath computes a request path to the delete action of calendar.
func DeleteCalendarPath(id string) string {
	return fmt.Sprintf("/appointments/%v", id)
}

// deletes a single appointment
func (c *Client) DeleteCalendar(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteCalendarRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteCalendarRequest create the request corresponding to the delete action endpoint of the calendar resource.
func (c *Client) NewDeleteCalendarRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// EditCalendarPayload is the calendar edit action payload.
type EditCalendarPayload struct {
	// The email of who's requesting it
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Internal notes about the Appointment
	Notes *string `form:"notes,omitempty" json:"notes,omitempty" xml:"notes,omitempty"`
	// The phone of the client
	Phone *string `form:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty"`
	// Option specifying that the work is problematic for some reason
	Problematic bool `form:"problematic" json:"problematic" xml:"problematic"`
	// Option to send email to the client
	SendEmail bool `form:"send_email" json:"send_email" xml:"send_email"`
	// Option to send sms to the client
	SendSms bool `form:"send_sms" json:"send_sms" xml:"send_sms"`
	// The status of the work. Can be one of todo/doing/done
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Option specifying that the work is to be done now
	Urgent bool `form:"urgent" json:"urgent" xml:"urgent"`
	// What is the purpose of the appointment
	What string `form:"what" json:"what" xml:"what"`
	// When is the appointment scheduled
	When time.Time `form:"when" json:"when" xml:"when"`
	// Where is the appointment scheduled
	Where string `form:"where" json:"where" xml:"where"`
	// Who is requesting the appointment
	Who string `form:"who" json:"who" xml:"who"`
}

// EditCalendarPath computes a request path to the edit action of calendar.
func EditCalendarPath(id string) string {
	return fmt.Sprintf("/appointments/%v", id)
}

// change properties of a single appointment
func (c *Client) EditCalendar(ctx context.Context, path string, payload *EditCalendarPayload, contentType string) (*http.Response, error) {
	req, err := c.NewEditCalendarRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewEditCalendarRequest create the request corresponding to the edit action endpoint of the calendar resource.
func (c *Client) NewEditCalendarRequest(ctx context.Context, path string, payload *EditCalendarPayload, contentType string) (*http.Request, error) {
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
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ListCalendarPath computes a request path to the list action of calendar.
func ListCalendarPath() string {
	return fmt.Sprintf("/appointments")
}

// show a list of appointments for the selected week
func (c *Client) ListCalendar(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListCalendarRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListCalendarRequest create the request corresponding to the list action endpoint of the calendar resource.
func (c *Client) NewListCalendarRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ShowCalendarPath computes a request path to the show action of calendar.
func ShowCalendarPath(id string) string {
	return fmt.Sprintf("/appointments/%v", id)
}

// show the details of a single appointment
func (c *Client) ShowCalendar(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowCalendarRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowCalendarRequest create the request corresponding to the show action endpoint of the calendar resource.
func (c *Client) NewShowCalendarRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}
