package auth_test

import (
	"testing"

	"gopkg.in/hlandau/passlib.v1"

	"github.com/codeclysm/mara-api/auth"
	"github.com/codeclysm/rdbutils"
	"github.com/dancannon/gorethink"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/suite"
)

type LoginTestSuite struct {
	suite.Suite
	db     rdbutils.Database
	client auth.Client
}

// When you call auth.Client.Login() with an existing user
// It should return a valid jwt token
func (t *LoginTestSuite) TestLoginSuccess() {
	data := auth.LoginData{
		User:     "user",
		Password: "password",
	}
	tokenstring, err := t.client.Login(&data)
	t.Nil(err)
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	t.Nil(err)
	t.Equal(token.Claims["user"], "user")
}

// When you call auth.Client.Login() with a non existing user it should return
// an error
func (t *LoginTestSuite) TestLoginFailure() {
	data := auth.LoginData{
		User:     "notuser",
		Password: "password",
	}
	tokenstring, err := t.client.Login(&data)
	t.NotNil(err)
	t.Equal(tokenstring, "")
}

func (t *LoginTestSuite) SetupTest() {
	t.db = rdbutils.Database{Name: "test", Table: "users"}
	err := t.db.Connect()
	t.Nil(err)
	gorethink.DBCreate("test").RunWrite(t.db.Session)
	gorethink.DB("test").TableCreate("users").RunWrite(t.db.Session)
	t.client = auth.Client{DB: &t.db, SigningKey: "secret"}

	hash, _ := passlib.Hash("password")
	user := auth.User{Username: "user", Password: hash}
	gorethink.DB("test").Table("users").Insert(user).Exec(t.db.Session)
}

func (t *LoginTestSuite) TearDownTest() {
	gorethink.DB("test").Table("users").Delete().Exec(t.db.Session)
}

func TestLogin(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}
