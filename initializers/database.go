package initializers

import (
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

func ConnectToDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("dbserver.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}
