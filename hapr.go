package hapr

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func Text(t string) *html.Node {
	return &html.Node{
		Type: html.TextNode,
		Data: t,
	}
}

func Attrs(attrs ...html.Attribute) []html.Attribute {
	return attrs
}

func Attr(k, v string) html.Attribute {
	return html.Attribute{
		Key: k,
		Val: v,
	}
}

func New(dataAtom atom.Atom, attrs []html.Attribute, children ...*html.Node) *html.Node {
	n := &html.Node{
		Type:     html.ElementNode,
		DataAtom: dataAtom,
		Data:     dataAtom.String(),
		Attr:     attrs,
	}

	for _, c := range children {
		if c != nil {
			n.AppendChild(c)
		}
	}
	return n
}
