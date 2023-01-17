package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
	"useLaborartory-backend/repository"
	"useLaborartory-backend/repository/dbrepo"
)

const port = 8082

type application struct {
	DSN          string
	Domain       string
	DB           repository.DatabaseRepo
	auth         Auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
}

func main() {
	var app application
	flag.StringVar(&app.DSN, "dsn", "root:1234@tcp(127.0.0.1:3306)/use-laboratory?charset=utf8mb4&parseTime=True&loc=Local", "MariaDB Connection")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "verysecret", "singing secret")
	flag.StringVar(&app.JWTIssuer, "jwt-issuer", "example.com", "singing issuer")
	flag.StringVar(&app.JWTAudience, "jwt-audience", "example.com", "singing audience")
	flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "cookie domain")
	flag.StringVar(&app.Domain, "domain", "example.com", "domain")
	flag.Parse()

	conn, err := app.connectDatabase(app.DSN)
	app.DB = &dbrepo.MariaDBRepo{DB: conn}

	app.auth = Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookiePath:    "/",
		CookieName:    "__Host-refresh_token",
		CookieDomain:  app.CookieDomain,
	}
	log.Println("Starting application on port", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
