package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"

	"./controllers"
	"./models"
)

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "root:root@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	//db.Exec("CREATE DATABASE articles_db")
	db.Exec("USE articles_db")

	// Migration to create tables for Article and Item schema
	db.AutoMigrate(models.Article{})
}

func main() {
	ctrlArticle := controllers.NewCtrlArticle()
	r := mux.NewRouter().StrictSlash(true)
	// Create New Article
	r.HandleFunc("/articles", ctrlArticle.CreateArticle).Methods("POST")
	// Read an Article filtered by author name
	//r.HandleFunc("/articles/{author}", ctrlArticle.GetArticleByAuthor).Methods("GET")
	// Read-all Articles
	r.HandleFunc("/articles", ctrlArticle.GetArticles).Methods("GET")
	// Initialize db connection
	initDB()

	fmt.Println("Listen and serve at port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
