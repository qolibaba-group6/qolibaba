
package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.Default()

	// Apply Metrics Middleware
	r.Use(MetricsMiddleware())

	// Expose metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Add other routes here
	log.Println("User Panel Service is running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
