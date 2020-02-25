package main

import (
	"assessment1/models"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createArticle(res http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	models.Database.CreateArticle(reqBody)
	// defer reqBody.Close()
	// fmt.Fprintf(res, "Success")
	fmt.Fprintf(res, "%+v", string(reqBody))
}

func updateArticle(res http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	models.Database.UpdateArticle(reqBody)
	// defer reqBody.Close()
	// fmt.Fprintf(res, "Success")
	fmt.Fprintf(res, "%+v", string(reqBody))
}

func getArticle(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key, err1 := strconv.ParseUint(vars["id"], 10, 32)
	if err1 != nil {
		fmt.Fprintf(res, "Invalid key")
	}
	id := uint(key)
	var article, err2 = models.Database.GetArticle(id)
	if err2 != nil {
		fmt.Fprintf(res, err2.Error())

	}
	json.NewEncoder(res).Encode(article)

}

func deleteArticle(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key, err1 := strconv.ParseUint(vars["id"], 10, 32)
	if err1 != nil {
		fmt.Fprintf(res, "Invalid key")
	}
	id := uint(key)
	models.Database.DeleteArticle(id)
	fmt.Fprintf(res, "Deleted article with id %d", id)

}

func getArticles(res http.ResponseWriter, req *http.Request) {
	var category string
	var publisher string
	var createdAt string
	var publishedAt string
	if key, ok := req.URL.Query()["category"]; ok {
		category = key[0]
	}
	if key, ok := req.URL.Query()["publisher"]; ok {
		publisher = key[0]
	}
	if key, ok := req.URL.Query()["created_at"]; ok {
		createdAt = key[0]
	}
	if key, ok := req.URL.Query()["published_at"]; ok {
		publishedAt = key[0]
	}

	var articles, err = models.Database.GetArticles(category, publisher, createdAt, publishedAt)
	if err != nil {
		fmt.Fprintf(res, err.Error())
	}
	json.NewEncoder(res).Encode(articles)

}

func main() {
	var host, user, dbname, password string
	var port uint
	flag.StringVar(&host, "host", "localhost", "Host name")
	flag.StringVar(&user, "user", "assess1", "Database user's name")
	flag.StringVar(&dbname, "dbname", "assessment", "Database base name")
	flag.StringVar(&password, "password", "1234test", "Database password")
	flag.UintVar(&port, "port", 8010, "Server port")
	flag.Parse()
	dbString := fmt.Sprintf("host=%s port%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	db, err := gorm.Open("postgres", dbString)
	if err != nil {
		fmt.Println("Database error", err.Error())
	}
	defer db.Close()
	models.InitDb(&models.Connector{Db: db})
	initializeDb(db)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/article", createArticle).Methods("POST")
	router.HandleFunc("/article", updateArticle).Methods("PUT")
	router.HandleFunc("/articles", getArticles).Methods("GET")
	router.HandleFunc("/article/{id}", getArticle).Methods("GET")
	router.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	fmt.Printf("Starting server at port :%d\n", port)
	log.Fatal(http.ListenAndServe(":8011", router))
}

func initializeDb(db *gorm.DB) {
	// db.Debug().DropTableIfExists(&models.ArticleModel{})
	// db.Debug().DropTableIfExists(&models.PublisherModel{})
	// db.Debug().DropTableIfExists(&models.CategoryModel{})
	db.Debug().AutoMigrate(&models.ArticleModel{})
	db.Debug().AutoMigrate(&models.PublisherModel{})
	db.Debug().AutoMigrate(&models.CategoryModel{})

}
