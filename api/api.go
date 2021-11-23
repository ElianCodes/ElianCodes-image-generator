package api

import (
	"fmt"
	imagegenerator "github.com/elianvancutsem/eliancodes-image-generator/api/imagegenerate"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"os"
)

func StartApi() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://59a9d679a16448a0888bb626e7dcc957@o1030206.ingest.sentry.io/6035002",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	} else {
		fmt.Println("Sentry initialized")
	}

	// create the app
	app := gin.Default()

	// Initialise middleware
	app.Use(sentrygin.New(sentrygin.Options{}))

	// Set up routes
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	// Set up routes
	app.GET("/health", func(c *gin.Context) {
		c.String(200, "Health seems fine")
	})

	app.POST("/generate", func(c *gin.Context) {
		var newImage imagegenerator.GenerateImageFromAPI

		if err := c.BindJSON(&newImage); err != nil {
			sentry.CaptureException(err)
			c.String(http.StatusBadRequest, "Something went wrong while building the image")
			return
		}

		var defaultSocialImageSize imagegenerator.Size = imagegenerator.Size{Width: 2024, Height: 1012}
		var randomColor color.Color = imagegenerator.GetRandomColor().Color
		var finalImage imagegenerator.SocialImage = imagegenerator.GenerateImage(imagegenerator.SocialImage{Name: "defaultBanner", Size: defaultSocialImageSize, BaseColor: randomColor, Title: imagegenerator.Line{Content: newImage.Title, Color: randomColor, Size: 32, Font: "Medium"}})
		c.Writer.Header().Set("Content-Type", "image/png")
		f, _ := os.Create(finalImage.Name + ".png")
		png.Encode(f, finalImage.Src)
		http.ServeFile(c.Writer, c.Request, finalImage.Name+".png")
		os.Remove(finalImage.Name + ".png")
	})

	// Run application
	fmt.Println("Starting API")
	app.Run(":3000")
}
