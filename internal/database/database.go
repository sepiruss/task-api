package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"

	"github.com/sepiruss/task-api/internal/config"
)

func Connect(cfg *config.Config) *sql.DB {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	var db *sql.DB
	var err error

	for i := 1; i <= 10; i++ {

		db, err = sql.Open("postgres", dsn)
		if err == nil {

			err = db.Ping()
			if err == nil {
				log.Println("Connected to PostgreSQL")
				return db
			}
		}

		log.Printf("Database not ready... retry %d/10", i)

		time.Sleep(3 * time.Second)
	}

	log.Fatal(err)

	return nil
}
