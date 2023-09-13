package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Resource struct {
	Material string
	Text     string
	Period   string
	Place    string
	ImageURL string
}

var materials = []Resource{
	{"Титан", "2 кг", "01.01.2023 - 01.06.2023", "Море Восточное", "/image/titanium.png"},
	{"Алюминий", "11 кг", "12.02.2023 - 25.05.2023", "Океан Бурь", "/image/aluminium.jpg"},
	{"Железо", "6 кг", "15.01.2023 - 07.04.2023", "Море Влажности", "/image/ferrum.jpg"},
}

func StartServer() {
	log.Println("Server start up")

	r := gin.Default()
	// r - сущность (структура) типа Engine* с встроенным логгером* и Recovery middleware*
	// Engine - сущность фреймворка с muxer'ом (это мультиплексор HTTP запросов),
	//								  конфигурацией и миддлварой (слой обработки ошибок)

	// c *gin.Context отвечает за передачу данных между миддлварами
	//							  проверку того, что json приходит в нужном формате
	//							  рендер json ответа
	// H - сокращение от map[string]any

	r.GET("/home", func(c *gin.Context) {
		queryParam, ok := c.GetQuery("search")

		if queryParam == "титан" && ok {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Ресурсы:",
				"res1":  "Титан",
				"res2":  "",
				"res3":  "",
				"btn":   1,
			})
		} else if queryParam == "алюминий" && ok {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Ресурсы:",
				"res1":  "",
				"res2":  "Алюминий",
				"res3":  "",
				"btn":   1,
			})
		} else if queryParam == "железо" && ok {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Ресурсы:",
				"res1":  "",
				"res2":  "",
				"res3":  "Железо",
				"btn":   1,
			})
		} else {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Ресурсы:",
				"res1":  "Титан",
				"res2":  "Алюминий",
				"res3":  "Железо",
				"btn":   0,
			})
		}

	})

	r.GET("/home/titanium", func(c *gin.Context) {
		c.HTML(http.StatusOK, "material.html", gin.H{
			"Material": "Титан",
			"Text":     "2 кг	",
			"Period":   "01.01.2023 - 01.06.2023",
			"Place":    "Море Восточное",
			"ImageURL": "/image/titanium.png", // URL-адрес изображения для Титана
		})
	})

	r.GET("/home/aluminium", func(c *gin.Context) {
		c.HTML(http.StatusOK, "material.html", gin.H{
			"Material": "Алюминий",
			"Text":     "11 кг",
			"Period":   "12.02.2023 - 25.05.2023",
			"Place":    "Океан Бурь",
			"ImageURL": "/image/aluminium.jpg", // URL-адрес изображения для Алюминия
		})
	})

	r.GET("/home/ferrum", func(c *gin.Context) {
		c.HTML(http.StatusOK, "material.html", gin.H{
			"Material": "Железо",
			"Text":     "6 кг",
			"Period":   "15.01.2023 - 07.04.2023",
			"Place":    "Море Влажности",
			"ImageURL": "/image/ferrum.jpg", // URL-адрес изображения для Железа
		})
	})

	r.LoadHTMLGlob("templates/*") // подгружаем html файлы по паттерну из templates

	r.Static("/image", "./resources") // это нужно чтобы картинки грузились ?

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}
