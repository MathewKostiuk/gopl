package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif" // register GIF decoder
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var inputflag string
var outputflag string

func init() {
	flag.StringVar(&inputflag, "i", "", "specify the path of the input image")
	flag.StringVar(&outputflag, "o", "jpg", "specify the image's output format")
	flag.Parse()
}

func main() {
	if inputflag == "" {
		fmt.Fprintln(os.Stderr, "jpeg: please specify the path of the input image")
		os.Exit(1)
	}

	reader, err := os.Open(inputflag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
	defer reader.Close()
	if err := convertImage(reader, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func convertImage(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	switch outputflag {
	case "jpg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, &gif.Options{})
	default:
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	}
}
