package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/golang-migrate/migrate"
)

const DBTimeout = time.Second * 3

func Init(dbSource string) *sql.DB {
	conn, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal(err)

	}
	return conn
}

func RunDBMigration(migrationPath string, dbSource string) {
	migration, err := migrate.New(migrationPath, dbSource)
	if err != nil {
		log.Fatal("cannot create migration:", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up:", err)
	}

	log.Println("DB migrated successfully")

}
