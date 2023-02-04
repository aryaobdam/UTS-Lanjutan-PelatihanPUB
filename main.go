package main

import (
	"UTS-LANJUTAN-ARYA/connection"
	"UTS-LANJUTAN-ARYA/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := connection.ConnectDB()
	rh := &router.Handlers{
		DB: db,
		R:  r,
	}

	rh.Routes()

	r.Run(":8080")
}
