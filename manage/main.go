package main

import (
	"flag"
	"log"

	"gopkg.in/hlandau/passlib.v1"

	"github.com/codeclysm/mara-api/auth"
	"github.com/codeclysm/rdbutils"
	"github.com/dancannon/gorethink"
)

var (
	command    = flag.String("command", "setup", "The command to execute")
	signingKey = flag.String("signingKey", "secret", "The key used to create jwt tokens")
)

func main() {
	flag.Parse()
	dbu := rdbutils.Database{Name: "mara", Table: "users"}
	if err := dbu.Connect(); err != nil {
		panic("Missing database")
	}

	switch *command {
	case "setup":
		log.Print("Creating the database and tables")
		gorethink.DBCreate("mara").RunWrite(dbu.Session)
		gorethink.DB("mara").TableCreate("users").RunWrite(dbu.Session)
	case "admin":
		log.Print("Creating an admin user")
		hash, _ := passlib.Hash("password")
		user := auth.User{Username: "user", Password: hash}
		gorethink.DB("mara").Table("users").Insert(user).Exec(dbu.Session)
	}

}
