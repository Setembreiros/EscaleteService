package main

import (
	"context"
	"escalateservice/cmd/startup"
	"os"
	"strings"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	env := strings.TrimSpace(os.Getenv("ENVIRONMENT"))
	connStr := strings.TrimSpace(os.Getenv("CONN_STR"))

	app := &startup.App{
		Ctx:     ctx,
		Cancel:  cancel,
		Env:     env,
		ConnStr: connStr,
	}

	app.Startup()
}
