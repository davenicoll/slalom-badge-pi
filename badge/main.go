package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
)

var epd_width int = 250
var epd_height int = 122
var fonts []string
var test_output_file_count int = 1

func epd_test() {
	// e := epaper.CreateEpd()
	// defer e.Close()
	// defer e.Clear()
	// e.Init()
	// e.Clear()

	//fmt.Printf("Display\n")

	// fmt.Printf("sleeping\n")
	// time.Sleep(10 * time.Second)
}

func display_font(font_file string) {

	fmt.Printf("Displaying %s\n", font_file)

	var font_name string = font_file[strings.LastIndex(font_file, "/")+1:]
	var sample_text string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var sample_nums string = "0123456789"
	var sample_special string = "!\"£$%^&*()_+{}[]:;@'~#<>?/\\|"
	var sample_emoji string = "❤️😀🙁👍"
	var all_text string = strings.ToLower(sample_text) + "\n" + sample_text + "\n" + sample_nums + "\n" + sample_special

	bg := image.NewRGBA(image.Rect(0, 0, epd_height, epd_width))

	// PAGE 1
	dc := gg.NewContext(epd_width, epd_height)
	dc.DrawImage(bg, 0, 0)
	dc.SetColor(color.White)
	dc.DrawRectangle(0, 0, 250.0, 122.0)
	dc.Fill()
	dc.SetColor(color.Black)

	dc.DrawStringWrapped(font_name, 100, 10, 0.5, 0.5, 180, 1.4, gg.AlignLeft)
	if err := dc.LoadFontFace(font_file, 48); err != nil {
		fmt.Printf("ERROR: %s is an invalid TTF font, skipping\n", font_file)
		return
	}
	dc.DrawStringWrapped(sample_emoji, 100, 60, 0.5, 0.5, 180, 1, gg.AlignLeft)
	dc.SavePNG("test-output/screen_" + strconv.Itoa(test_output_file_count) + "_1.png")

	// PAGE 2
	dc.SetColor(color.White)
	dc.DrawRectangle(0, 0, 250.0, 122.0)
	dc.Fill()
	dc.SetColor(color.Black)

	if err := dc.LoadFontFace(font_file, 8); err != nil {
		fmt.Printf("ERROR: %s is an invalid TTF font, skipping\n", font_file)
		return
	}
	dc.DrawStringWrapped(all_text, 92, 18, 0.5, 0.5, 180, 1.4, gg.AlignLeft)

	if err := dc.LoadFontFace(font_file, 9); err != nil {
		panic(err)
	}
	dc.DrawStringWrapped(all_text, 92, 56, 0.5, 0.5, 180, 1.4, gg.AlignLeft)

	if err := dc.LoadFontFace(font_file, 10); err != nil {
		panic(err)
	}
	dc.DrawStringWrapped(all_text, 92, 98, 0.5, 0.5, 180, 1.4, gg.AlignLeft)

	//addLabel(bg, 20, 30, "Hello")
	//e.Display(e.GetBuffer(dc.Image()))
	dc.SavePNG("test-output/screen_" + strconv.Itoa(test_output_file_count) + "_2.png")

	// PAGE 3
	dc.SetColor(color.White)
	dc.DrawRectangle(0, 0, 250.0, 122.0)
	dc.Fill()
	dc.SetColor(color.Black)

	if err := dc.LoadFontFace(font_file, 11); err != nil {
		panic(err)
	}
	dc.DrawStringWrapped(all_text, 92, 24, 0.5, 0.5, 180, 1.5, gg.AlignLeft)

	if err := dc.LoadFontFace(font_file, 12); err != nil {
		panic(err)
	}
	dc.DrawStringWrapped(all_text, 92, 80, 0.5, 0.5, 180, 1.5, gg.AlignLeft)

	dc.SavePNG("test-output/screen_" + strconv.Itoa(test_output_file_count) + "_3.png")

	test_output_file_count++
}

func find_fonts(path string) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		if info.IsDir() {
			//fmt.Printf("Folder Name: %s\n", info.Name())
		} else {
			if strings.Contains(strings.ToLower(info.Name()), ".ttf") {
				//fmt.Printf("File Name: %s\n", info.Name())
				fonts = append(fonts, path)
			}
		}
		return nil
	})
}

func main() {
	font_directory, err := filepath.Abs("../fonts")
	if err != nil {
		log.Fatal(err)
	}
	find_fonts(font_directory)

	for len(fonts) > 0 {
		n := len(fonts) - 1
		display_font(fonts[n])
		fonts = fonts[:n]
	}
}
