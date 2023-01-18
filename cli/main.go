package main

import (
	"fmt"
	"github.com/boombuler/led"
	"image/color"
	"time"
)

var RED color.RGBA = color.RGBA{0xFF, 0x00, 0x00, 0xFF}
var GREEN color.RGBA = color.RGBA{0x00, 0xFF, 0x00, 0xFF}
var BLUE color.RGBA = color.RGBA{0x00, 0x00, 0xFF, 0xFF}

func main() {
	for devInfo := range led.Devices() {
		dev, err := devInfo.Open()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer dev.Close()
		dev.SetColor(RED)
		time.Sleep(1 * time.Second) // Wait 1 second since device turns off after call to close()
		dev.SetColor(GREEN)
		time.Sleep(1 * time.Second)
		dev.SetColor(BLUE)
		time.Sleep(1 * time.Second)
	}
}
