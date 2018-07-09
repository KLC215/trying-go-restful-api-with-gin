package user

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/package/errors"
	"apiserver/util"

	"github.com/lexkong/log/lager"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Create(c *gin.Context) {

	log.Info("[User]->Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	var r CreateRequest

	if err := c.Bind(&r); err != nil {
		SendResponse(c, errors.BindError, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// Validate data
	if err := u.Validate(); err != nil {
		SendResponse(c, errors.ValidationError, nil)
		return
	}

	// Encrypt password
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errors.EncryptError, nil)
		return
	}

	// Insert user to database
	if err := u.Create(); err != nil {
		SendResponse(c, errors.DatabaseError, nil)
		return
	}

	res := CreateResponse{
		Username: r.Username,
	}

	// Show user details
	SendResponse(c, nil, res)
}
