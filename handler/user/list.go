package user

import (
	. "apiserver/handler"
	"apiserver/package/errors"
	"apiserver/service"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var r ListRequest

	if err := c.Bind(&r); err != nil {
		SendResponse(c, errors.BindError, nil)
		return
	}

	users, count, err := service.ListUser(r.Username, r.Offset, r.Limit)

	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   users,
	})
}
