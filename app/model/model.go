package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Person struct {
	gorm.Model
	Name        string `gorm:"unique" json:"name"`
	Citizenship string `json:"citizenship"`
	Occupation  string `json:"occupation"`
	Age         int    `json:"age"`
	Status      bool   `json:"status"`
}

func (e *Person) Disable() {
	e.Status = false
}

func (p *Person) Enable() {
	p.Status = true
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Person{})
	return db
}
