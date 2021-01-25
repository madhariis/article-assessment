package controllers

import (
	"../models"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type CtrlArticle struct{}

func NewCtrlArticle() *CtrlArticle {
	return &CtrlArticle{}
}

func (c *CtrlArticle) CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dataSourceName := "root:root@tcp(localhost:3306)/?parseTime=True"
	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}else {
		fmt.Println("success to connect database")
	}
	db.Exec("USE articles_db")

	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	//fmt.Println("reqBody", reqBody)
	var article models.Article
	_ = json.Unmarshal(reqBody, &article)
	db.Create(&article)
	db.Model(&article).Update("Created", time.Now())

	fmt.Println("Endpoint Hit: Creating New Article")
	_ = json.NewEncoder(w).Encode(article)
}

func (c *CtrlArticle) GetArticles(w http.ResponseWriter, r *http.Request) {
	w. Header().Set("Content-Type", "application/json")
	dataSourceName := "root:root@tcp(localhost:3306)/?parseTime=True"
	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}else {
		fmt.Println("success to connect database")
	}
	db.Exec("USE articles_db")

	var articles []models.Article
	//check is user sorting by Author?
	author := r.URL.Query()
	inputAuthor := author.Get("author")
	//inputAuthor, ok := r.URL.Query()["author"]
	if inputAuthor == "" {
		log.Println("not sorting by Author")

		//check is user search a keyword?
		query := r.URL.Query()
		inputQuery := query.Get("query")
		if inputQuery == "" {
			log.Println("doesn't search a keyword")
			//SEARCH ALL
			db.Order("created desc").Find(&articles)
		}else {
			log.Println("search a keyword: ", inputQuery)
			db.Where("title LIKE ? OR body LIKE ?", "%"+inputQuery+"%", "%"+inputQuery+"%").Find(&articles)
		}
	}else {
		query := r.URL.Query()
		inputQuery := query.Get("query")
		if inputQuery != "" {
			log.Println("sorting by author and keyword")
			db.Where("title LIKE ? OR body LIKE ? AND author = ?", "%"+inputQuery+"%", "%"+inputQuery+"%", inputAuthor).Find(&articles)
		}else {
			log.Println("sorting by Author: ", inputAuthor)
			db.Where("author = ?", inputAuthor).Find(&articles)
		}
	}

	_ = json.NewEncoder(w).Encode(articles)
}