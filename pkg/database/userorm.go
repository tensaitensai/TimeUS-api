package database

import "github.com/tensaitensai/TimeUS-api/pkg/model"

func CreateUser(user *model.User) {
	db.Create(user)
}

func FindUser(u *model.User) model.User {
	var user model.User
	db.Where(u).First(&user)
	return user
}
