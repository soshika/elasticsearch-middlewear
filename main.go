package main

import (
	"github.com/gin-gonic/gin"
	"github.com/soshika/elasticsearch-middlewear/app"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	app.StartApplication()
}
