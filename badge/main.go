package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	epaper "github.com/bestbug456/epaper"
	"github.com/fogleman/gg"
)

var epd_width int = 250
var epd_height int = 122

var epd = epaper.CreateEpd()

var fonts = []struct {
	name string
	path string
	size int
}{
	{"sfpro", "/modern/SF-Pro.ttf", 20},
	{"emoji", "/emoji/AndroidEmoji.ttf", 48},
	{"minecraft", "/8bit/minecraft_10.ttf", 16},
	{"m5x7", "/8bit/m5x7.ttf", 16},
	{"retroville", "/8bit/Retroville NC.ttf", 10},
	{"retrogaming", "/8bit/Retro Gaming.ttf", 11},
	{"pixeloid", "/8bit/PixeloidSans.ttf", 9},
	{"pixeloid-bold", "/8bit/PixeloidSans-Bold.ttf", 9},
	{"pixelade", "/8bit/PIXELADE.TTF", 13},
	{"iconbittwo", "/8bit/IconBitTwo.ttf", 10},
	{"iconbitone", "/8bit/IconBitOne.ttf", 10},
	{"atari", "/8bit/AtariST8x16SystemFont.ttf", 16},
}

func display_font_test(font_path string, font_size int) {

	fmt.Printf("Displaying %s\n", font_path)

	var font_name string = font_path[strings.LastIndex(font_path, "/")+1:]
	var sample_text string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var sample_nums string = "0123456789"
	var sample_special string = "!\"£$%^&*()_+{}[]:;@'~#<>?/\\|"
	//var sample_emoji string = "❤️😀🙁👍"
	var all_text string = strings.ToLower(sample_text) + "\n" + sample_text + "\n" + sample_nums + "\n" + sample_special

	bg := image.NewRGBA(image.Rect(0, 0, epd_height, epd_width))

	dc := gg.NewContext(epd_width, epd_height)
	dc.DrawImage(bg, 0, 0)
	dc.SetColor(color.White)
	dc.DrawRectangle(0, 0, 250.0, 122.0)
	dc.Fill()
	dc.SetColor(color.Black)

	dc.DrawStringWrapped(font_name+", size "+strconv.Itoa(font_size), 92, 20, 0.5, 0.5, 180, 1.4, gg.AlignLeft)

	if err := dc.LoadFontFace(font_path, float64(font_size)); err != nil {
		fmt.Printf("ERROR: %s is an invalid TTF font, skipping\n", font_path)
		return
	}
	dc.DrawStringWrapped(all_text, 92, 70, 0.5, 0.5, 180, 1.4, gg.AlignLeft)

	if img, ok := dc.Image().(*image.RGBA); ok {
		epd.Display(epd.GetBuffer(img))
	}
	time.Sleep(3 * time.Second)
}

func main() {
	// Initialize e-ink display
	defer epd.Close()
	defer epd.Clear()
	epd.Init()
	epd.Clear()

	font_directory, err := filepath.Abs("../fonts")
	if err != nil {
		log.Fatal(err)
	}

	for _, font := range fonts {
		display_font_test(font_directory+font.path, font.size)
	}
}
