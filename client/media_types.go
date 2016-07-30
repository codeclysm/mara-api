//************************************************************************//
// API "mara": Application Media Types
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/codeclysm/mara-api/design
// --out=$(GOPATH)/src/github.com/codeclysm/mara-api
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	uuid "github.com/goadesign/goa/uuid"
	"net/http"
	"time"
)

// MaraAppointment media type (default view)
//
// Identifier: application/vnd.mara.appointment
type MaraAppointment struct {
	// The email of who's requesting it
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// The url of the appointment
	Href *string `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
	// The id of the appointment
	ID *uuid.UUID `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
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
	What *string `form:"what,omitempty" json:"what,omitempty" xml:"what,omitempty"`
	// When is the appointment scheduled
	When *time.Time `form:"when,omitempty" json:"when,omitempty" xml:"when,omitempty"`
	// Where is the appointment scheduled
	Where *string `form:"where,omitempty" json:"where,omitempty" xml:"where,omitempty"`
	// Who is requesting the appointment
	Who *string `form:"who,omitempty" json:"who,omitempty" xml:"who,omitempty"`
}

// Validate validates the MaraAppointment media type instance.
func (mt *MaraAppointment) Validate() (err error) {
	if mt.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *mt.Email, goa.FormatEmail, err2))
		}
	}
	if mt.Status != nil {
		if ok := goa.ValidatePattern(`todo|doing|done`, *mt.Status); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.status`, *mt.Status, `todo|doing|done`))
		}
	}
	return
}

// DecodeMaraAppointment decodes the MaraAppointment instance encoded in resp body.
func (c *Client) DecodeMaraAppointment(resp *http.Response) (*MaraAppointment, error) {
	var decoded MaraAppointment
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// MaraAppointmentCollection is the media type for an array of MaraAppointment (default view)
//
// Identifier: application/vnd.mara.appointment; type=collection
type MaraAppointmentCollection []*MaraAppointment

// Validate validates the MaraAppointmentCollection media type instance.
func (mt MaraAppointmentCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Email != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Email); err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFormatError(`response[*].email`, *e.Email, goa.FormatEmail, err2))
			}
		}
		if e.Status != nil {
			if ok := goa.ValidatePattern(`todo|doing|done`, *e.Status); !ok {
				err = goa.MergeErrors(err, goa.InvalidPatternError(`response[*].status`, *e.Status, `todo|doing|done`))
			}
		}
	}
	return
}

// DecodeMaraAppointmentCollection decodes the MaraAppointmentCollection instance encoded in resp body.
func (c *Client) DecodeMaraAppointmentCollection(resp *http.Response) (MaraAppointmentCollection, error) {
	var decoded MaraAppointmentCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// MaraToken media type (default view)
//
// Identifier: application/vnd.mara.token
type MaraToken struct {
	// The token to use in subsequent api calls
	Token *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
}

// DecodeMaraToken decodes the MaraToken instance encoded in resp body.
func (c *Client) DecodeMaraToken(resp *http.Response) (*MaraToken, error) {
	var decoded MaraToken
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
