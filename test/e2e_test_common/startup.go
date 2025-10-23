package e2e_test_common

import (
	"context"
	"escalateservice/cmd/startup"
	"net/http"
	"testing"
	"time"
)

var app *startup.App

func SetUpE2E(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	app = &startup.App{
		Ctx:     ctx,
		Cancel:  cancel,
		Env:     Env,
		ConnStr: ConnStr,
	}

	go app.Startup()

	waitForPing(t)
}

func waitForPing(t *testing.T) {
	const maxRetries = 20
	const delay = 1000 * time.Millisecond
	url := "http://localhost:3333/test/escalateservice/ping"

	for i := 0; i < maxRetries; i++ {
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == http.StatusOK {
			resp.Body.Close()
			return
		}
		if resp != nil {
			resp.Body.Close()
		}
		time.Sleep(delay)
	}
	t.Fatalf("Timed out waiting for ping endpoint to become available at %s", url)
}
