package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tensaitensai/TimeUS-api/internal/database"
	"github.com/tensaitensai/TimeUS-api/internal/model"
)

type jwtCustomClaims struct {
	UID   int    `json:"uid"`
	Email string `json:"email"`
	jwt.StandardClaims
}

var signingKey = []byte("secret")

var Config = middleware.JWTConfig{
	Claims:     &jwtCustomClaims{},
	SigningKey: signingKey,
}

func Signup(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	if user.Email == "" || user.Password == "" {
		return APIResponseError(c, http.StatusBadRequest, "invalid email or password")
	}

	if u := database.FindUser(&model.User{Email: user.Email}); u.ID != 0 {
		return APIResponseError(c, http.StatusConflict, "email already exists")
	}

	database.CreateUser(user)
	user.Password = ""

	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	user := database.FindUser(&model.User{Email: u.Email})
	if user.ID == 0 || user.Password != u.Password {
		return APIResponseError(c, http.StatusUnauthorized, "invalid email or password")
	}

	claims := &jwtCustomClaims{
		user.ID,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func userIDFromToken(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	uid := claims.UID
	return uid
}
