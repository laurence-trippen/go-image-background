package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"

	"github.com/go-playground/colors"
)

func loadImageFromFile(path string) (image.Image, error) {
	f, err := os.Open(path)

	if nil != err {
		log.Fatal(err)
		return nil, err
	}

	defer f.Close()

	image, _, err := image.Decode(f)

	return image, err
}

func main() {
	cwd, err := os.Getwd()
	if nil != err {
		panic(err)
	}
	fmt.Println(cwd)

	if len(os.Args) != 3 {
		log.Fatal("Please specify arguments like this: <imgpath> <color>")
	}

	path := os.Args[1] // TODO: Validate path
	colorStr := os.Args[2]

	color, err := colors.Parse(colorStr)
	if nil != err {
		log.Fatal("Invalid color")
	}

	fmt.Println(color.String())

	inputImage, err := loadImageFromFile(path)
	if nil != err {
		log.Fatal(err)
	}

	fmt.Println(inputImage.Bounds())

	newImage := image.NewRGBA(inputImage.Bounds())

	// Draw Background
	draw.Draw(newImage, inputImage.Bounds(), &image.Uniform{color}, image.ZP, draw.Src)

	// Draw Image
	draw.Draw(newImage, inputImage.Bounds(), inputImage, image.ZP, draw.Over)

	out, err := os.Create("out.png")
	if nil != err {
		log.Fatal(err)
	}
	defer out.Close()

	err = png.Encode(out, newImage)
	if nil != err {
		log.Fatal(err)
	}

	fmt.Println("Done!")
}
