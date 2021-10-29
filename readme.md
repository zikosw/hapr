# Hapr

HTML builder in Golang


```golang
[]*html.Node{
    Input(Attrs(
        Attr(a.Name, "Book"),
        Attr(a.Value, "30"),
    )),
    Input(Attrs(
        Attr(a.Name, "TV"),
        Attr(a.Value, "100"),
    )),
    Input(Attrs(
        Attr(a.Name, "Paper"),
        Attr(a.Value, "2"),
    )),
}
```

Same thing with `ForEach`

```golang
type Thing struct {
    Name  string
    Price int
}

things := []Thing{
    {
        Name:  "Book",
        Price: 30,
    },
    {
        Name:  "TV",
        Price: 100,
    },
    {
        Name:  "Paper",
        Price: 2,
    },
}

ForEach(
    things,
    func(t Thing) *html.Node {
        return New(a.Input, Attrs(
            Attr(a.Name, t.Name),
            Attr(a.Value, fmt.Sprint(t.Price)),
        ))
    },
)

```


Render table

```golang

Table(
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
)

// same thing with MakeTable

MakeTable(MakeTableOpts{
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
})


```
