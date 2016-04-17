package main

import (
	"flag"

	"github.com/codeclysm/mara-api/auth"
	"github.com/codeclysm/rdbutils"
)

var (
	port       = flag.String("port", ":9000", "The port where the api would listen")
	signingKey = flag.String("signingKey", "secret", "The key used to create jwt tokens")
)

func main() {
	dbu := rdbutils.Database{Name: "mara", Table: "users"}
	if err := dbu.Connect(); err != nil {
		panic("Missing database")
	}
	authClient := auth.Client{DB: &dbu, SigningKey: *signingKey}
	routes := Server{Auth: authClient}
	server := routes.New()
	server.Run(*port)
}
