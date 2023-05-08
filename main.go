package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.AbortWithError(400, err)
			return
		}
		defer file.Close()

		f, err := os.Create(header.Filename)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		c.String(200, "File uploaded successfully")

		r.GET("/download/:filename", func(c *gin.Context) {
			filename := c.Param("filename")
			 c.File(filename)
		})
	})

	r.Run(":8080")
}
