/*
Package flagicon provides a function to generate flag icons from image files

By default the flag icon will be of size 14x14 pixels; the original image will
be resized to this size, with a 2px border with color #808080.

If the image is not square, flagicon first crops the image (centered at the center
of the image) such that it is square before it resizes the image.

*/

package main

import (
	"log"
	"github.com/alexflint/go-arg"
	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
	"image"
	"math"
)

// resizeImage resizes the base image to a 14x14 square used for the flagicon
// If the base image is not square, resizeImage first crops the image into a
// square shape, centered on the center of the image.
func resizeImage(base image.Image) image.Image {
	if base.Bounds().Dx() != base.Bounds().Dy() {
		newHeight := int(math.Min(float64(base.Bounds().Dx()), float64(base.Bounds().Dy())))
		cropped, err := cutter.Crop(base, cutter.Config{
			Width:  newHeight,
			Height: newHeight,
			Mode:   cutter.Centered,
		})
		if err != nil {
			log.Fatal(err)
		}
		return resizeToThumbnail(cropped)
	}
	return resizeToThumbnail(base)
}

// resizeToThumbnail resizes the base image to a 14x14 square
func resizeToThumbnail(base image.Image) image.Image {
	return resize.Thumbnail(uint(14), uint(14), base, resize.Bilinear)
}

func main() {
	var args struct {
		Input string `arg:"positional" help:"Path to input image file"`
		Output string `arg:"-o" help:"Output file name"`
	}	
	args.Output = "out.png"
	arg.MustParse(&args)

	im, err := gg.LoadImage(args.Input)
	if err != nil {
		log.Fatal(err)
	}
	image := resizeImage(im)

	dc := gg.NewContext(14, 14)

	// Clip and draw circle image
	dc.DrawCircle(7, 7, 7)
	dc.Clip()
	dc.DrawImageAnchored(image, 7, 7, 0.5, 0.5)

	// Draw image border
	dc.DrawCircle(7, 7, 7)
	dc.SetHexColor("#808080")
	dc.SetLineWidth(2)
	dc.Stroke()

	dc.SavePNG(args.Output)
}
