package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main()  {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	// Handle Index
	router.GET("/", showIndexPage)
	//Handle GET requests at /article/view/some_article_id
	router.GET("/article/view/:article_id", getArticle)

	//router.GET("/", func(c *gin.Context) {
	//	c.HTML(
	//		http.StatusOK,
	//		"index.html",
	//		gin.H{
	//			"title": "Home Page",
	//		},
	//		)
	//})

	router.Run(":8089")
}