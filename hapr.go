package hapr

import (
	"fmt"
	"reflect"

	"golang.org/x/net/html"
	a "golang.org/x/net/html/atom"
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

func Attr(k a.Atom, v string) html.Attribute {
	return html.Attribute{
		Key: k.String(),
		Val: v,
	}
}

func New(dataAtom a.Atom, attrs []html.Attribute, children ...*html.Node) *html.Node {
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

func NewNA(dataAtom a.Atom, children ...*html.Node) *html.Node {
	return New(dataAtom, nil, children...)
}

var ErrItemsNotSlice = fmt.Errorf("item must be slice")
var ErrRenderNotFunc = fmt.Errorf("render must be function")
var ErrRenderArgsNotMatch = fmt.Errorf("render args not match")
var ErrRenderReturnNotMatch = fmt.Errorf("render return not match")

var _renderReturn *html.Node

func ForEach(items interface{}, render interface{}) ([]*html.Node, error) {
	if items == nil {
		return nil, nil
	}

	typeOfItems := reflect.TypeOf(items)
	// fmt.Printf("%+v\n", typeOfItems)
	if typeOfItems.Kind() != reflect.Slice {
		// fmt.Println(`kind`, typeOfItems.Kind())
		return nil, ErrItemsNotSlice
	}

	if render == nil {
		return nil, ErrRenderNotFunc
	}
	valOfRenderer := reflect.ValueOf(render)
	typeOfRenderer := reflect.TypeOf(render)

	// fmt.Printf("val render:%+v\n", valOfRenderer)
	// fmt.Printf("type render:%+v\n", typeOfRenderer)
	if valOfRenderer.IsNil() {
		return nil, ErrRenderNotFunc
	}

	// fmt.Println("num in", typeOfRenderer.NumIn())
	if typeOfRenderer.NumIn() != 1 {
		return nil, ErrRenderArgsNotMatch
	}

	typeOfArg := typeOfRenderer.In(0)
	// fmt.Println("type of arg", typeOfArg)
	// fmt.Println("type of items", typeOfItems)
	// fmt.Println("type of items elem", typeOfItems.Elem())
	typeOfItemsElem := typeOfItems.Elem()
	// fmt.Println("-item elem type", typeOfItemsElem.String())
	// fmt.Println("-renderer arg type", typeOfArg.String())
	if typeOfArg.String() != typeOfItemsElem.String() {
		return nil, ErrRenderArgsNotMatch
	}

	// typeOfRenderOut
	if typeOfRenderer.NumOut() != 1 {
		return nil, ErrRenderReturnNotMatch
	}

	typeOfRet := typeOfRenderer.Out(0)
	// typeOfrenret := reflect.TypeOf(_renderReturn)
	// fmt.Println("ret", reflect.TypeOf(_renderReturn))
	// fmt.Println("ret2", typeOfrenret.String())
	// fmt.Println("ret3", typeOfrenret.Name())
	if typeOfRet.String() != reflect.TypeOf(_renderReturn).String() {
		return nil, ErrRenderReturnNotMatch
	}

	valOfItems := reflect.ValueOf(items)
	if valOfItems.Len() == 0 {
		return nil, nil
	}

	// ---- items and renderer is correct ----

	results := []*html.Node{}
	for i := 0; i < valOfItems.Len(); i++ {
		item := valOfItems.Index(i)
		outValues := valOfRenderer.Call([]reflect.Value{item})
		if len(outValues) == 1 {
			out := outValues[0]
			// fmt.Printf("out:%+v\n", out.Type())
			// out.Type()

			res, ok := out.Interface().(*html.Node)
			if !ok {
				panic("why cant cast out to it type?")
			}
			results = append(results, res)
		} else {
			panic("why no out values?")
		}
	}

	return results, nil
}

func AttrClass(class string) []html.Attribute {
	return Attrs(Attr(a.Class, class))
}
