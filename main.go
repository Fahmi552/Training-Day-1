package main

import (
	"Hello/router"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi router Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Definisikan route
	router.SetupRouter(r)

	// Jalankan server pada port 8080
	log.Println("Running server on port 8080")
	r.Run(":8080")
	//fmt.Println("Hello World!")
	// inisialisasi router gin
	//r := gin.Default()

	//Definisikan roten gin
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hallo dari Gin!",
	// 	})
	// })

	// //POST
	// r.POST("/post", func(c *gin.Context) {
	// 	var json struct {
	// 		Message string `json:"message"`
	// 	}

	// 	if err := c.ShouldBindJSON(&json); err == nil {
	// 		c.JSON(200, gin.H{"message": json.Message})
	// 	} else {
	// 		c.JSON(400, gin.H{"error": err.Error()})
	// 	}
	// })

	// // Jalankan server pada port 8080
	// r.Run(":8080")

	// //router
	// gin.SetMode(gin.ReleaseMode)
	// //r := gin.Default()
	// router.SetupRouter(r)

	// //jalankan
	// log.Println("Running server on port 8080")
	// r.Run(":8080")

}
