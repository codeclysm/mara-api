//************************************************************************//
// API "mara": Application User Types
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/codeclysm/mara-api/design
// --out=$(GOPATH)/src/github.com/codeclysm/mara-api
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"time"
)

// appointment user type.
type appointment struct {
	// The email of who's requesting it
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Internal notes about the Appointment
	Notes *string `form:"notes,omitempty" json:"notes,omitempty" xml:"notes,omitempty"`
	// The phone of the client
	Phone *string `form:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty"`
	// Option specifying that the work is problematic for some reason
	Problematic *bool `form:"problematic,omitempty" json:"problematic,omitempty" xml:"problematic,omitempty"`
	// Option to send email to the client
	SendEmail *bool `form:"send_email,omitempty" json:"send_email,omitempty" xml:"send_email,omitempty"`
	// Option to send sms to the client
	SendSms *bool `form:"send_sms,omitempty" json:"send_sms,omitempty" xml:"send_sms,omitempty"`
	// The status of the work. Can be one of todo/doing/done
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Option specifying that the work is to be done now
	Urgent *bool `form:"urgent,omitempty" json:"urgent,omitempty" xml:"urgent,omitempty"`
	// What is the purpose of the appointment
	What *string `form:"what,omitempty" json:"what,omitempty" xml:"what,omitempty"`
	// When is the appointment scheduled
	When *time.Time `form:"when,omitempty" json:"when,omitempty" xml:"when,omitempty"`
	// Where is the appointment scheduled
	Where *string `form:"where,omitempty" json:"where,omitempty" xml:"where,omitempty"`
	// Who is requesting the appointment
	Who *string `form:"who,omitempty" json:"who,omitempty" xml:"who,omitempty"`
}

// Finalize sets the default values for appointment type instance.
func (ut *appointment) Finalize() {
	var defaultProblematic = false
	if ut.Problematic == nil {
		ut.Problematic = &defaultProblematic
	}
	var defaultSendEmail = false
	if ut.SendEmail == nil {
		ut.SendEmail = &defaultSendEmail
	}
	var defaultSendSms = false
	if ut.SendSms == nil {
		ut.SendSms = &defaultSendSms
	}
	var defaultUrgent = false
	if ut.Urgent == nil {
		ut.Urgent = &defaultUrgent
	}
}

// Validate validates the appointment type instance.
func (ut *appointment) Validate() (err error) {
	if ut.Who == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "who"))
	}
	if ut.What == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "what"))
	}
	if ut.When == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "when"))
	}
	if ut.Where == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "where"))
	}

	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Status != nil {
		if ok := goa.ValidatePattern(`todo|doing|done`, *ut.Status); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.status`, *ut.Status, `todo|doing|done`))
		}
	}
	return
}

// Publicize creates Appointment from appointment
func (ut *appointment) Publicize() *Appointment {
	var pub Appointment
	if ut.Email != nil {
		pub.Email = ut.Email
	}
	if ut.Notes != nil {
		pub.Notes = ut.Notes
	}
	if ut.Phone != nil {
		pub.Phone = ut.Phone
	}
	if ut.Problematic != nil {
		pub.Problematic = *ut.Problematic
	}
	if ut.SendEmail != nil {
		pub.SendEmail = *ut.SendEmail
	}
	if ut.SendSms != nil {
		pub.SendSms = *ut.SendSms
	}
	if ut.Status != nil {
		pub.Status = ut.Status
	}
	if ut.Urgent != nil {
		pub.Urgent = *ut.Urgent
	}
	if ut.What != nil {
		pub.What = *ut.What
	}
	if ut.When != nil {
		pub.When = *ut.When
	}
	if ut.Where != nil {
		pub.Where = *ut.Where
	}
	if ut.Who != nil {
		pub.Who = *ut.Who
	}
	return &pub
}

// Appointment user type.
type Appointment struct {
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

// Validate validates the Appointment type instance.
func (ut *Appointment) Validate() (err error) {
	if ut.Who == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "who"))
	}
	if ut.What == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "what"))
	}
	if ut.Where == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "where"))
	}

	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Status != nil {
		if ok := goa.ValidatePattern(`todo|doing|done`, *ut.Status); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.status`, *ut.Status, `todo|doing|done`))
		}
	}
	return
}

// login user type.
type login struct {
	// The password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// The username
	User *string `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
}

// Publicize creates Login from login
func (ut *login) Publicize() *Login {
	var pub Login
	if ut.Password != nil {
		pub.Password = ut.Password
	}
	if ut.User != nil {
		pub.User = ut.User
	}
	return &pub
}

// Login user type.
type Login struct {
	// The password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// The username
	User *string `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
}

// reset user type.
type reset struct {
	// The user that will receive the new password
	User *string `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
}

// Publicize creates Reset from reset
func (ut *reset) Publicize() *Reset {
	var pub Reset
	if ut.User != nil {
		pub.User = ut.User
	}
	return &pub
}

// Reset user type.
type Reset struct {
	// The user that will receive the new password
	User *string `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
}
