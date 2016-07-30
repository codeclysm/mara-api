package calendar

import "time"

// Appointment is defined by 4 why question plus some other collateral info.
type Appointment struct {
	What   string    `gorethink:"what"`
	When   time.Time `gorethink:"when"`
	Where  string    `gorethink:"where"`
	Who    string    `gorethink:"who"`
	Status string    `gorethink:"status"`

	Email       string `gorethink:"email"`
	Notes       string `gorethink:"notes"`
	Phone       string `gorethink:"phone"`
	Problematic bool   `gorethink:"problematic"`
	SendEmail   bool   `gorethink:"send_email"`
	SendSMS     bool   `gorethink:"send_sms"`
	Urgent      bool   `gorethink:"urgent"`
}
