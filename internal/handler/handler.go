package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tensaitensai/TimeUS-api/internal/database"
	"github.com/tensaitensai/TimeUS-api/internal/model"
)

/*
func GetPost(c echo.Context) error {
}

func GetRanking(c echo.Context) error {
}
*/

func AddPost(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := database.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return APIResponseError(c, http.StatusBadRequest, "Bad Request")
	}

	post := new(model.Post)
	if err := c.Bind(post); err != nil {
		return err
	}

	if pUID, err := strconv.Atoi(c.Param("uid")); err != nil || pUID != uid {
		return APIResponseError(c, http.StatusBadRequest, "invalid url (Userid is not int)")
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

	if pUID, err := strconv.Atoi(c.Param("uid")); err != nil || pUID != uid {
		return APIResponseError(c, http.StatusBadRequest, "invalid url (Userid is not int)")
	}
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return APIResponseError(c, http.StatusBadRequest, "invalid url (Postid is not int)")
	}

	if err := database.DeletePost(&model.Post{ID: pID, UserID: uid}); err != nil {
		return APIResponseError(c, http.StatusBadRequest, "") //治す
	}

	return c.NoContent(http.StatusNoContent)
}

func UpdatePost(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := database.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return APIResponseError(c, http.StatusBadRequest, "Bad Request")
	}

	p := new(model.Post)
	if err := c.Bind(p); err != nil {
		return err
	}

	if pUID, err := strconv.Atoi(c.Param("uid")); err != nil || pUID != uid {
		return APIResponseError(c, http.StatusBadRequest, "invalid url (Userid is not int)")
	}
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return APIResponseError(c, http.StatusBadRequest, "invalid url (Postid is not int)")
	}
	if p.Subjectname == "" {
		return APIResponseError(c, http.StatusBadRequest, "invalid Subjectname of post")
	}

	post := database.FindGetPost(&model.Post{ID: pID, UserID: uid})
	if post.Subjectname == "" {
		return APIResponseError(c, http.StatusUnauthorized, "invalid post")
	}

	post.Subjectname = p.Subjectname
	post.Subjectstarttime = p.Subjectstarttime
	post.Subjectendtime = p.Subjectendtime

	if err := database.UpdatePost(post); err != nil {
		return APIResponseError(c, http.StatusBadRequest, "") //治す
	}

	return c.NoContent(http.StatusNoContent)
}
