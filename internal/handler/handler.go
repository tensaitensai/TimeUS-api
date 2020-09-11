package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tensaitensai/TimeUS-api/internal/database"
	"github.com/tensaitensai/TimeUS-api/internal/model"
)

func GetPost(c echo.Context) error {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return APIResponseErrorLog(c, http.StatusBadRequest, "invalid url (userid is not int)", err)
	}

	item, err := strconv.Atoi(c.QueryParam("item"))
	if err != nil {
		return APIResponseErrorLog(c, http.StatusBadRequest, "invalid url (item is not int)", err)
	}

	layout := "2006-01-02"
	startdate, err := time.Parse(layout, c.QueryParam("startdate"))
	if err != nil {
		return APIResponseErrorLog(c, http.StatusBadRequest, "invalid url (startdate is not int)", err)
	}
	enddate, err := time.Parse(layout, c.QueryParam("enddate"))
	if err != nil {
		return APIResponseErrorLog(c, http.StatusBadRequest, "invalid url (enddate is not int)", err)
	}

	posts, err := database.FindPostsOnConditions(&model.Post{UserID: uid}, item, startdate, enddate)
	if err != nil {
		return APIResponseErrorLog(c, http.StatusBadRequest, "Could not find Posts", err)
	}

	return c.JSON(http.StatusOK, posts)
}

/*
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

	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return APIResponseErrorLog(c, http.StatusBadRequest, "invalid url (Postid is not int)", err)
	}

	if err := database.DeletePost(&model.Post{ID: pID, UserID: uid}); err != nil {
		return APIResponseErrorLog(c, http.StatusBadRequest, "Could not find Post to delete", err)
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

	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return APIResponseErrorLog(c, http.StatusBadRequest, "invalid url (Postid is not int)", err)
	}
	if p.Subjectname == "" {
		return APIResponseError(c, http.StatusBadRequest, "invalid Subjectname of post")
	}

	p.ID = pID
	p.UserID = uid
	if err := database.UpdatePost(p); err != nil {
		return APIResponseErrorLog(c, http.StatusBadRequest, "Could not find Post to update", err)
	}

	return c.NoContent(http.StatusNoContent)
}
