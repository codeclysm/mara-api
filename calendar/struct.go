package calendar

import "time"

// Appointment is defined by 4 why question plus some other collateral info.
type Appointment struct {
	ID     string    `gorethink:"id,omitempty"`
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

func (a *Appointment) Merge(new *Appointment) {
	if new.What != "" {
		a.What = new.What
	}
	time0 := time.Time{}
	if new.When != time0 {
		a.When = new.When
	}
	if new.Where != "" {
		a.Where = new.Where
	}
	if new.Who != "" {
		a.Who = new.Who
	}
	if new.Status != "" {
		a.Status = new.Status
	}
	if new.Email != "" {
		a.Email = new.Email
	}
	if new.Notes != "" {
		a.Notes = new.Notes
	}
	if new.Phone != "" {
		a.Phone = new.Phone
	}

	a.Problematic = new.Problematic
	a.SendEmail = new.SendEmail
	a.SendSMS = new.SendSMS
	a.Urgent = new.Urgent
}
