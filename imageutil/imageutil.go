// Package imageutil adds functions for working with images.
package imageutil // import "zgo.at/utils/imageutil"

import (
	"crypto/md5"
	"fmt"
)

// ColorHash generates a random RGB background colour based on the input string
// with a foreground colour to match. The foreground colour is either all black
// or white.
func ColorHash(s string) (bg, fg string) {
	h := md5.New() // fnv is faster, but doesn't give a good distribution for this.
	h.Write([]byte(s))
	c := string(h.Sum(nil))

	fg = "#000000"
	if luma(c[0], c[1], c[2]) < 150 {
		fg = "#ffffff"
	}
	return fmt.Sprintf("#%.2x%.2x%.2x", c[0], c[1], c[2]), fg
}

// Get luma/brightness (0-255). Constants from BT.601.
func luma(r, g, b byte) float32 {
	return .299*float32(r) + .587*float32(g) + .114*float32(b)
}
