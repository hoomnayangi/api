package model

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo"
)

// User data model
type User struct {
	Base
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`

	AccessToken *string `json:"access_token,omitempty" sql:"-"`
}

const (
	uidKey string = "uid"
)

// SetUserIDToCtx - set a uid to echo context for later user's session usage
func SetUserIDToCtx(c echo.Context, uid int) {
	c.Set(uidKey, uid)
}

// GetUserIDFromCtx - get uid from echo context
func GetUserIDFromCtx(c echo.Context) int {
	id, _ := strconv.Atoi(fmt.Sprintf("%v", c.Get(uidKey)))
	return id
}
