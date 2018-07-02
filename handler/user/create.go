package user

import (
	"apiserver/package/errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Create(c *gin.Context) {
	var r struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var err error

	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": errors.BindError})
		return
	}

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)

	if r.Username == "" {
		err = errors.New(errors.UserNotFoundError,
			fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
		log.Errorf(err, "Get an error")
	}

	if errors.IsUserNotFoundError(err) {
		log.Debug("Error type is UserNotFoundError")
	}

	if r.Password == "" {
		err = fmt.Errorf("Password is empty")
	}

	code, message := errors.DecodeError(err)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
