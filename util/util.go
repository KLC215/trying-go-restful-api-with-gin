package util

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

func GenShortId() (string, error) {
	return shortid.Generate()
}

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")

	if !ok {
		return ""
	}

	if reqID, ok := v.(string); ok {
		return reqID
	}

	return ""
}
