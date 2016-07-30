package calendar_test

import (
	"testing"
	"time"

	"gopkg.in/dancannon/gorethink.v1"

	"github.com/codeclysm/mara-api/calendar"
	"github.com/codeclysm/rdbutils"
	"github.com/stretchr/testify/suite"
)

func TestCalendar(t *testing.T) {
	suite.Run(t, new(CalendarTestSuite))
}

type CalendarTestSuite struct {
	suite.Suite
	db     rdbutils.Database
	client calendar.Client
}

func (t *CalendarTestSuite) SetupTest() {
	t.db = rdbutils.Database{Name: "test", Table: "appointments"}
	err := t.db.Connect()
	t.Nil(err)
	gorethink.DBCreate("test").RunWrite(t.db.Session)
	gorethink.DB("test").TableCreate("appointments").RunWrite(t.db.Session)
	t.client = calendar.Client{DB: &t.db}

	ap1 := calendar.Appointment{Where: "here", When: time.Date(2010, time.November, 10, 22, 0, 0, 0, time.UTC)}
	err = gorethink.DB("test").Table("appointments").Insert(ap1).Exec(t.db.Session)
	if err != nil {
		panic(err)
	}
	ap2 := calendar.Appointment{Where: "there", When: time.Date(2010, time.April, 10, 22, 0, 0, 0, time.UTC)}
	err = gorethink.DB("test").Table("appointments").Insert(ap2).Exec(t.db.Session)
	if err != nil {
		panic(err)
	}
	ap3 := calendar.Appointment{Where: "there", When: time.Date(2010, time.November, 10, 22, 0, 0, 0, time.UTC)}
	err = gorethink.DB("test").Table("appointments").Insert(ap3).Exec(t.db.Session)
	if err != nil {
		panic(err)
	}
}

func (t *CalendarTestSuite) TearDownTest() {
	gorethink.DB("test").Table("appointments").Delete().Exec(t.db.Session)
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
