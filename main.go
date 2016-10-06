//go:generate goagen bootstrap -d github.com/codeclysm/mara-api/design

package main

import (
	"flag"
	"log"
	"net/smtp"
	"time"

	"github.com/codeclysm/mara-api/app"
	"github.com/codeclysm/mara-api/auth"
	"github.com/codeclysm/mara-api/calendar"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/goa/middleware/security/jwt"
	r "gopkg.in/dancannon/gorethink.v2"
)

var (
	port       = flag.String("port", ":9000", "The port where the api would listen")
	signingKey = flag.String("signingKey", "secret", "The key used to create jwt tokens")
	smtpHost   = flag.String("smtp-host", "smtp.gmail.com", "The smtp host")
	smtpUser   = flag.String("smtp-user", "user@gmail.com", "The smtp user")
	smtpPass   = flag.String("smtp-pass", "password", "The smtp password")

	// Auth is the client to handle authentication
	Auth auth.Client
	// Calendar is the client that handles the appointments
	Calendar calendar.Client
)

func main() {
	flag.Parse()
	// Create service
	service := goa.New("mara")

	// Connect to db
	opts := r.ConnectOpts{
		Database:   "mara",
		NumRetries: 3,
	}
	db, err := r.Connect(opts)
	if err != nil {
		panic("Missing database")
	}

	// Users Client
	Auth = auth.Client{DB: db, Table: "users", SigningKey: *signingKey}

	// Calendar Client
	Calendar = calendar.Client{DB: db, Table: "appointments"}

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount Jwt security
	app.UseJWTMiddleware(service, jwt.New(*signingKey, nil, app.NewJWTSecurity()))

	// Mount "auth" controller
	c := NewAuthController(service)
	app.MountAuthController(service, c)
	// Mount "calendar" controller
	c2 := NewCalendarController(service)
	app.MountCalendarController(service, c2)

	// Start loop to send notifications
	go loop()

	// Test email
	// ap := calendar.Appointment{Email: "matteo.suppo@gmail.com", SendEmail: true}
	// sendEmail(ap)

	// Start service
	if err := service.ListenAndServe(*port); err != nil {
		service.LogError("startup", "err", err)
	}
}

func loop() {
	interval := 1 * time.Minute
	c := time.Tick(interval)
	for now := range c {
		now = now.Add(24 * time.Hour)
		list, err := Calendar.Between("", now, now.Add(interval))
		if err != nil {
			log.Println(err.Error())
			continue
		}
		for _, ap := range list {
			sendSMS(ap)
			sendEmail(ap)
		}
	}
}

func sendSMS(ap calendar.Appointment) {

}

func sendEmail(ap calendar.Appointment) {
	if !ap.SendEmail || ap.Email == "" {
		return
	}

	auth := smtp.PlainAuth(
		"",
		*smtpUser,
		*smtpPass,
		*smtpHost,
	)
	log.Println(auth)
	err := smtp.SendMail(
		*smtpHost+":2525",
		auth,
		*smtpUser,
		[]string{ap.Email},
		[]byte("This is the email body."),
	)
	if err != nil {
		log.Fatal(err)
	}
}
