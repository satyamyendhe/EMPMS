package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var (
	host     string
	port     string
	user     string
	password string
	dbname   string

	once sync.Once
	db   *sql.DB
)

// initDB initializes the database connection and creates tables if they do not exist
func initDB() error {
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASS")
	dbname = os.Getenv("DB_NAME")

	if strings.TrimSpace(host) == "" {
		host = "0.0.0.0"
	}
	if strings.TrimSpace(port) == "" {
		port = "5432"
	}
	if strings.TrimSpace(user) == "" {
		user = "satyam"
	}
	if strings.TrimSpace(password) == "" {
		password = "vsys!234"
	}
	if strings.TrimSpace(dbname) == "" {
		dbname = "empms"
	}

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		return fmt.Errorf("database environment variables are not properly set")
	}

	log.Printf("Connecting to database: host=%s port=%s user=%s dbname=%s", host, port, user, dbname)

	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlConn)
	if err != nil {
		return fmt.Errorf("error while making connection to DB: %w", err)
	}

	// Ping the database to ensure it is reachable
	if err = db.Ping(); err != nil {
		return fmt.Errorf("error while pinging the database: %w", err)
	}

	log.Println("Database connection established successfully")

	// Create tables if they do not exist
	schema := []string{
		`CREATE TABLE IF NOT EXISTS "user" (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			pass VARCHAR(255) NOT NULL,
			isAdmin BOOL NOT NULL,
			createdOn DATE DEFAULT CURRENT_DATE
		);`,
		`CREATE TABLE IF NOT EXISTS employee (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL CHECK (name <> ''),
			department VARCHAR(255) NOT NULL CHECK (department <> ''),
			designation VARCHAR(255) NOT NULL CHECK (designation <> ''),
			doj DATE NOT NULL,
			dob DATE NOT NULL,
			gender VARCHAR(50) NOT NULL CHECK (gender <> ''),
			email VARCHAR(255) NOT NULL UNIQUE CHECK (email <> ''),
			address VARCHAR(255) NOT NULL CHECK (address <> ''),
			mobile VARCHAR(50) NOT NULL CHECK (mobile <> ''),
			salary INT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS logs (
			id SERIAL PRIMARY KEY,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			emp_name VARCHAR(100) NOT NULL CHECK (emp_name <> ''),
			emp_email VARCHAR(255) NOT NULL CHECK (emp_email <> ''),
			emp_designation VARCHAR(255) NOT NULL CHECK (emp_designation <> ''),
			operation TEXT NOT NULL CHECK (operation <> ''),
			updated_by VARCHAR(100) NOT NULL CHECK (updated_by <> '')
		);`,
	}

	for _, stmt := range schema {
		if _, err = db.Exec(stmt); err != nil {
			return fmt.Errorf("error while creating tables: %w", err)
		}
	}

	return nil
}

// DBconn returns a singleton instance of the database connection, retrying on failure
func DBconn() *sql.DB {
	const maxTry = 5
	const retryInterval = 5 * time.Second

	var err error
	once.Do(func() {
		err = initDB()
	})

	for i := 0; i < maxTry; i++ {
		if err == nil {
			return db
		}
		log.Printf("Failed to connect to database, attempt %d/%d: %v", i+1, maxTry, err)
		time.Sleep(retryInterval)
		err = initDB() // Retry initialization
	}

	if err != nil {
		log.Fatalf("Unable to establish database connection after %d attempts: %v", maxTry, err)
	}

	return nil
}
