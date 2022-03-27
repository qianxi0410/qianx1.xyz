package main

import (
	"api/handlers"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	defaultPort = "8080"
)

var port string

func init() {
	flag.StringVar(&port, "p", "", "server port that runs on")
}

func main() {
	flag.Parse()
	e := gin.Default()
	e.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	blog := e.Group("/blog")
	{
		var postHandler handlers.PostHandler
		blog.GET("/post/:id", postHandler.Post())
		blog.GET("post/share/:display_id", postHandler.PostWithDisplayId())
		blog.GET("/posts/:page/:size", postHandler.Posts())
		blog.GET("/posts/count", postHandler.Count())
	}

	if port == "" {
		port = defaultPort
	}

	if err := e.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
