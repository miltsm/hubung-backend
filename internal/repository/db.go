package repository

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	PGUser         = "POSTGRES_USER"
	PGPassword     = "POSTGRES_PASSWORD"
	PGDb           = "POSTGRES_DB"
	PGPort         = "POSTGRES_PORT"
	PGHost         = "POSTGRES_HOST"
	PGPasswordPath = "POSTGRES_PASSWORD_FILE"
)

func establishDB() *sql.DB {

	POSTGRES_USER, POSTGRES_DB, POSTGRES_PORT, POSTGRES_HOST, POSTGRES_PASSWORD_FILE := os.Getenv(PGUser), os.Getenv(PGDb), os.Getenv(PGPort), os.Getenv(PGHost), os.Getenv(PGPasswordPath)

	//get password
	root, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	rootFs := os.DirFS(root)
	pwByte, err := fs.ReadFile(rootFs, POSTGRES_PASSWORD_FILE)
	if err != nil {
		panic(err.Error())
	}

	pswd := string(pwByte)

	url := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v", POSTGRES_USER, pswd, POSTGRES_HOST, POSTGRES_PORT, POSTGRES_DB)
	db, err := sql.Open("pgx", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	if err := db.PingContext(context.Background()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connected!")

	return db
}
