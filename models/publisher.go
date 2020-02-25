package models

import (
	"github.com/jinzhu/gorm"
)

//PublisherModel maps publisher details to the database
type PublisherModel struct {
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"type:varchar(100)"`
}

//GetOrCreatePublisher get or create publisher
func GetOrCreatePublisher(db *gorm.DB, name string) PublisherModel {
	var publisher = PublisherModel{}
	db.Where("Name = ?", name).First(&publisher)
	if publisher.Name == "" {
		publisher.Name = name
		db.Create(&publisher)
	}
	return publisher
}
