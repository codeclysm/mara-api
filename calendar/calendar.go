package calendar

import (
	"time"

	r "gopkg.in/dancannon/gorethink.v2"

	"github.com/juju/errors"
)

// Client is the way you use this module. Just instantiate it with a db instance
// and call its methods
type Client struct {
	DB    r.QueryExecutor
	Table string
}

// Between returns all the appointments in a certain location, after the time start and before the time end, included.
func (c *Client) Between(location string, start, end time.Time) ([]Appointment, error) {
	// Build the query
	query := r.Table(c.Table)
	if location != "" {
		query = r.Table(c.Table).Filter(map[string]string{"where": location})
	}

	query = query.Filter(func(row r.Term) r.Term {
		return row.Field("when").During(start, end)
	})

	// Execute the query
	cursor, err := query.Run(c.DB)
	if err != nil {
		return nil, errors.Annotatef(err, "while executing the query %v", query)
	}

	// Build the array
	appointments := []Appointment{}
	err = cursor.All(&appointments)
	if err != nil {
		return nil, errors.Annotatef(err, "while executing the query %v", query)
	}
	return appointments, nil
}

// Get an appointment with a specific id
func (c *Client) Get(id string) (*Appointment, error) {
	query := r.Table(c.Table).Get(id)
	cursor, err := query.Run(c.DB)
	if err != nil {
		return nil, errors.Annotatef(err, "while executing the query %v", query)
	}

	if cursor.IsNil() {
		return nil, errors.NotFoundf("with id %s", id)
	}

	var app Appointment
	err = cursor.One(&app)
	if err != nil {
		return nil, errors.Annotatef(err, "while unmarshaling cursor %v", cursor)
	}

	return &app, nil
}

// Save persists the appointment in database
func (c *Client) Save(app *Appointment) error {
	options := r.InsertOpts{Conflict: "replace"}
	query := r.Table(c.Table).Insert(app, options)
	_, err := query.RunWrite(c.DB)
	if err != nil {
		return errors.Annotatef(err, "while executing the query %v", query)
	}
	return nil
}

// Delete removes an appointment from the database
func (c *Client) Delete(app *Appointment) error {
	return r.Table(c.Table).Get(app.ID).Delete().Exec(c.DB)
}
