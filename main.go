package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
)

type hexColor color.RGBA

func (s hexColor) String() string { return fmt.Sprintf("#%02x%02x%02x", s.R, s.G, s.B) }

func (s *hexColor) Set(v string) error {
	if _, err := fmt.Sscanf(v, "#%02x%02x%02x", &s.R, &s.G, &s.B); err != nil {
		return fmt.Errorf("invalid color: %q", v)
	}
	s.A = 0xff
	return nil
}

type options struct {
	colorA, colorB hexColor
	stripeHeight   int
}

func main() {
	var opts options
	flag.Var(&opts.colorA, "a", "Topmost stripe color")
	flag.Var(&opts.colorB, "b", "Bottom-most stripe color")
	flag.IntVar(&opts.stripeHeight, "w", 30, "Height of the stripes")
	flag.Parse()
	if err := printStripes(os.Stdout, &opts); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// fill fills pix with the color c, assuming it is the in-memory representation
// of a RGBA image.
// It panics if len(pix) is not a multiple of 4.
func fill(pix []byte, c color.RGBA) {
	for i := 0; i < len(pix); i += 4 {
		pix[i] = c.R
		pix[i+1] = c.G
		pix[i+2] = c.B
		pix[i+3] = c.A
	}
}

func printStripes(out io.Writer, opts *options) error {
	// Since the image is meant to be tiled, its width can be brought down as low as
	// 1 pixel; but then it would be hard to make out the colors in an image picker.
	img := image.NewRGBA(image.Rect(0, 0, 16, opts.stripeHeight*2))
	m := len(img.Pix) / 2
	fill(img.Pix[:m], color.RGBA(opts.colorA))
	fill(img.Pix[m:], color.RGBA(opts.colorB))
	return png.Encode(out, img)
}
