package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// set the addr to :4000 as a default value
	// For env variables, use os.Getenv()
	addr := flag.String("addr", ":4000", "HTTP network address")

	// parses -addr="..."
	flag.Parse()

	// create new logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app := &application{
		logger: logger,
	}

	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
