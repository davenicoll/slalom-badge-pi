package main

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	epaper "github.com/bestbug456/epaper"
)

var epd_width int = 122
var epd_height int = 250

func main() {
	e := epaper.CreateEpd()
	defer e.Close()
	defer e.Clear()
	e.Init()
	e.Clear()

	fmt.Printf("Display\n")
	bg := image.NewRGBA(image.Rect(0, 0, epd_height, epd_width))
	addLabel(bg, 20, 30, "Hello")
	e.Display(e.GetBuffer(bg))

	fmt.Printf("sleeping\n")
	time.Sleep(10 * time.Second)
}

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}
