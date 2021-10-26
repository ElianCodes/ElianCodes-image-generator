package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
)

type size struct {
	width  int
	height int
}

type randomColor struct {
	name  string
	color color.RGBA
}

type Line struct {
	content string
	color   color.Color
	size    float64
	font    string
}

type socialImage struct {
	size        size
	name        string
	baseColor   color.Color
	src         image.Image
	link        Line
	title       Line
	subtitle    Line
	generalText Line
}

func main() {
	err := sentry.Init(sentry.ClientOptions{
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
			c.JSON(400, gin.H{
				"error": "Something went wrong while building the image",
			})
			return
		}

		var defaultSocialImageSize size = size{width: 2024, height: 1012}
		var randomColor color.Color = getRandomColor().color
		generateImage(socialImage{name: "defaultBanner", size: defaultSocialImageSize, baseColor: randomColor, title: Line{content: newImage.Title, color: randomColor, size: 32, font: "Medium"}})
		c.IndentedJSON(http.StatusCreated, newImage)
	})

	// Run application
	app.Run(":3000")
}

type generateImageFromAPI struct {
	Title string `json:"title"`
}

func getRandomColor() randomColor {
	availableColors := [5]randomColor{
		{name: "Orange", color: color.RGBA{253, 186, 116, 0xff}},
		{name: "Green", color: color.RGBA{134, 239, 172, 0xff}},
		{name: "Blue", color: color.RGBA{147, 197, 253, 0xff}},
		{name: "Red", color: color.RGBA{252, 165, 165, 0xff}},
		{name: "Purple", color: color.RGBA{240, 171, 252, 0xff}},
	}
	rand.Seed(time.Now().UnixNano())
	var pickedColor randomColor = availableColors[rand.Intn(len(availableColors))]
	fmt.Println("The color is picked, it is " + pickedColor.name + "!")
	return pickedColor
}

func generateImage(generation socialImage) {
	fmt.Println("starting generation of " + generation.name)

	// initialize a new (empty) image
	img := image.NewRGBA(image.Rect(0, 0, generation.size.width, generation.size.height))
	generation.src = img
	uniformColor := image.NewUniform(generation.baseColor)

	// Draw the background
	draw.Draw(img, image.Rect(generation.size.width/2, 0, generation.size.width, generation.size.height), uniformColor, image.Point{}, draw.Src)
	draw.Draw(img, image.Rect(0, 0, generation.size.width/2, generation.size.height), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// Set the ElianCodes image in place
	readHeroImg, err := os.Open("./assets/hero.png")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer readHeroImg.Close()
	heroImg, err := png.Decode(readHeroImg)
	heroImg = resize.Resize(600, 0, heroImg, resize.Lanczos3)

	rightMiddlePart := image.Pt((generation.size.width/2*-1)-200, -200)
	draw.Draw(img, img.Bounds(), heroImg, rightMiddlePart, draw.Over)

	// Set the text
	addText(generation, img)

	writeImage(generation)
}

func addText(generation socialImage, sourceImg draw.Image) {
	fmt.Println("starting generation of Text")
	ctx := freetype.NewContext()
	ctx.SetDPI(300)
	ctx.SetClip(generation.src.Bounds())
	ctx.SetDst(sourceImg)

	addLine(*ctx, generation.title)
}

func addLine(ctx freetype.Context, line Line) {
	var baseColor image.Image = image.NewUniform(line.color)
	ctx.SetFontSize(float64(line.size))
	font := getFont(line.font)
	ctx.SetFont(font)
	ctx.SetSrc(baseColor)
	pos := freetype.Pt(0, 0+int(ctx.PointToFixed(line.size)>>6))
	ctx.DrawString(line.content, pos)
}

func getFont(wantedFont string) *truetype.Font {
	fontBytes, err := ioutil.ReadFile("./fonts/Rubik/static/Rubik-" + wantedFont + ".ttf")
	if err != nil {
		fmt.Println("ohoh, I encountered an error while fetching the font: " + err.Error())
	}
	font, fonterr := freetype.ParseFont(fontBytes)
	if fonterr != nil {
		fmt.Println("ohoh, I encountered an error while parsing the font: " + fonterr.Error())
	}
	return font
}

func writeImage(img socialImage) {
	// Encode as PNG.
	f, _ := os.Create(img.name + ".png")
	png.Encode(f, img.src)
	fmt.Println("The " + img.name + " image is ready and outputted as " + img.name + ".png !")
}
