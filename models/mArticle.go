package models

import "time"

// Article represents the model for an article
// Default table name will be `articles`
type Article struct {
	// gorm.Model
	ArticleID uint      `redis:"articleId" form:"articleId" json:"articleId" gorm:"primary_key"`
	Author    string    `redis:"author" form:"author" json:"author" gorm:"type:text"`
	Title     string    `redis:"title" form:"title" json:"title" gorm:"type:text"`
	Body      string    `redis:"body" form:"body" json:"body" gorm:"type:text"`
	Created   time.Time `redis:"created" form:"created" json:"created" gorm:"type:timestamp"`
}
