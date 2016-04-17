package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codeclysm/mara-api/auth"
	"github.com/codeclysm/rdbutils"
	"github.com/dancannon/gorethink"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gopkg.in/hlandau/passlib.v1"
)

type HandlersTestSuite struct {
	suite.Suite
	db     rdbutils.Database
	auth   auth.Client
	server *gin.Engine
}

type Table struct {
	Method       string
	Body         []byte
	ExpectedCode int
	ValidateBody func(*bytes.Buffer)
}

// TestLogin uses a table-driven test to test the Login handler
func (t *HandlersTestSuite) TestLogin() {

	table := map[string]Table{
		"wrong-method": Table{"GET", nil, 404, nil},
		"existing-user": Table{"POST", []byte(`{"username": "user", "password": "password"}`), 201, func(body *bytes.Buffer) {
			data := new(struct {
				Token string `json:"token"`
			})
			json.Unmarshal(body.Bytes(), data)
			token, err := jwt.Parse(data.Token, func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			})
			t.Nil(err)
			t.Equal(token.Claims["user"], "user")
		}},
		"not-existing-user": Table{"POST", []byte(`{"username": "user1", "password": "password"}`), 404, nil},
		"wrong-password":    Table{"POST", []byte(`{"username": "user", "password": "pass"}`), 400, nil},
	}
	t.TableTest(table)
}

func (t *HandlersTestSuite) TableTest(table map[string]Table) {
	for name, test := range table {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(test.Method, "/login", bytes.NewBuffer(test.Body))
		t.server.ServeHTTP(w, req)
		t.Equal(test.ExpectedCode, w.Code, name)
		if test.ValidateBody != nil {
			test.ValidateBody(w.Body)
		}
	}
}

func (t *HandlersTestSuite) SetupTest() {
	t.db = rdbutils.Database{Name: "test", Table: "users"}
	err := t.db.Connect()
	t.Nil(err)
	gorethink.DBCreate("test").RunWrite(t.db.Session)
	gorethink.DB("test").TableCreate("users").RunWrite(t.db.Session)
	t.auth = auth.Client{DB: &t.db, SigningKey: "secret"}

	r := Server{Auth: t.auth}
	t.server = r.New()

	hash, _ := passlib.Hash("password")
	user := auth.User{Username: "user", Password: hash}
	gorethink.DB("test").Table("users").Insert(user).Exec(t.db.Session)
}

func (t *HandlersTestSuite) TearDownTest() {
	gorethink.DB("test").Table("users").Delete().Exec(t.db.Session)
}

func TestHandlers(t *testing.T) {
	suite.Run(t, new(HandlersTestSuite))
}
