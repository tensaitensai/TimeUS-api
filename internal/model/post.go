package model

import (
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
