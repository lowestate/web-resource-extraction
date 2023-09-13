package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Resource struct {
	Material    string
	Text        string
	Description string
	ImageURL    string
}

func StartServer() {
	log.Println("Server start up")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/home/titanium", func(c *gin.Context) {
		c.HTML(http.StatusOK, "material.html", gin.H{
			"Material": "Титан",
			"Text":     "Добыто 2 кг титана",
			"ImageURL": "/image/titanium.png", // URL-адрес изображения для Титана
		})
	})

	r.GET("/home/aluminium", func(c *gin.Context) {
		c.HTML(http.StatusOK, "material.html", gin.H{
			"Material": "Алюминий",
			"Text":     "Добыто 11 кг алюминия",
			"ImageURL": "/image/aluminium.jpg", // URL-адрес изображения для Алюминия
		})
	})

	r.GET("/home/ferrum", func(c *gin.Context) {
		c.HTML(http.StatusOK, "material.html", gin.H{
			"Material": "Железо",
			"Text":     "Добыто 6 кг железа",
			"ImageURL": "/image/ferrum.jpg", // URL-адрес изображения для Железа
		})
	})

	r.LoadHTMLGlob("templates/*")

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Ресурсы:",
			"res1":  "Титан",
			"res2":  "Алюминий",
			"res3":  "Железо",
		})
	})

	r.Static("/image", "./resources")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}
