package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tensaitensai/TimeUS-api/pkg/database"
	"github.com/tensaitensai/TimeUS-api/pkg/model"
)

/*
func GetPostsToday(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := database.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return APIResponseError(c, http.StatusBadRequest, "Bad Request")
	}

	posts := database.FindPosts(&model.Post{UID: uid})
	return c.JSON(http.StatusOK, posts)
}*/

func AddPost(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := database.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return APIResponseError(c, http.StatusBadRequest, "Bad Request")
	}

	post := new(model.Post)
	if err := c.Bind(post); err != nil {
		return err
	}

	if post.Subjectname == "" {
		return APIResponseError(c, http.StatusBadRequest, "invalid Subjectname of post")
	}

	post.UserID = uid
	database.CreatePost(post)

	return c.JSON(http.StatusCreated, post)
}

func DeletePost(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := database.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return APIResponseError(c, http.StatusBadRequest, "Bad Request")
	}

	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return APIResponseError(c, http.StatusBadRequest, "Postid is not int")
	}

	if err := database.DeletePost(&model.Post{ID: postID, UserID: uid}); err != nil {
		return APIResponseError(c, http.StatusBadRequest, err) //治す
	}

	return c.NoContent(http.StatusNoContent)
}

func UpdatePost(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := database.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return APIResponseError(c, http.StatusBadRequest, "Bad Request")
	}

	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return APIResponseError(c, http.StatusBadRequest, "Postid is not int")
	}

	posts := database.FindPosts(&model.Post{ID: postID, UserID: uid})
	if len(posts) != 1 {
		return APIResponseError(c, http.StatusBadRequest, "Bad Request")
	}

	post := posts[0]
	if err := database.UpdatePost(&post); err != nil {
		return APIResponseError(c, http.StatusBadRequest, err) //治す
	}

	return c.NoContent(http.StatusNoContent)
}
