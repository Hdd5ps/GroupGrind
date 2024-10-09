package main

import (
  "database/sql"
  _ "github.com/lib/pq"
)

func connectDB() (*sql.DB, error) {
  connStr := "user=username dbname=groupgrind sslmode=disable"
  return sql.Open("postgres", connStr)
}
package main

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var db *gorm.DB

func init() {
  var err error
  dsn := "user=username password=password dbname=groupgrind port=5432 sslmode=disable"
  db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
  db.AutoMigrate(&User{})
}
