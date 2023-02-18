package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/api/cat/add", AddCat)

	r.GET("/api/cats", GetAllCats)

	r.GET("/api/cat/:id", GetCat)

	r.DELETE("/api/cat/:id", DeleteCat)

	r.PUT("/api/cat/:id", EditCat)

	r.POST("/api/dog/add", AddDog)

	r.GET("/api/dogs", GetAllDogs)

	r.GET("/api/dog/:id", GetDog)

	r.DELETE("/api/dog/:id", DeleteDog)

	r.PUT("/api/dog/:id", EditDog)

	r.Run("0.0.0.0:8888")
}
