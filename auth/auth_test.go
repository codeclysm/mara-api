package auth_test

import (
	"testing"

	"gopkg.in/hlandau/passlib.v1"

	"github.com/codeclysm/mara-api/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/suite"
	r "gopkg.in/dancannon/gorethink.v2"
)

type LoginTestSuite struct {
	suite.Suite
	db     r.QueryExecutor
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
	claims := token.Claims.(jwt.MapClaims)
	t.Equal(claims["user"], "user")
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
	opts := r.ConnectOpts{
		Database: "test",
	}
	var err error
	t.db, err = r.Connect(opts)

	if err != nil {
		t.FailNow("Can't connect to database")
	}

	r.DBCreate("test").RunWrite(t.db)
	r.DB("test").TableCreate("users").RunWrite(t.db)
	t.client = auth.Client{DB: t.db, Table: "users", SigningKey: "secret"}

	hash, _ := passlib.Hash("password")
	user := auth.User{Username: "user", Password: hash}
	r.DB("test").Table("users").Insert(user).Exec(t.db)
}

func (t *LoginTestSuite) TearDownTest() {
	r.DB("test").Table("users").Delete().Exec(t.db)
}

func TestLogin(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}
