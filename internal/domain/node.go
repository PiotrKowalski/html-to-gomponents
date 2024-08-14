package domain

import (
	"fmt"
	"strings"
)

type CustomNode struct {
	ParentNode *CustomNode
	Level      uint
	Type       string
	Attrs      []attr
	Nodes      []*CustomNode
}

func (n *CustomNode) AddAttr(key, value string) {
	n.Attrs = append(n.Attrs, attr{key: key, value: value})
}

func (n *CustomNode) String() string {
	str := ""

	str += n.Type
	str += "("

	if len(n.Attrs) > 0 {
		for _, v := range n.Attrs {
			if len(v.value) > 0 {
				str = fmt.Sprintf("%s%s(\"%s\"),", str, v.key, v.value)
			} else {
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

type attr struct {
	key, value string
}
