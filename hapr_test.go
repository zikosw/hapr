package hapr

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestX(t *testing.T) {

	w := bufio.NewWriter(os.Stdout)

	hh := New(atom.Div, nil,
		New(atom.P, nil, Text("hi")),
		New(atom.P, nil, Text("bye")),
		New(atom.Img,
			Attrs(
				Attr("height", "100px"),
				Attr("width", "100px"),
			),
		),
	)

	spew.Dump(hh)

	fmt.Printf("\n\n\n")
	fmt.Printf("%#+v\n\n\n", hh)
	html.Render(w, hh)

	w.Flush()

}
