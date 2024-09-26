package domain

import (
	"fmt"
	"slices"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CustomNode struct {
	ParentNode *CustomNode
	Level      uint
	customType bool
	Type       string
	Attrs      []Attr
	Nodes      []*CustomNode
}

func (n *CustomNode) SetType(in string) {
	switch in {
	case "head":
		n.Type = "c.Head"
	case "thead":
		n.Type = "THead"
	case "tbody":
		n.Type = "TBody"
	case "id":
		n.Type = "ID"
	case "path":
		n.Type = "path"
		n.customType = true
	case "circle":
		n.Type = "circle"
		n.customType = true
	case "rect":
		n.Type = "rect"
		n.customType = true
	case "line":
		n.Type = "line"
		n.customType = true
	case "polyline":
		n.Type = "line"
		n.customType = true
	case "svg":
		n.Type = "SVG"

	default:
		n.Type = cases.Title(language.English).String(in)

	}
}

func (n *CustomNode) AddAttr(key, value string) {
	if slices.Contains([]string{"xmlns", "fill", "viewBox", "stroke", "stroke-width", "fill-rule", "d", "stroke-linecap", "stroke-linejoin", "cx", "cy", "r", "x", "y", "rx", "ry", "x1", "x2", "y1", "y2", "points"}, key) {
		n.Attrs = append(n.Attrs, Attr{
			custom: true,
			key:    key,
			value:  value,
		})
		return
	}

	switch key {
	case "id":
		n.Attrs = append(n.Attrs, Attr{key: "ID", value: value})
		return
	case "g.Text":
		n.Attrs = append(n.Attrs, Attr{key: key, value: value})
		return
	default:
		n.Attrs = append(n.Attrs, Attr{key: cases.Title(language.English).String(key), value: value})
		return
	}
}

func (n *CustomNode) String() string {
	str := ""

	if n.customType {
		str += "g.El(\"" + n.Type + "\","
	} else {
		str += n.Type + "("
	}

	if len(n.Attrs) > 0 {
		for _, v := range n.Attrs {
			if v.custom {
				str = fmt.Sprintf("%sg.Attr(\"%s\",\"%s\"),", str, v.key, v.value)
			} else if len(v.value) > 0 {
				str = fmt.Sprintf("%s%s(\"%s\"),", str, v.key, v.value)
			} else if len(v.value) == 0 {
				str = fmt.Sprintf("%s%s(),", str, v.key)
			}
		}
	}

	if len(n.Nodes) > 0 {
		for _, v := range n.Nodes {
			if v.Type != "" {
				str = fmt.Sprintf("%s\n%s%s,", str, strings.Repeat(" ", int(n.Level)), v)
			}
		}
	}

	str = fmt.Sprintf("%s\n%s)", str, strings.Repeat(" ", int(n.Level)))

	return removeBrackets(str)
}

func removeBrackets(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "[", ""), "]", "")
}

type Attr struct {
	custom     bool
	key, value string
}
