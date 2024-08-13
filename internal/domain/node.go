package domain

import (
	"fmt"
	"strings"
)

type CustomNode struct {
	ParentNode *CustomNode
	Level      uint
	Type       string
	Attrs      map[string]string
	Nodes      []*CustomNode
}

func (c CustomNode) String() string {
	str := ""

	str += c.Type
	str += "("

	if len(c.Attrs) > 0 {
		for k, v := range c.Attrs {
			if len(v) > 0 {
				str = fmt.Sprintf("%s%s(\"%s\"),", str, k, v)
			} else {
				str = fmt.Sprintf("%s%s(),", str, k)
			}

		}
	}

	if len(c.Nodes) > 0 {
		for _, v := range c.Nodes {
			if v.Type != "" {
				str = fmt.Sprintf("%s\n%s%s,", str, strings.Repeat(" ", int(c.Level)), v)
			}
		}
	}

	str = fmt.Sprintf("%s\n%s)", str, strings.Repeat(" ", int(c.Level)))

	return removeBrackets(str)
}

func removeBrackets(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "[", ""), "]", "")
}
