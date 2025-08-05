package main

import (
	"fmt"

	"github.com/DustinMeyer1010/converters/internal/image/png"
)

func main() {

	test, _ := png.CreatePNG("/Users/dustinmeyer/Documents/Github/image/internal/image/testimages/valid.png")

	fmt.Print(test)

}
