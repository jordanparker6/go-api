// Defines a database struct and constructor
package core

import (
	"database/sql"
	"time"

	driver "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseI interface {
	Connect()
}

type database struct {
	pool   *sql.DB
	orm    *gorm.DB
	Config DatabaseConfig
}

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

	db.orm = orm
}

func NewDatabase(dialect string, uri string) DatabaseI {
	db := new(database)
	db.Config = DatabaseConfig{dialect, uri}
	return db
}
