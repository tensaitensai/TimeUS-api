package database

import (
	"fmt"

	"github.com/tensaitensai/TimeUS-api/pkg/model"
)

func CreatePost(post *model.Post) {
	db.Create(post)
}

func FindPosts(p *model.Post) model.Posts {
	var posts model.Posts
	db.Where(p).Find(&posts)
	return posts
}

func DeletePost(p *model.Post) error {
	if rows := db.Where(p).Delete(&model.Post{}).RowsAffected; rows == 0 {
		return fmt.Errorf("Could not find Post (%v) to delete", p)
	}
	return nil
}

func UpdatePost(p *model.Post) error {
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
