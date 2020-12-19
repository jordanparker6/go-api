// Package core defines core API objects and functions
// (e.g. database, security and config methods).
// The file database.go defines the database object.
package core

import (
	"database/sql"
	"time"

	driver "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ** Interface and Struct

type DatabaseI interface {
	Connect()
	Migrate(models []interface{})
	Orm() *gorm.DB
}

type database struct {
	pool   *sql.DB
	orm    *gorm.DB
	Config DatabaseConfig
}

// ** Database Methods

// Connect establishes a connection and builds a pool and orm.
func (db *database) Connect() {
	orm, err := gorm.Open(driver.Open(db.Config.Uri), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.orm = orm
	pool, err := db.orm.DB()
	db.pool = pool
	db.pool.SetMaxIdleConns(10)
	db.pool.SetMaxOpenConns(100)
	db.pool.SetConnMaxLifetime(time.Hour)
}

// Migrate uses the ORM system to build the required tables.
func (db *database) Migrate(models []interface{}) {
	db.orm.AutoMigrate(models...)
}

// Orm return the ORM object for use.
func (db *database) Orm() *gorm.DB {
	return db.orm
}

// ** Constructor

// NewDatabase is the database struct constructor.
func NewDatabase(dialect string, uri string) DatabaseI {
	db := new(database)
	db.Config = DatabaseConfig{dialect, uri}
	return db
}
