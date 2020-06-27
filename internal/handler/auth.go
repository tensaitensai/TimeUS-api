package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tensaitensai/TimeUS-api/internal/database"
	"github.com/tensaitensai/TimeUS-api/internal/model"
	"golang.org/x/crypto/bcrypt"
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
		return APIResponseError(c, http.StatusBadRequest, "invalid email or password", nil)
	}

	if len(user.Password) > 32 || len(user.Password) < 8 {
		return APIResponseError(c, http.StatusBadRequest, "password is long or short", nil)
	}

	if u := database.FindUser(&model.User{Email: user.Email}); u.ID != 0 {
		return APIResponseError(c, http.StatusConflict, "email already exists", nil)
	}

	hash, err := hashpassword(user.Password)
	if err != nil {
		return APIResponseError(c, http.StatusBadRequest, "couldn't hash password", err)
	}

	user.Password = hash
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
	if user.ID == 0 {
		return APIResponseError(c, http.StatusUnauthorized, "invalid email", nil)
	}
	if err := comparepassword(user.Password, u.Password); err != nil {
		return APIResponseError(c, http.StatusUnauthorized, "invalid password", err)
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

func hashpassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func comparepassword(hash, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
}
