# rdbutils
--
    import "github.com/codeclysm/rdbutils"

Package rdbutils provides some convenience functions wrapping
https://github.com/dancannon/gorethink

## Usage

#### func  Exec

```go
func Exec(query gorethink.Term, conn *gorethink.Session, Retries int) error
```
Exec is a wrapper aroung gorethink queries. It tries the same query up to
Retries times when there's a failure, and each time it waits a number of seconds
longer.


Usage

When you have a query like `r.DB("test").Table("test").Exec(conn)` you can
trasform it into `rdbutils.Exec(r.DB("test").Table("test"), conn, 3)`

#### func  Run

```go
func Run(query gorethink.Term, conn *gorethink.Session, Retries int) (*gorethink.Cursor, error)
```
Run is a wrapper aroung gorethink queries. It tries the same query up to Retries
times when there's a failure, and each time it waits a number of seconds longer.


Usage

When you have a query like `r.DB("test").Table("test").Run(conn)` you can
trasform it into `rdbutils.Run(r.DB("test").Table("test"), conn, 3)`

#### func  RunWrite

```go
func RunWrite(query gorethink.Term, conn *gorethink.Session, Retries int) (*gorethink.Cursor, error)
```
RunWrite is a wrapper aroung gorethink queries. It tries the same query up to
Retries times when there's a failure, and each time it waits a number of seconds
longer.


Usage

When you have a query like `r.DB("test").Table("test").RunWrite(conn)` you can
trasform it into `utils.RunWrite(r.DB("test").Table("test"), conn, 3)`

#### type Database

```go
type Database struct {
	Host    string
	Key     string
	Name    string
	Table   string
	Retries int
	Session *gorethink.Session
}
```

Database contains the informations about the rdbutils, the rethinkdb session,
and the options such as the number of times it should retry a query to account
for network failures

#### func (*Database) Connect

```go
func (d *Database) Connect() error
```
Connect establish a connection with the rethinkdb server and caches the
connection

#### func (*Database) Exec

```go
func (d *Database) Exec(query gorethink.Term) error
```
Exec is a wrapper around gorethink queries. It works like its sibling function
with the same name, but it uses the cached session and the selected Retries

#### func (*Database) Query

```go
func (d *Database) Query() gorethink.Term
```
Query returns a default query from which to build your own.

#### func (*Database) Run

```go
func (d *Database) Run(query gorethink.Term) (*gorethink.Cursor, error)
```
Run is a wrapper around gorethink queries. It works like its sibling function
with the same name, but it uses the cached session and the selected Retries

#### func (*Database) RunWrite

```go
func (d *Database) RunWrite(query gorethink.Term) (*gorethink.Cursor, error)
```
RunWrite is a wrapper around gorethink queries. It works like its sibling
function with the same name, but it uses the cached session and the selected
Retries

#### type Interface

```go
type Interface interface {
	Query() gorethink.Term
	Run(gorethink.Term) (*gorethink.Cursor, error)
	RunWrite(gorethink.Term) (*gorethink.Cursor, error)
	Exec(gorethink.Term) error
}
```

Interface is the contract the Database struct implements
