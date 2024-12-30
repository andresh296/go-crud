package main

import (
	"github.com/andresh296/go-crud/cmd/api"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	os.Setenv("PORT", "8081")
	app := gin.Default()
	api.Boostrap(app)

	app.Run()
}