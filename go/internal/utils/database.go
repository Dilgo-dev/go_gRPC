package utils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Log struct {
  gorm.Model
  message  string
}

func Database() *gorm.DB {
  db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
  if err != nil {
    panic("Error: SQLITE file doesn't open")
  }
  return db 
}

func InitDatabase() {
  db := Database()
  db.AutoMigrate(&Log{})
}
