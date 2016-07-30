// Package auth deals with the authentication of the user. It's fairly simple:
// A user has a username and a password, which is saved hashed in the database.
package auth

import (
	"fmt"
	"time"

	"github.com/codeclysm/rdbutils"
	"github.com/dgrijalva/jwt-go"
	"github.com/juju/errors"
	"gopkg.in/dancannon/gorethink.v1"
	"gopkg.in/hlandau/passlib.v1"
)

// Client is the way you use this module. Just instantiate it with a db instance
// and call its methods
type Client struct {
	DB         *rdbutils.Database
	SigningKey string
}

// Login exchanges a user and a password for a jwt token. It returns an error if
// the user doesn't exist (or if there is a problem with the database)
func (c *Client) Login(data *LoginData) (string, error) {
	user, err := c.ByID(data.User)
	if err != nil {
		return "", ErrorNotFound{User: data.User, Message: err.Error()}
	}
	_, err = passlib.Verify(data.Password, user.Password)
	if err != nil {
		return "", ErrorWrongPass{User: data.User, Password: data.Password, Message: err.Error()}
	}

	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["user"] = user.Username
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, err := token.SignedString([]byte(c.SigningKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ByID returns the user given the id
func (c *Client) ByID(id string) (*User, error) {
	cursor, err := c.DB.Run(c.DB.Query().Get(id))
	if err != nil {
		return nil, err
	}
	user := new(User)
	err = cursor.One(user)
	return user, err
}

// Save persists an user on the database
func (c *Client) Save(user *User) error {
	options := gorethink.InsertOpts{Conflict: "replace"}
	query := c.DB.Query().Insert(user, options)
	_, err := c.DB.RunWrite(query)
	if err != nil {
		return errors.Annotatef(err, "while executing the query %v", query)
	}
	return nil
}

// LoginData is a helper struct containing a username and a password. You can use
// it to marshal/unmarshal json data
type LoginData struct {
	User     string `json:"username"`
	Password string `json:"password"`
}

// User has a username and an hashed password
type User struct {
	Username string `gorethink:"id"`
	Password string `gorethink:"password"`
	Email    string `gorethink:"email"`
}

// SetPassword changes the password of the user
func (u *User) SetPassword(password string) {
	hash, _ := passlib.Hash(password)
	u.Password = hash
}

// ErrorNotFound is returned when the requested user is not present in the db
type ErrorNotFound struct {
	User    string
	Message string
}

func (e ErrorNotFound) Error() string { return fmt.Sprintf("NotFound user %s: %s", e.User, e.Message) }

// ErrorWrongPass is returned when the password doesn't match the one in the database
type ErrorWrongPass struct {
	User     string
	Password string
	Message  string
}

func (e ErrorWrongPass) Error() string {
	return fmt.Sprintf("WrongPass %s for user %s: %s", e.Password, e.User, e.Message)
}
