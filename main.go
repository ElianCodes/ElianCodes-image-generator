package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/golang/freetype"
	"github.com/nfnt/resize"
)

type size struct {
	width  int
	height int
}

type socialImage struct {
	size      size
	name      string
	baseColor color.Color
	src       image.Image
}

func main() {
	var defaultSocialImageSize size = size{width: 2024, height: 1012}
	generateImage("test", defaultSocialImageSize)
}

func getRandomColor() color.Color {
	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	availableColors := [5]color.Color{
		color.RGBA{253, 186, 116, 0xff}, // ElianCodes Orange
		color.RGBA{134, 239, 172, 0xff}, // ElianCodes Green
		color.RGBA{147, 197, 253, 0xff}, // ElianCodes Blue
		color.RGBA{252, 165, 165, 0xff}, // ElianCodes Red
		color.RGBA{240, 171, 252, 0xff}, // ElianCodes Purple
	}
	fmt.Println("The color is picked!")
	return availableColors[rand.Intn(len(availableColors))]
}

func generateImage(imageName string, size size) {
	fmt.Println("starting generation of " + imageName)

	// initialize a new (empty) image
	img := image.NewRGBA(image.Rect(0, 0, size.width, size.height))

	usedColor := getRandomColor()
	uniformColor := image.NewUniform(usedColor)

	var socialImg socialImage = socialImage{size: size, name: imageName, baseColor: usedColor, src: img}

	// Draw the background
	draw.Draw(img, image.Rect(size.width/2, 0, size.width, size.height), uniformColor, image.Point{}, draw.Src)
	draw.Draw(img, image.Rect(0, 0, size.width/2, size.height), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// Set the ElianCodes image in place
	readHeroImg, err := os.Open("./assets/hero.png")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer readHeroImg.Close()
	heroImg, err := png.Decode(readHeroImg)
	heroImg = resize.Resize(600, 0, heroImg, resize.Lanczos3)

	rightMiddlePart := image.Pt((size.width/2*-1)-200, -200)
	draw.Draw(img, img.Bounds(), heroImg, rightMiddlePart, draw.Over)

	// Set the text
	ctx := freetype.NewContext()
	ctx.SetDPI(300)
	ctx.SetFontSize(42)
	ctx.SetClip(img.Bounds())
	fontBytes, err := ioutil.ReadFile("./fonts/Rubik/static/Rubik-Regular.ttf")
	if err != nil {
		fmt.Println("ohoh, I encountered an error while fetching the font: " + err.Error())
	}
	font, err := freetype.ParseFont(fontBytes)
	ctx.SetFont(font)
	ctx.SetDst(img)
	ctx.SetSrc(image.NewUniform(usedColor))
	ctx.DrawString("Elian Codes", freetype.Pt(0, 0+int(ctx.PointToFixed(42)>>6)))

	writeImage(socialImg)
}

func writeImage(img socialImage) {
	// Encode as PNG.
	f, _ := os.Create(img.name + ".png")
	png.Encode(f, img.src)
	fmt.Println("The " + img.name + " image is ready!")
}
