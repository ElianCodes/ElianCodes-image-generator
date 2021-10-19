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
	"time"

	"github.com/golang/freetype"
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

func generateImage(imageName string, size size) {
	fmt.Println("starting generation of " + imageName)

	// initialize a new (empty) image
	img := image.NewRGBA(image.Rect(0, 0, size.width, size.height))

	usedColor := getRandomColor().color
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
