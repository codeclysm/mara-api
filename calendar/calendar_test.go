package calendar_test

import (
	"testing"
	"time"

	r "gopkg.in/dancannon/gorethink.v2"

	"github.com/codeclysm/mara-api/calendar"
	"github.com/stretchr/testify/suite"
)

func TestCalendar(t *testing.T) {
	suite.Run(t, new(CalendarTestSuite))
}

type CalendarTestSuite struct {
	suite.Suite
	db     r.QueryExecutor
	client calendar.Client
}

func (t *CalendarTestSuite) SetupTest() {
	opts := r.ConnectOpts{
		Database: "test",
	}
	var err error
	t.db, err = r.Connect(opts)

	if err != nil {
		t.FailNow("Can't connect to database")
	}

	r.DBCreate("test").RunWrite(t.db)
	r.DB("test").TableCreate("appointments").RunWrite(t.db)
	t.client = calendar.Client{DB: t.db, Table: "appointments"}

	ap1 := calendar.Appointment{Where: "here", When: time.Date(2010, time.November, 10, 22, 0, 0, 0, time.UTC)}
	err = r.DB("test").Table("appointments").Insert(ap1).Exec(t.db)
	if err != nil {
		panic(err)
	}
	ap2 := calendar.Appointment{Where: "there", When: time.Date(2010, time.April, 10, 22, 0, 0, 0, time.UTC)}
	err = r.DB("test").Table("appointments").Insert(ap2).Exec(t.db)
	if err != nil {
		panic(err)
	}
	ap3 := calendar.Appointment{Where: "there", When: time.Date(2010, time.November, 10, 22, 0, 0, 0, time.UTC)}
	err = r.DB("test").Table("appointments").Insert(ap3).Exec(t.db)
	if err != nil {
		panic(err)
	}
}

func (t *CalendarTestSuite) TearDownTest() {
	r.DB("test").Table("appointments").Delete().Exec(t.db)
}

var TestBetweenData = []struct {
	ID             string
	Location       string
	Start          time.Time
	End            time.Time
	ExpectedNumber int
}{
	{
		"nothing", "here",
		time.Date(2009, time.November, 10, 22, 0, 0, 0, time.UTC),
		time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), 0,
	},

	{
		"2010", "here",
		time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2010, time.December, 31, 23, 59, 0, 0, time.UTC), 1,
	},
	{
		"2010", "there",
		time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2010, time.December, 31, 23, 59, 0, 0, time.UTC), 2,
	},
	{
		"november 2010", "there",
		time.Date(2010, time.November, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2010, time.November, 31, 23, 59, 0, 0, time.UTC), 1,
	},
}

func (t *CalendarTestSuite) TestBetween() {
	for _, test := range TestBetweenData {
		apps, err := t.client.Between(test.Location, test.Start, test.End)
		t.Nil(err)
		if err == nil {
			t.Equal(test.ExpectedNumber, len(apps))
		}
	}
}
