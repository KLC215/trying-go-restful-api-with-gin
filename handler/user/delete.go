package user

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/package/errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errors.DatabaseError, nil)
		return
	}

	SendResponse(c, nil, nil)
}
