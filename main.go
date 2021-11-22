package main

import (
	"github.com/elianvancutsem/eliancodes-image-generator/api"
)

func main() {
	api.StartApi()
	/*err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://59a9d679a16448a0888bb626e7dcc957@o1030206.ingest.sentry.io/6035002",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
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
		var newImage generateImageFromAPI

		if err := c.BindJSON(&newImage); err != nil {
			sentry.CaptureException(err)
			c.String(http.StatusBadRequest, "Something went wrong while building the image")
			return
		}

		var defaultSocialImageSize size = size{width: 2024, height: 1012}
		var randomColor color.Color = getRandomColor().color
		generateImage(socialImage{name: "defaultBanner", size: defaultSocialImageSize, baseColor: randomColor, title: Line{content: newImage.Title, color: randomColor, size: 32, font: "Medium"}})
		c.IndentedJSON(http.StatusCreated, newImage)
	})

	// Run application
	app.Run(":3000")*/
}
