package model

import (
	"fmt"
	"time"
)

//Post の定義
type Post struct {
	ID               int       `json:"id" gorm:"column:id;UNIQUE;NOT NULL;AUTO_INCREMENT"`
	UserID           int       `json:"userid" gorm:"column:userid;praimaly_key;NOT NULL;AUTO_INCREMENT"`
	Subjectname      string    `json:"subjectname" gorm:"column:subjectname;type:varchar(45);NOT NULL"`
	Subjectstarttime time.Time `json:"subjectstarttime" gorm:"column:subjectstarttime;type:DATETIME"`
	Subjectendtime   time.Time `json:"subjectendtime" gorm:"column:subjectendtime;type:DATETIME"`
}

type Posts []Post

func CreatePost(post *Post) {
	db.Create(post)
}

func FindPosts(p *Post) Posts {
	var posts Posts
	db.Where(p).Find(&posts)
	return posts
}

func DeletePosts(p *Post) error {
	if rows := db.Where(p).Delete(&Post{}).RowsAffected; rows == 0 {
		return fmt.Errorf("Could not find Post (%v) to delete", p)
	}
	return nil
}

func UpdatePosts(p *Post) error {
	rows := db.Model(p).Update(map[string]interface{}{
		"subjectname":      p.Subjectname,
		"subjectstarttime": p.Subjectstarttime,
		"subjectendtime":   p.Subjectendtime,
	}).RowsAffected
	if rows == 0 {
		return fmt.Errorf("Could not find Post (%v) to update", p)
	}
	return nil
}
