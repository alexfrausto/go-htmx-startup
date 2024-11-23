package config

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"io"
	"log"
	"os"
)

var DB *sqlx.DB

func connectDB() {
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.DbHost, config.DbUser, config.DbPass, config.DbName, config.DbPort, config.DbSSLMode,
	)

	DB, err = sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatalln(err)
	}
}

func InitializeDB() {
	connectDB()

	schema, err := readSchemaFile()
	if err != nil {
		log.Fatalf("Failed to read schema file: %s", err)
	}

	DB.MustExec(schema)
}

func readSchemaFile() (string, error) {
	file, err := os.Open("sql/schema.sql")
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(byteValue), nil
}
