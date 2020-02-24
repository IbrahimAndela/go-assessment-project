package models

import (
	"github.com/jinzhu/gorm"
)

//ArticleModel is model that holds articels from the external server
type ArticleModel struct {
	ID          int    `gorm:"primary_key"`
	Title       string `gorm:"type:varchar(100)"`
	Body        string
	PublishedAt string         `gorm:"type:varchar(19)"`
	CreatedAt   string         `gorm:"type:varchar(19)"`
	Publisher   PublisherModel `gorm:"ForeignKeyId:id"`
	Category    CategoryModel  `gorm:"ForeignKeyId:id"`
	PublisherID int
	CategoryID  int
}

//Article used in marchling and unmarchling of articles from and to json
type Article struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	Publisher   string `json:"publisher"`
	Category    string `json:"category"`
	PublishedAt string `json:"published_at"`
	CreatedAt   string `json:"created_at"`
}

//Articles contains a list of articles
type Articles struct {
	Articles []Article
}

func (p *ArticleModel) create(db *gorm.DB, article *ArticleModel, name string, title string) error {

	if name != "" && title != "" {
		var publisher = GetOrCreatePublisher(db, name)
		var category = GetOrCreateCategory(db, title)
		db.Where("Name = ?", name).Find(&publisher)
		db.Where("Title = ?", category.Title).Find(&category)
		db.Create(&article).Related(&publisher).Related(&category)
	} else if name != "" {
		var publisher = GetOrCreatePublisher(db, name)
		db.Create(&article).Related(&publisher)
	} else if title != "" {
		var category = GetOrCreateCategory(db, title)
		db.Create(&article).Related(&category)
	} else {
		db.Create(&article)
	}
	return nil
}
