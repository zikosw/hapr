package hapr

import (
	"fmt"

	"golang.org/x/net/html"
)

type TableColumn struct {
	Label  string
	Class  string
	Render func(interface{}) *html.Node
}

type TableRow struct {
	Data   interface{}
	Render func(interface{}) *html.Node
	Class  string
}

// type TableCell struct {
// 	DataIndex string
// 	Data      interface{}
// 	Render    func(interface{}) *html.Node
// 	Class     string
// }

type MakeTableOpts struct {
	Columns      []string
	ColumnRender interface{}
	Class        string
	HeaderClass  string
	BodyClass    string
	// Data         []map[string]interface{}
	Data   interface{}
	Render interface{}
}

func MakeTable(opts MakeTableOpts) (*html.Node, error) {

	if len(opts.Columns) == 0 {
		return nil, fmt.Errorf("column length is 0")
	}

	headers, err := ForEach(opts.Columns, opts.ColumnRender)
	if err != nil {
		return nil, fmt.Errorf("heads render: %w", err)
	}

	// rows := []*html.Node{}
	// for _, d := range data {
	// 	r :=
	// 	rows = append(rows, r)
	// }
	rows, err := ForEach(opts.Data, opts.Render)
	if err != nil {
		return nil, fmt.Errorf("rows render: %w", err)
	}

	return Table(
		AttrClass(opts.Class),
		Thead(
			AttrClass(opts.HeaderClass),
			Tr_(
				headers...,
			),
		),
		Tbody(
			AttrClass(opts.BodyClass),
			rows...,
		),
	), nil
}

// (defn table3 [{:keys [columns data table-classes head-wrapper-classes body-classes row-classes]}]
// 	[:table {:class table-classes}
// 	 [:thead
// 	  [:tr {:class head-wrapper-classes}
// 	   (for [col columns]
// 		 (tb-header3 col))]]
// 	 [:tbody {:class body-classes}
// 	  (for [row data]
// 		[:tr {:class row-classes}
// 		 (for [col columns]
// 		   (tb-row3 (merge col {:data row})))])]])
