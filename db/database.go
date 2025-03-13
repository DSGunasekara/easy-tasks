package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to the database:", err)
	}

	// Auto-migrate the Task model
	err = DB.AutoMigrate(&Task{}, &Project{})
	if err != nil {
		log.Fatal("❌ Failed to migrate database:", err)
	}

	fmt.Println("✅ Database connected and migrated successfully!")
}

type Project struct {
	gorm.Model
	Name  string
	Tasks []Task `gorm:"foreignKey:TripId"`
}

type Task struct {
	gorm.Model
	TripId uint
	Name   string
	Status string
}
