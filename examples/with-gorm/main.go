package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ls6-events/gengo"
	gengoGin "github.com/ls6-events/gengo/inputs/gin"
	"github.com/ls6-events/gengo/outputs/openapi"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"withgorm/types"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&types.Post{}, &types.Author{}, &types.Comment{})
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/posts", GetPosts)
	r.GET("/posts/:id", GetPost)
	r.POST("/posts", CreatePost)
	r.PUT("/posts/:id", UpdatePost)
	r.DELETE("/posts/:id", DeletePost)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	gen := gengo.New(gengoGin.WithGinInput(r), openapi.WithOpenAPIOutput("openapi.generated.yaml"))

	config := gengo.Config{
		Title:   "Example API",
		Version: "1.0.0",
		Host:    "localhost",
		Port:    8000,
	}

	gen.SetConfig(&config)

	err = gen.Parse()
	if err != nil {
		panic(err)
	}

	err = r.Run(":8000")
	if err != nil {
		panic(err)
	}
}
