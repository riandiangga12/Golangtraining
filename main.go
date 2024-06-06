package main

import (
	"log"

	"training/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi router Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Definisikan route
	router.SetupRouter(r)
	//fmt.Println("Hello")
	// Jalankan server pada port 8080
	log.Println("Running server on port 8080")
	r.Run(":8080")
}
