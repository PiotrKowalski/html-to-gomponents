package parser

import (
	"bytes"
	"golang.org/x/net/html"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"html-to-gomponents/internal/domain"
	"strings"
)

type Parser struct {
}

func (p Parser) FromBytes(in []byte) (*domain.CustomNode, error) {
	hNode, err := html.Parse(bytes.NewReader(in))
	if err != nil {
		return nil, err
	}

	var findBody func(n *html.Node) *html.Node
	findBody = func(n *html.Node) *html.Node {
		if n.Data == "body" {
			return n
		}
		var e *html.Node
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			e = findBody(c)
		}

		return e
	}

	body := findBody(hNode)

	var f func(*html.Node, *domain.CustomNode) *domain.CustomNode
	f = func(n *html.Node, cNode *domain.CustomNode) *domain.CustomNode {
		if n.Type == html.ElementNode {
			cNode.Type = htmlToGomponentsName(n.Data)

			for _, attr := range n.Attr {
				key := htmlToGomponentsName(attr.Key)
				cNode.Attrs[key] = attr.Val
			}
		}

		if n.Type == html.TextNode && len(strings.TrimSpace(n.Data)) > 0 {
			cNode.ParentNode.Attrs["g.Text"] = strings.TrimSpace(n.Data)
		}

		var i uint
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			cNode.Nodes = append(cNode.Nodes, &domain.CustomNode{ParentNode: cNode, Level: cNode.Level + 1, Attrs: make(map[string]string)})
			cNode.Nodes[i] = f(c, cNode.Nodes[i])
			i++
		}
		return cNode
	}
	output := &domain.CustomNode{}
	out := f(body, output)

	return out, nil
}

func New() Parser {
	return Parser{}
}

func htmlToGomponentsName(in string) string {
	switch in {
	case "head":
		return "c.Head"
	case "thead":
		return "THead"
	case "tbody":
		return "TBody"
	case "id":
		return "ID"
	default:
		return cases.Title(language.English).String(in)
	}
}
