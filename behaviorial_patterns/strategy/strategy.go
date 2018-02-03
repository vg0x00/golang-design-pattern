package strategy

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

// Different implementtations with the same interface

type OutputStrategy interface {
	Draw() error
}

type ConsoleSquare struct {
}

func (c *ConsoleSquare) Draw() error {
	// return errors.New("not implemented yet")
	fmt.Println("console square")
	return nil
}

type ImageSquare struct {
	Path string
}

func (i *ImageSquare) Draw() error { // for exercise image std lib
	// return errors.New("not implemented yet")
	width := 800
	height := 600

	// Uniform: for single color image
	backgroundImage := image.Uniform{color.RGBA{R: 70, G: 80, B: 70, A: 0}}
	origin := image.Point{0, 0}
	quality := &jpeg.Options{Quality: 75}
	backgroundRect := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: width, Y: height},
	})

	draw.Draw(finalImage, finalImage.Bounds(), &backgroundImage, origin, draw.Src)

	// draw square
	squareWidth := 200
	squareHeight := 200
	squareColor := image.Uniform{color.RGBA{R: 255, G: 0, B: 0, A: 1}}
	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (width / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})
	squareImage := image.NewRGBA(square)
	// draw new red square on the finalImage
	draw.Draw(finalImage, squareImage.Bounds(), &squareColor, origin, draw.Src)

	// write content to file
	imgFile, err := os.Create(i.Path)
	if err != nil {
		return fmt.Errorf("image file create failed: %s\n", err.Error())
	}
	defer imgFile.Close()

	if err = jpeg.Encode(imgFile, finalImage, quality); err != nil {
		return fmt.Errorf("error in wring image to disk: %s\n", err.Error())
	}

	// now, image should locate  at Path
	return nil
}

var option string
var outputObj OutputStrategy

func Main() error {
	switch option {
	case "console":
		outputObj = &ConsoleSquare{}
	default:
		outputObj = &ImageSquare{Path: "/Users/vg0x00/Desktop/test.jpeg"}
	}

	return outputObj.Draw()
}
