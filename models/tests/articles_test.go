package models

import (
	"assessment1/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func load() *models.Articles {
	fl, err := os.Open("../data/articles.json")
	defer fl.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	var articles models.Articles
	readContent, err := ioutil.ReadAll(fl)
	if err != nil {
		fmt.Printf(err.Error())
	}
	json.Unmarshal(readContent, &articles.Articles)
	return &articles
}

func TestCreateArticles(t *testing.T) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=assess1 dbname=assessment_test password=test1234 sslmode=disable")
	if err != nil {
		fmt.Println("Database error", err.Error())
	}
	defer db.Close()
	db.Delete(&models.ArticleModel{})
	var article = models.ArticleModel{Title: "Title"}
	db.Create(&article)

	var articles = []models.ArticleModel{}
	db.Find(&articles)
	var count = len(articles)
	if !reflect.DeepEqual(1, len(articles)) {
		t.Errorf("Expected: %#v\nActual   %#v\n", 1, count)
	}
	db.Rollback()
}

func TestArticles(t *testing.T) {
	expected := []string{"1 Lorem ipsum dolor sit amet", "2 Lorem ipsum dolor sit amet", "3 Lorem ipsum dolor sit amet"}
	articles := load()
	actual := []string{}
	for _, article := range articles.Articles {
		actual = append(actual, article.Title)
	}
	if len(articles.Articles) != 3 {
		panic(len(articles.Articles))
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %#v\nActual:   %#v\n", expected, actual)
	}

}
