package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func main() {
	var format string
	flag.StringVar(&format, "f", "", "output format. Required. One of png, jpg, gif.")
	flag.Parse()
	if len(flag.Args()) > 0 {
		fmt.Fprintln(os.Stderr, "usage: imgconv -t=png|jpg|gif < INPUT > OUTPUT")
		os.Exit(1)
	}
	info, _ := os.Stdout.Stat()
	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Fprintln(os.Stderr, "Refusing to write to character device. Redirect output to a pipe or reqular file.")
		os.Exit(1)
	}
	img, _, err := image.Decode(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	format = strings.ToLower(format)
	switch format {
	case "jpg", "jpeg":
		err = jpeg.Encode(os.Stdout, img, nil)
	case "png":
		err = png.Encode(os.Stdout, img)
	case "gif":
		err = gif.Encode(os.Stdout, img, nil)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
