package user

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/package/errors"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	username := c.Param("username")

	user, err := model.GetUser(username)

	if err != nil {
		SendResponse(c, errors.UserNotFoundError, nil)
		return
	}

	SendResponse(c, nil, user)
}
