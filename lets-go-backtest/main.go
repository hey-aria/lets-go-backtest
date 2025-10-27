package main

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	go startGinServer()
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "lets-go-backtest",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}

func startGinServer() {
	// create web server
	r := gin.Default()

	// create /tmp dir for handling any temporary files
	if _, err := os.Stat("./tmp"); os.IsNotExist(err) {
		os.Mkdir("./tmp", 0700)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("data")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read file"})
			return
		}

		out, err := os.Create("./tmp/upload~" + header.Filename)
		fmt.Println(err)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to write temp file"})
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to write csv to temp file"})
			return
		}
		fmt.Println(err)

		c.Status(http.StatusOK)
	})

	r.Run(":8080")
}
