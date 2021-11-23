package imagegenerate

import (
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type Size struct {
	Width  int
	Height int
}

type RandomColor struct {
	Name  string
	Color color.RGBA
}

type Line struct {
	Content string
	Color   color.Color
	Size    float64
	Font    string
}

type SocialImage struct {
	Size        Size
	Name        string
	BaseColor   color.Color
	Src         image.Image
	Link        Line
	Title       Line
	Subtitle    Line
	GeneralText Line
}

type GenerateImageFromAPI struct {
	Title string `json:"title"`
}

func GetRandomColor() RandomColor {
	availableColors := [5]RandomColor{
		{Name: "Orange", Color: color.RGBA{253, 186, 116, 0xff}},
		{Name: "Green", Color: color.RGBA{134, 239, 172, 0xff}},
		{Name: "Blue", Color: color.RGBA{147, 197, 253, 0xff}},
		{Name: "Red", Color: color.RGBA{252, 165, 165, 0xff}},
		{Name: "Purple", Color: color.RGBA{240, 171, 252, 0xff}},
	}
	rand.Seed(time.Now().UnixNano())
	var pickedColor RandomColor = availableColors[rand.Intn(len(availableColors))]
	fmt.Println("The color is picked, it is " + pickedColor.Name + "!")
	return pickedColor
}

func GenerateImage(generation SocialImage) SocialImage {
	fmt.Println("starting generation of " + generation.Name)

	// initialize a new (empty) image
	img := image.NewRGBA(image.Rect(0, 0, generation.Size.Width, generation.Size.Height))
	generation.Src = img
	uniformColor := image.NewUniform(generation.BaseColor)

	// Draw the background
	draw.Draw(img, image.Rect(generation.Size.Width/2, 0, generation.Size.Width, generation.Size.Height), uniformColor, image.Point{}, draw.Src)
	draw.Draw(img, image.Rect(0, 0, generation.Size.Width/2, generation.Size.Height), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// Set the ElianCodes image in place
	readHeroImg, err := os.Open("./assets/hero.png")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer readHeroImg.Close()
	heroImg, err := png.Decode(readHeroImg)
	heroImg = resize.Resize(600, 0, heroImg, resize.Lanczos3)

	rightMiddlePart := image.Pt((generation.Size.Width/2*-1)-200, -200)
	draw.Draw(img, img.Bounds(), heroImg, rightMiddlePart, draw.Over)

	// Set the text
	addText(generation, img)

	//writeImage(generation)
	fmt.Println("finished generation of " + generation.Name)
	return generation
}

func addText(generation SocialImage, sourceImg draw.Image) {
	fmt.Println("starting generation of Text")
	ctx := freetype.NewContext()
	ctx.SetDPI(300)
	ctx.SetClip(generation.Src.Bounds())
	ctx.SetDst(sourceImg)

	addLine(*ctx, generation.Title)
}

func addLine(ctx freetype.Context, line Line) {
	var baseColor image.Image = image.NewUniform(line.Color)
	ctx.SetFontSize(float64(line.Size))
	font := getFont(line.Font)
	ctx.SetFont(font)
	ctx.SetSrc(baseColor)
	pos := freetype.Pt(0, 0+int(ctx.PointToFixed(line.Size)>>6))
	ctx.DrawString(line.Content, pos)
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

func writeImage(img SocialImage) {
	// Encode as PNG.
	f, _ := os.Create(img.Name + ".png")
	png.Encode(f, img.Src)
	fmt.Println("The " + img.Name + " image is ready and outputted as " + img.Name + ".png !")
}
