package postgres

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresCredentials struct {
	host         string
	port         int
	databaseName string
	username     string
	password     string
	sslmode      string
}

var connection *PostgresCredentials

func GetPostgresCredentials() *PostgresCredentials {
	if connection == nil {
		port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
		if err != nil {
			log.Fatalf("Can't get POSTGRES_PORT from environment, it must be type int and not nil: %s\n", err)
		}
		connection = &PostgresCredentials{
			host:         os.Getenv("POSTGRES_HOST"),
			port:         port,
			databaseName: os.Getenv("POSTGRES_DATABASENAME"),
			username:     os.Getenv("POSTGRES_USERNAME"),
			password:     os.Getenv("POSTGRES_PASSWORD"),
			sslmode:      os.Getenv("SSLMODE"),
		}
	}
	return connection
}

func (pgcs *PostgresCredentials) PostgresConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		pgcs.host, pgcs.port, pgcs.databaseName, pgcs.username, pgcs.password, pgcs.sslmode,
	)
}

func NewPostgresDBConnection() *sqlx.DB {
	db, err := sqlx.Connect(os.Getenv("POSTGRES_DRIVER"), GetPostgresCredentials().PostgresConnectionString())
	if err != nil {
		log.Fatalf("Can't connect to Postgres database (blog): %s\n", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error during connection check in Postgres database (blog): %s\n", err)
	}
	log.Print("Connection to Postgres database (blog) successfully established\n")
	return db
}
