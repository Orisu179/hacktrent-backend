package main

import (
	"context"
	"flag"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"hacktrent.orisu179.com/internal/models"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger  *slog.Logger
	animals *models.AnimalModel
}

func main() {
	// set the addr to :4000 as a default value
	// For env variables, use os.Getenv()
	addr := flag.String("addr", ":4000", "HTTP network address")

	// parses -addr="..."
	flag.Parse()
	// create new logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	err := godotenv.Load()
	if err != nil {
		logger.Error("Doesn't work", "error", err.Error())
		os.Exit(1)
	}
	db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Error("Cannot connect to database", "error", err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := &application{
		logger:  logger,
		animals: &models.AnimalModel{DB: db},
	}

	err = app.createDb()
	if err != nil {
		logger.Error("Cannot init database", "error", err.Error())
		os.Exit(1)
	}

	logger.Info("starting server", "addr", *addr)
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDb(key string) (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), key)
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		err2 := db.Close(context.Background())
		if err2 != nil {
			return nil, err2
		}
		return nil, err
	}

	return db, nil
}
