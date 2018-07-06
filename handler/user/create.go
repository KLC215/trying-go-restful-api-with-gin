package user

import (
	. "apiserver/handler"
	"apiserver/package/errors"
	"apiserver/util"
	"fmt"

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

	if err := r.checkParams(); err != nil {
		SendResponse(c, err, nil)
		return
	}

	username := c.Param("username")
	log.Infof("URL username: %s", username)

	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)

	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)

	if r.Username == "" {
		SendResponse(c,
			errors.New(errors.UserNotFoundError,
				fmt.Errorf("Username cannot found in db: xx.xx.xx.xx")),
			nil)

		return
	}

	if r.Password == "" {
		SendResponse(c, fmt.Errorf("Password is empty"), nil)

		return
	}

	res := CreateResponse{
		Username: r.Username,
	}

	SendResponse(c, nil, res)
}
