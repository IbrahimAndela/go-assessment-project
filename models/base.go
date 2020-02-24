package models

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

//DatabaseInterface connects all database methods
type DatabaseInterface interface {
	CreateArticle(data []byte) (*ArticleModel, error)
	GetArticles(category string, publisher string, createdAt string, publishedAt string) ([]*ArticleModel, error)
	GetArticle(id uint) (*ArticleModel, error)
	UpdateArticle(data []byte) (*ArticleModel, error)
	DeleteArticle(id uint)
}

//Connector provides database interface
type Connector struct {
	Db *gorm.DB
}

//CreateArticle creates article
func (con *Connector) CreateArticle(data []byte) (*ArticleModel, error) {
	var jsonData Article
	json.Unmarshal(data, &jsonData)
	var category = GetOrCreateCategory(con.Db, jsonData.Category)
	var publisher = GetOrCreatePublisher(con.Db, jsonData.Publisher)
	var article = ArticleModel{
		Title: jsonData.Title, Body: jsonData.Body, PublishedAt: jsonData.PublishedAt, CreatedAt: jsonData.CreatedAt}
	if category.ID != 0 && publisher.ID != 0 {
		article.Category = category
		article.Publisher = publisher
		con.Db.Create(&article)
	} else if category.ID != 0 {
		article.Category = category
		con.Db.Create(&article).Related(&category)
	} else if publisher.ID != 0 {
		article.Publisher = publisher
		con.Db.Create(&article).Related(&publisher)
	} else {
		con.Db.Create(&article)
	}
	return &article, nil
}

//UpdateArticle updates or creats an article
func (con *Connector) UpdateArticle(data []byte) (*ArticleModel, error) {
	var jsonData Article
	json.Unmarshal(data, &jsonData)
	var category = GetOrCreateCategory(con.Db, jsonData.Category)
	var publisher = GetOrCreatePublisher(con.Db, jsonData.Publisher)
	var article = ArticleModel{
		Title: jsonData.Title, Body: jsonData.Body, PublishedAt: jsonData.PublishedAt, CreatedAt: jsonData.CreatedAt}
	if category.ID != 0 && publisher.ID != 0 {
		article.Category = category
		article.Publisher = publisher
	} else if category.ID != 0 {
		article.Category = category
	} else if publisher.ID != 0 {
		article.Publisher = publisher
	}
	var existingArticle = ArticleModel{}
	tx := con.Db.Begin()
	tx.Where("title = ?", jsonData.Title).First(&existingArticle)
	if existingArticle.ID != 0 {
		article.ID = existingArticle.ID
		tx.Save(&article)
		tx.Commit()
	} else {
		con.Db.Create(&article)
	}
	return &article, nil
}

//GetArticles gets all articles from the database
func (con *Connector) GetArticles(category string, publisher string, createdAt string, publishedAt string) ([]*ArticleModel, error) {
	var articles = []*ArticleModel{}
	var categoryM = CategoryModel{}
	var publisherM = PublisherModel{}
	transactional := false
	con.Db.Where("title = ?", category).First(&categoryM)
	con.Db.Where("name = ?", publisher).First(&publisherM)
	var query = map[string]interface{}{}
	tx := con.Db.Begin()
	if categoryM.Title != "" {
		transactional = true
		query["category_id"] = categoryM.ID
	}
	if publisherM.Name != "" {
		transactional = true
		query["publisher_id"] = publisherM.ID
	}
	if createdAt != "" {
		transactional = true
		query["created_at"] = createdAt
	}
	if publishedAt != "" {
		transactional = true
		tx = tx.Where("published_at = ?", publishedAt)
	}

	if transactional {
		tx.Where(query)
		tx.Preload("Category").Preload("Publisher").Find(&articles)
	} else {
		con.Db.Preload("Category").Preload("Publisher").Find(&articles)
	}
	return articles, nil
}

//GetArticle gets a single utem from the database
func (con *Connector) GetArticle(id uint) (*ArticleModel, error) {
	var article = ArticleModel{}
	con.Db.Where("ID = ?", id).Preload("Category").Preload("Publisher").Find(&article)
	return &article, nil
}

//DeleteArticle deletes article from the database
func (con *Connector) DeleteArticle(id uint) {
	con.Db.Where("ID = ?", id).Delete(&ArticleModel{})
}

//Database object
var Database DatabaseInterface

//InitDb initializes database
func InitDb(d DatabaseInterface) {
	Database = d

}
