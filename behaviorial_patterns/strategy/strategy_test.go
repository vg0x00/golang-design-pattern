package strategy

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"testing"
)

func TestStrategy(t *testing.T) {
	t.Run("TestDrawImpl", func(t *testing.T) {
		consoleOutput := ConsoleSquare{}
		err := consoleOutput.Draw()
		if err != nil {
			t.Errorf("console implement wrong: %s\n", err.Error())
		}

		imageOutput := ImageSquare{Path: "/Users/vg0x00/Desktop/image.jpeg"}
		err = imageOutput.Draw()
		if err != nil {
			t.Errorf("image implement wrong: %s\n", err.Error())
		}
	})
}

func TestDraw(t *testing.T) {
	option = "console"
	err := Main()
	if err != nil {
		t.Error(err)
	}
	option = "image"
	err = Main()
	if err != nil {
		t.Error(err)
	}
}

func TestImage(t *testing.T) {
	img := image.Uniform{color.RGBA{R: 70, G: 80, B: 70, A: 0}}
	file, _ := os.Create("/Users/vg0x00/Desktop/img.jpeg")
	defer file.Close()
	err := jpeg.Encode(file, &img, nil)
	if err != nil {
		fmt.Println("image writing error: %s\n", err.Error())
	}

}
