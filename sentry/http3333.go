package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
    Dsn:          "http://164e2f7807824e25b779850e0364ed2d@localhost/1",
		//Dsn:   "", // Set DSN here or set SENTRY_DSN environment variable
		Debug: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(time.Second)

	sentryMiddleware := sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	})

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// Important: Chi has a middleware stack and thus it is important to put the
	// Sentry handler on the appropriate place. If using middleware.Recoverer,
	// the Sentry middleware must come afterwards (and configure it with
	// Repanic: true).
	r.Use(sentryMiddleware.Handle)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})
	r.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		hub := sentry.GetHubFromContext(r.Context())
		hub.CaptureException(errors.New("test error"))
	})
	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("server panic")
	})

	http.ListenAndServe("localhost:3333", r)
}
