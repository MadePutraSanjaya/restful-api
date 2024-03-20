// File: migrate.go

package main

import (
    "crud/initializers"
    "crud/models"
)

func main() {
    db, err := initializers.ConnectToDb()
    if err != nil {
        panic(err)
    }

    err = db.AutoMigrate(&models.Post{})
    if err != nil {
        panic(err)
    }
}
