package model

import "time"

//User の定義
type User struct {
	ID         int       `json:"id" gorm:"column:id;praimaly_key;NOT NULL;AUTO_INCREMENT"`
	Email      string    `json:"email" gorm;"column:email;type:varchar(255);UNIQUE;NOT NULL"`
	Password   string    `json:"password" gorm;"column:password;type:varchar(32);NOT NULL"`
	Createtime time.Time `json:"createtime" gorm:"column:createtime;type:DATETIME(6);NOT NULL"`
	Updatetime time.Time `json:"updatetime" gorm:"column:updatetime;type:DATETIME(6);NOT NULL"`
	Name       string    `json:"name" gorm:"column:name;type:varchar(16);NOT NULL"`
	Bio        string    `json:"bio" gorm:"column:bio;type:varchar(255)"`
	Myurl      string    `json:"myurl" gorm:"column:myurl;type:varchar(255)"`
}
