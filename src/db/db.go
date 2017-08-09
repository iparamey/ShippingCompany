package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"errors"
)

type Config struct {
	Host      string
	Port      int
	User      string
	Password  string
	DBName    string
	SSLMode   string
}

func (config *Config) Default() {
	config.Host = "localhost"
	config.Port = 5432
	config.User = "postgres"
	config.Password = "postgres"
	config.DBName = "shipping_company"
	config.SSLMode = "disable"
}

var (
	ConnectionError = errors.New("Database connection failed")
	ValueError = errors.New("Input value is incorrect")
)

type DB struct {
	configuration *Config
	database *sql.DB
}

func (db *DB) Initialize (config *Config) error {
	db.configuration = config

	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		db.configuration.Host, db.configuration.Port, db.configuration.User, db.configuration.Password, db.configuration.DBName, db.configuration.SSLMode)

	db.database, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		//panic(err)
		return err
	}
	//fmt.Println("Successfully connected!")
	return nil
}

func (db *DB) GetConnection() *sql.DB {
	return db.database
}

func GetDBConnection() *sql.DB {
	config := Config{}
	config.Default()
	database := DB{}
	database.Initialize(&config)
	return database.GetConnection()
}

func (db *DB) CloseConnection() {
	if db.database != nil {
		db.database.Close()
		db.database = nil
	}
}
