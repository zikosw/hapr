package hapr

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"golang.org/x/net/html"
	a "golang.org/x/net/html/atom"
)

func TestX(t *testing.T) {

	w := bufio.NewWriter(os.Stdout)

	hh := New(a.Div, nil,
		New(a.P, nil, Text("hi")),
		New(a.P, nil, Text("bye")),
		New(a.Img,
			Attrs(
				Attr(a.Height, "100px"),
				Attr(a.Width, "100px"),
				Attr(a.Style, "border-radius:10px;color:red"),
			),
		),
	)

	spew.Dump(hh)

	fmt.Printf("\n\n\n")
	fmt.Printf("%#+v\n\n\n", hh)
	html.Render(w, hh)

	w.Flush()

	fmt.Println("\nLookUP", a.Lookup([]byte("style")))

}
