package handler

import (
	"net/http"

	"github.com/Rob8150/apibest/platform/newsfeed"
	"github.com/gin-gonic/gin"
)

//NewsfeedPostRequest model
type NewsfeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

//NewsfeedPost gin function
func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := NewsfeedPostRequest{}
		c.Bind(&requestBody)
		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		feed.Add(item)

		c.Status(http.StatusNoContent)
	}
}
