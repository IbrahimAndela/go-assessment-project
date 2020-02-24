package models

import (
	"github.com/jinzhu/gorm"
)

//CategoryModel maps categories to the database
type CategoryModel struct {
	ID    int    `gorm:"primary_key"`
	Title string `gorm:"type:varchar(100)"`
}

func (p *CategoryModel) save(db *gorm.DB) (bool, error) {
	return true, nil
}

//GetOrCreateCategory get or create category
func GetOrCreateCategory(db *gorm.DB, title string) CategoryModel {
	var category = CategoryModel{}
	db.Where("Title = ?", title).First(&category)
	if category.Title == "" {
		category.Title = title
		db.Create(&category)
	}
	return category
}
