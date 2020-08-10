package main

import (
	"gin-mvc/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.Serve(r)

	r.Run()
}
