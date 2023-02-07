package postgres_connector

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var db *sql.DB

func init() {
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	if POSTGRES_USER == "" {
		log.Panicln("POSTGRES_USER is empty")
	}

	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	if POSTGRES_PASSWORD == "" {
		log.Panicln("POSTGRES_PASSWORD is empty")
	}

	POSTGRES_DB := os.Getenv("POSTGRES_DB")
	if POSTGRES_DB == "" {
		log.Panicln("POSTGRES_DB is empty")
	}

	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	if POSTGRES_DB == "" {
		log.Panicln("POSTGRES_HOST is empty")
	}

	POSTGRES_PORT := os.Getenv("POSTGRES_PORT")
	if POSTGRES_DB == "" {
		log.Panicln("POSTGRES_PORT is empty")
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB, POSTGRES_HOST, POSTGRES_PORT)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	createSchemaString := `create schema if not exists web_control`
	createTableString := fmt.Sprintf(`create table if not exists web_control.user_password
(
    id            integer not null
        constraint id
            primary key,
    username      text    not null,
    password_hash text    not null
)`)

	_, err = db.Exec(createSchemaString)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(createTableString)
	if err != nil {
		panic(err)
	}

}
