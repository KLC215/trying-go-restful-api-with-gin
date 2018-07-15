package user

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/package/auth"
	"apiserver/package/errors"
	"apiserver/package/token"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var u model.UserModel

	// Bind request data
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errors.BindError, nil)
		return
	}

	// Get username from database
	d, err := model.GetUser(u.Username)
	if err != nil {
		SendResponse(c, errors.UserNotFoundError, nil)
		return
	}

	// Compare user passwords
	if err := auth.ComparePassword(d.Password, u.Password); err != nil {
		SendResponse(c, errors.PasswordIncorrectError, nil)
		return
	}

	// Generate JWT
	t, err := token.Sign(c, token.Context{ID: d.Id, Username: d.Username}, "")
	if err != nil {
		SendResponse(c, errors.TokenError, nil)
		return
	}

	SendResponse(c, nil, model.Token{Token: t})
}
