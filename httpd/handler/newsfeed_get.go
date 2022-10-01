package handler

import (
	"net/http"

	"github.com/Rob8150/apibest/platform/newsfeed"
	"github.com/gin-gonic/gin"
)

//NewsfeedGet gin function
func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
