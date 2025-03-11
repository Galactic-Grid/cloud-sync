package db

import (
	"fmt"
	"sync"
	"time"

	"github.com/Galactic-Grid/cloud-sync/pkg/db/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	instance *gorm.DB
	once     sync.Once
	initErr  error
)

func create_conn() (*gorm.DB, error) {
	once.Do(func() {
		dsn := "postgres://postgres:yourpassword@localhost:5432/mydb?sslmode=disable"

		// Configure GORM
		config := &gorm.Config{
			NowFunc: func() time.Time {
				return time.Now().UTC()
			},
			Logger:      logger.Default.LogMode(logger.Info),
			PrepareStmt: true,
		}

		// Open connection
		db, err := gorm.Open(postgres.Open(dsn), config)
		if err != nil {
			initErr = fmt.Errorf("failed to connect to database: %v", err)
			return
		}

		// Configure connection pool
		sqlDB, err := db.DB()
		if err != nil {
			initErr = fmt.Errorf("failed to get underlying *sql.DB: %v", err)
			return
		}

		sqlDB.SetMaxIdleConns(20)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(30 * time.Minute)

		instance = db
	})

	if initErr != nil {
		return nil, initErr
	}

	if instance == nil {
		return nil, fmt.Errorf("failed to initialize database connection")
	}

	return instance, nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	if instance == nil {
		db, err := create_conn()
		if err != nil {
			panic(err)
		}
		return db
	}
	return instance
}

func SchemaInit() error {
	db := GetDB()
	schemaName := "cloud_sync"

	// Create schema
	if err := db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s;", schemaName)).Error; err != nil {
		return fmt.Errorf("failed to create schema: %v", err)
	}

	// Set search path
	if err := db.Exec(fmt.Sprintf("SET search_path TO %s;", schemaName)).Error; err != nil {
		return fmt.Errorf("failed to set search path: %v", err)
	}

	// Auto-migrate will create tables based on models
	if err := db.AutoMigrate(&model.Tenant{}); err != nil {
		return fmt.Errorf("failed to migrate tables: %v", err)
	}

	return nil
}
