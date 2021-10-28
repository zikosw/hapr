package hapr

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
	a "golang.org/x/net/html/atom"
)

func TestRender(t *testing.T) {

	w := bufio.NewWriter(os.Stdout)

	hh := New(a.Div, nil,
		New(a.P, nil, Text("hi")),
		P_(Text("bye")),
		Img(
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

func TestForEach(t *testing.T) {
	type args struct {
		items    interface{}
		renderer interface{}
	}

	type InputData struct {
		Name  string
		Value int
	}

	tests := []struct {
		name    string
		args    args
		want    []*html.Node
		wantErr error
	}{
		{
			name: "put nil get nil",
		},
		{
			name: "err if item not a slice",
			args: args{
				items: 1,
			},
			wantErr: ErrItemsNotSlice,
		},
		{
			name: "err nil renderer",
			args: args{
				items: []int{},
			},
			wantErr: ErrRendererNotFunc,
		},
		{
			name: "renderer no arg",
			args: args{
				items:    []int{},
				renderer: func() {},
			},
			wantErr: ErrRendererArgsNotMatch,
		},
		{
			name: "renderer args number not match",
			args: args{
				items:    []int{},
				renderer: func(int, int) {},
			},
			wantErr: ErrRendererArgsNotMatch,
		},
		{
			name: "items's type and renderer's param type not match",
			args: args{
				items:    []int{},
				renderer: func(string) {},
			},
			wantErr: ErrRendererArgsNotMatch,
		},
		{
			name: "renderer return not match",
			args: args{
				items:    []int{},
				renderer: func(int) int { return 0 },
			},
			wantErr: ErrRendererReturnNotMatch,
		},
		{
			name: "items's type and renderer's param type match, empty items got nil",
			args: args{
				items: []args{},
				renderer: func(args) *html.Node {
					return nil
				},
			},
			want: nil,
		},
		{
			name: "items's type and renderer's param type match, nil items got nil",
			args: args{
				items: nil,
				renderer: func(args) *html.Node {
					return nil
				},
			},
			want: nil,
		},
		{
			name: "correct length and values",
			args: args{
				items: []InputData{
					{
						Name:  "day",
						Value: 29,
					},
					{
						Name:  "month",
						Value: 10,
					},
					{
						Name:  "year",
						Value: 2021,
					},
				},
				renderer: func(in InputData) *html.Node {
					return New(a.Input, Attrs(
						Attr(a.Name, in.Name),
						Attr(a.Value, fmt.Sprint(in.Value)),
					))
				},
			},
			want: []*html.Node{
				Input(Attrs(
					Attr(a.Name, "day"),
					Attr(a.Value, "29"),
				)),
				Input(Attrs(
					Attr(a.Name, "month"),
					Attr(a.Value, "10"),
				)),
				Input(Attrs(
					Attr(a.Name, "year"),
					Attr(a.Value, "2021"),
				)),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)
			got, err := ForEach(tt.args.items, tt.args.renderer)
			if tt.wantErr != nil {
				r.Error(err)
				r.ErrorIs(err, tt.wantErr)
			} else {
				r.NoError(err)
				r.Len(got, len(tt.want))
				r.Equal(tt.want, got)
			}
		})
	}
}
