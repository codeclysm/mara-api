package calendar

import (
	"time"

	"gopkg.in/dancannon/gorethink.v1"

	"github.com/codeclysm/rdbutils"
)

// Client is the way you use this module. Just instantiate it with a db instance
// and call its methods
type Client struct {
	DB *rdbutils.Database
}

// Between returns all the appointments in a certain location, after the time start and before the time end, included.
func (c *Client) Between(location string, start, end time.Time) ([]Appointment, error) {
	query := c.DB.Query().Filter(map[string]string{"where": location})

	query = query.Filter(func(row gorethink.Term) gorethink.Term {
		return row.Field("when").During(start, end)
	})
	cursor, err := c.DB.Run(query)
	if err != nil {
		panic(err)
		return nil, err
	}
	appointments := []Appointment{}
	err = cursor.All(&appointments)
	return appointments, err
}
