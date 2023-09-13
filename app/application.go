package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {
	router.Use(CORSMiddleware())
	URLPatterns()

	router.Run(":9090")
}
