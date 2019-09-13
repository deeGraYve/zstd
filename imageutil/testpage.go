// +build go_run_only

package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func init() {
	rand.Seed(42)
}

func main() {
	var html strings.Builder
	html.WriteString(`<div style="display: flex"><style>
		.white { color: #fff; text-shadow: 1px 2px rgba(0, 0, 0, .5) }
		.black { color: #000; text-shadow: 1px 2px rgba(255, 255, 255, .5) }
	</style>`)

	write(&html, grey)
	write(&html, red)
	write(&html, green)
	write(&html, blue)
	write(&html, rnd)

	fmt.Println(html.String() + `</div>`)
}

func write(html *strings.Builder, getColor func(int) (int, int, int)) {
	html.WriteString(`<div>`)

	for i := 0; i <= 255; i++ {
		r, g, b := getColor(i)

		luma := .299*float32(r) + .587*float32(g) + .114*float32(b)
		class := "dark"
		if luma < 150 {
			class = "white"
		}

		html.WriteString(fmt.Sprintf(
			`<div class="%s" style="background-color: %s; padding: 1em;">It’s a bloody aardvark!</div>`,
			class, fmt.Sprintf("#%.2x%.2x%.2x", r, g, b)))
	}
	html.WriteString(`</div>`)
}

func grey(i int) (int, int, int)  { return i, i, i }
func red(i int) (int, int, int)   { return i, 0, 0 }
func green(i int) (int, int, int) { return 0, i, 0 }
func blue(i int) (int, int, int)  { return 0, 0, i }
func rnd(i int) (int, int, int)   { return rand.Intn(255), rand.Intn(255), rand.Intn(255) }
