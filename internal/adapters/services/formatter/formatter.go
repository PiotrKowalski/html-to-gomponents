package formatter

import (
	"go/format"
	"html-to-gomponents/internal/domain"
)

type Formatter struct {
}

func (f Formatter) Format(node *domain.CustomNode) (string, error) {
	b := []byte(`package example
import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)
func example() Node {
	return ` + node.String() + `
}`)
	dist, err := format.Source(b)
	if err != nil {
		return "", err
	}

	return string(dist), nil
}

func New() Formatter {
	return Formatter{}
}
