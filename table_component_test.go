package hapr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func TestMakeTable(t *testing.T) {

	type Coin struct {
		Name  string
		Price int
	}

	columnRender := func(col string) *html.Node {
		return Th_(Text(col))
	}
	rowRender := func(c Coin) *html.Node {
		return Tr_(
			Td_(Text(c.Name)),
			Td_(Text(fmt.Sprint(c.Price))),
		)
	}
	columns := []string{"Name", "Price"}

	type args struct {
		opts MakeTableOpts
	}
	tests := []struct {
		name    string
		args    args
		want    *html.Node
		wantErr error
	}{
		{
			name: "nil data",
			args: args{
				opts: MakeTableOpts{
					Columns:      columns,
					ColumnRender: columnRender,
					Render:       rowRender,
				},
			},
			want: Table(
				AttrClass(""),
				Thead(
					AttrClass(""),
					Tr_(
						Th_(Text("Name")),
						Th_(Text("Price")),
					),
				),
				Tbody(
					AttrClass(""),
				),
			),
		},
		{
			name: "lenght 0 data",
			args: args{
				opts: MakeTableOpts{
					Columns:      columns,
					ColumnRender: columnRender,
					Render:       rowRender,
					Data:         []Coin{},
				},
			},
			want: Table(
				AttrClass(""),
				Thead(
					AttrClass(""),
					Tr_(
						Th_(Text("Name")),
						Th_(Text("Price")),
					),
				),
				Tbody(
					AttrClass(""),
				),
			),
		},
		{
			name: "some data",
			args: args{
				opts: MakeTableOpts{
					Class:        "table-class",
					Columns:      columns,
					ColumnRender: columnRender,
					HeaderClass:  "head-class",
					BodyClass:    "body-class",
					Render:       rowRender,
					Data: []Coin{
						{Name: "Bitcoin", Price: 1800000},
						{Name: "Ethereum", Price: 300000},
					},
				},
			},
			want: Table(
				AttrClass("table-class"),
				Thead(
					AttrClass("head-class"),
					Tr_(
						Th_(Text("Name")),
						Th_(Text("Price")),
					),
				),
				Tbody(
					AttrClass("body-class"),
					Tr_(
						Td_(Text("Bitcoin")),
						Td_(Text("1800000")),
					),
					Tr_(
						Td_(Text("Ethereum")),
						Td_(Text("300000")),
					),
				),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)
			got, err := MakeTable(tt.args.opts)
			if tt.wantErr != nil {
				r.Error(err)
			} else {
				r.NoError(err)
				r.Equal(tt.want, got)
			}
		})
	}
}
