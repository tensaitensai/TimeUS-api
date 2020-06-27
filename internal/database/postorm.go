package database

import (
	"fmt"

	"github.com/tensaitensai/TimeUS-api/internal/model"
)

func CreatePost(post *model.Post) {
	db.Create(post)
}

func FindListPosts(p *model.Post) (model.Posts, error) {
	var posts model.Posts
	db.Where(p).Find(&posts)
	if len(posts) == 0 {
		return posts, fmt.Errorf("Could not find Posts (%v)", p)
	}
	return posts, nil
}

func FindGetPost(p *model.Post) (*model.Post, error) {
	var post model.Post
	if err := db.Where(p).First(&post).Error; err != nil {
		return &post, fmt.Errorf("Could not find Post (%v)", p)
	}
	return &post, nil
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
