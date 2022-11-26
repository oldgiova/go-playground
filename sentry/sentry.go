package main

import (
  "log"
  "time"
  "github.com/getsentry/sentry-go"
)

func() {
  defer sentry.Recover()
}

func main() {
  err := sentry.Init(sentry.ClientOptions{
    Dsn: "http://164e2f7807824e25b779850e0364ed2d@localhost/1",
    // Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
  })
  if err != nil {
    log.Fatalf("sentry.Init: %s", err)
  }
  // Flush buffered events before the program teminates.
  defer sentry.Flush(2 * time.Second)

}
