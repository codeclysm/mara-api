// Package rdbutils provides some convenience functions wrapping https://github.com/dancannon/gorethink
package rdbutils

import (
	"time"

	"gopkg.in/dancannon/gorethink.v1"
)

// Interface is the contract the Database struct implements
type Interface interface {
	Query() gorethink.Term
	Run(gorethink.Term) (*gorethink.Cursor, error)
	RunWrite(gorethink.Term) (*gorethink.Cursor, error)
	Exec(gorethink.Term) error
}

// Run is a wrapper aroung gorethink queries.
// It tries the same query up to Retries times when there's a failure, and each time it waits a
// number of seconds longer.
//
// Usage
//
// When you have a query like `r.DB("test").Table("test").Run(conn)`
// you can trasform it into `rdbutils.Run(r.DB("test").Table("test"), conn, 3)`
func Run(query gorethink.Term, conn *gorethink.Session, Retries int) (*gorethink.Cursor, error) {
	cursor, err := query.Run(conn)
	if err != nil && Retries > 0 {
		timeToSleep := time.Duration(3 - Retries)
		time.Sleep(timeToSleep * time.Second)
		return Run(query, conn, Retries-1)
	}
	return cursor, err
}

// RunWrite is a wrapper aroung gorethink queries.
// It tries the same query up to Retries times when there's a failure, and each time it waits a
// number of seconds longer.
//
// Usage
//
// When you have a query like `r.DB("test").Table("test").RunWrite(conn)`
// you can trasform it into `utils.RunWrite(r.DB("test").Table("test"), conn, 3)`
func RunWrite(query gorethink.Term, conn *gorethink.Session, Retries int) (*gorethink.Cursor, error) {
	cursor, err := query.Run(conn)
	if err != nil && err.Error() == "EOF" && Retries > 0 {
		timeToSleep := time.Duration(3 - Retries)
		time.Sleep(timeToSleep * time.Second)
		return Run(query, conn, Retries-1)
	}
	return cursor, err
}

// Exec is a wrapper aroung gorethink queries.
// It tries the same query up to Retries times when there's a failure, and each time it waits a
// number of seconds longer.
//
// Usage
//
// When you have a query like `r.DB("test").Table("test").Exec(conn)`
// you can trasform it into `rdbutils.Exec(r.DB("test").Table("test"), conn, 3)`
func Exec(query gorethink.Term, conn *gorethink.Session, Retries int) error {
	err := query.Exec(conn)
	if err != nil && Retries > 0 {
		timeToSleep := time.Duration(3 - Retries)
		time.Sleep(timeToSleep * time.Second)
		return Exec(query, conn, Retries-1)
	}
	return err
}

// Database contains the informations about the rdbutils, the rethinkdb session,
// and the options such as the number of times it should retry a query to
// account for network failures
type Database struct {
	Host    string
	Key     string
	Name    string
	Table   string
	Retries int
	Session *gorethink.Session
}

// Connect establish a connection with the rethinkdb server and caches the
// connection
func (d *Database) Connect() error {
	var host string
	if d.Host == "" {
		host = "localhost:28015"
	} else {
		host = d.Host
	}

	var err error
	d.Session, err = gorethink.Connect(gorethink.ConnectOpts{
		Address: host,
		AuthKey: d.Key,
	})
	return err
}

// Query returns a default query from which to build your own.
func (d *Database) Query() gorethink.Term {
	return gorethink.DB(d.Name).Table(d.Table)
}

// Run is a wrapper around gorethink queries. It works like its sibling function
// with the same name, but it uses the cached session and the selected Retries
func (d *Database) Run(query gorethink.Term) (*gorethink.Cursor, error) {
	return Run(query, d.Session, d.Retries)
}

// RunWrite is a wrapper around gorethink queries. It works like its sibling function
// with the same name, but it uses the cached session and the selected Retries
func (d *Database) RunWrite(query gorethink.Term) (*gorethink.Cursor, error) {
	return RunWrite(query, d.Session, d.Retries)
}

// Exec is a wrapper around gorethink queries. It works like its sibling function
// with the same name, but it uses the cached session and the selected Retries
func (d *Database) Exec(query gorethink.Term) error {
	return Exec(query, d.Session, d.Retries)
}
