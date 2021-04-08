package database

import (
	"database/sql"
	"example-service/internal/config"
	"fmt"
	"log"
	"time"

	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Module ...
var Module = fx.Provide(NewDB)

// DB ...
type DB struct {
	Connection *gorm.DB
	sql        *sql.DB
	config     *config.Configuration
}

// NewDB is start connection database.
func NewDB(config *config.Configuration) (*DB, error) {
	// dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	// db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Silent),
	// })
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBName, config.DBPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if config.Environment != "production" {
		db.Logger = logger.Default.LogMode(logger.Info)
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	sqlDB.SetMaxOpenConns(7)
	sqlDB.SetMaxIdleConns(5)

	return &DB{
		Connection: db,
		sql:        sqlDB,
	}, nil
}

// Close Connection DB
func (db *DB) Close() {
	if err := db.sql.Close(); err != nil {
		log.Printf("Error closing db connection %s", err)
		return
	}
	log.Println("DB connection closed")
}
