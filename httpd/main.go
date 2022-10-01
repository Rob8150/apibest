package main

import (
	"github.com/Rob8150/apibest/httpd/handler"
	"github.com/Rob8150/apibest/platform/newsfeed"
	"github.com/gin-gonic/gin"
)

//cd apibest
//go test ./platform/newsfeed
//{"hello":"Found Me"}
// curl http://localhost:8080/newsfeed
// curl -H "Content-Type: application/json" -d '{"Title":"Robert","Post":"Whats up man"}' http://localhost:8080/newsfeed

//go get -u github.com/gin-gonic/gin
//go run main.go
//http://localhost:8080/ping

func main() {

	feed := newsfeed.New()

	r := gin.Default()

	r.GET("/ping", handler.PingGet())

	r.GET("/newsfeed", handler.NewsfeedGet(feed))

	r.POST("/newsfeed", handler.NewsfeedPost(feed))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	//feed := newsfeed.New()

	//fmt.Println(feed)

	//feed.Add(newsfeed.Item{"Hello", "How ya' doing mate?"})

	//fmt.Println(feed)

}
