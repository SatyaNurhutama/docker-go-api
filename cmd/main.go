package main

import (
	"docker-go-api/pkg/books"
	"docker-go-api/pkg/common/config"
	"docker-go-api/pkg/common/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Error at config", err)
	}

	h := db.Init(&c)
	r := gin.Default()

	books.RegisterRoutes(r, h)

	r.Run(c.Port)
}
