package models

import (
	"assessment1/models"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func setup() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=assess1 dbname=assessment_test password=test1234 sslmode=disable")
	if err != nil {
		fmt.Println("Database error", err.Error())
	}
	models.InitDb(&models.Connector{Db: db})
	initializeDb(db)
	return db

}

func getArticle() *models.ArticleModel {
	return &models.ArticleModel{Title: "Lorem ipsum dolor sit amet", Body: "Body"}
}

func TestBaseCreateArticles(t *testing.T) {
	var db = setup()
	defer db.Close()

	fl, err := os.Open("models/data/article.json")
	defer fl.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	readContent, err := ioutil.ReadAll(fl)
	if err != nil {
		fmt.Printf(err.Error())
	}

	// var createdArticle *models.ArticleModel
	if createdArticle, err := models.Database.CreateArticle(readContent); err == nil {
		var article models.ArticleModel
		article.Publisher = models.PublisherModel{ID: 1}
		article.Category = models.CategoryModel{ID: 1}
		db.First(&article)
		defer db.Rollback()
		if !reflect.DeepEqual(article, *createdArticle) {
			t.Errorf("Expected: %#v\nActual:   %#v\n", article, *createdArticle)
		}

	}

}

func TestBaseGetArticle(t *testing.T) {
	var db = setup()
	defer db.Close()
	defer db.Rollback()
	var article = getArticle()
	article.Publisher = models.PublisherModel{}
	article.Category = models.CategoryModel{}
	var expectingArticle = &models.ArticleModel{}
	db.Create(&article)
	expectingArticle, _ = models.Database.GetArticle(1)
	if !reflect.DeepEqual(expectingArticle, article) {
		t.Errorf("Expected: %#v\nActual:   %#v\n", expectingArticle, article)
	}

}

func TestBaseGetArticles(t *testing.T) {
	var db = setup()
	defer db.Close()
	defer db.Rollback()
	var article = getArticle()
	var articles = []*models.ArticleModel{}
	article.Publisher = models.PublisherModel{}
	article.Category = models.CategoryModel{}
	var expectingArticle = []*models.ArticleModel{}
	db.Create(&article)
	articles = append(articles, article)
	expectingArticle, _ = models.Database.GetArticles("", "", "", "")
	if !reflect.DeepEqual(expectingArticle, articles) {
		t.Errorf("Expected: %#v\nActual:   %#v\n", expectingArticle, articles)
	}

}

// func TestBaseUpdateArticle(t *testing.T) {
// 	var db = setup()
// 	defer db.Close()
// 	defer db.Rollback()
// 	var article = getArticle()
// 	article.Publisher = models.PublisherModel{}
// 	article.Category = models.CategoryModel{}
// 	var expectingArticle = &models.ArticleModel{}
// 	db.Create(&article)
// 	expectingArticle, _ = models.Database.UpdateArticle([]byte(""))
// 	if !reflect.DeepEqual(expectingArticle, article) {
// 		t.Errorf("Expected: %#v\nActual:   %#v\n", expectingArticle, article)
// 	}

// }

func initializeDb(db *gorm.DB) {
	db.Debug().DropTableIfExists(&models.ArticleModel{})
	db.Debug().DropTableIfExists(&models.PublisherModel{})
	db.Debug().DropTableIfExists(&models.CategoryModel{})
	db.Debug().AutoMigrate(&models.ArticleModel{})
	db.Debug().AutoMigrate(&models.PublisherModel{})
	db.Debug().AutoMigrate(&models.CategoryModel{})

}
