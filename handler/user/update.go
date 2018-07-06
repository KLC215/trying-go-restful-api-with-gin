package user

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/package/errors"
	"apiserver/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	userId, _ := strconv.Atoi(c.Param("id"))

	// Bind user data
	var u model.UserModel

	if err := c.Bind(&u); err != nil {
		SendResponse(c, errors.BindError, nil)
		return
	}

	u.Id = uint64(userId)

	// Validate user data
	if err := u.Validate(); err != nil {
		SendResponse(c, errors.ValidationError, nil)
		return
	}

	// Encrypt user password
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errors.EncryptError, nil)
		return
	}

	// Save changed fields
	if err := u.Update(); err != nil {
		SendResponse(c, errors.DatabaseError, nil)
		return
	}

	SendResponse(c, nil, nil)

}
