package app

import (
	"golang.org/x/net/context"
	"html-to-gomponents/internal/domain"
	"html-to-gomponents/internal/requests"
	"html-to-gomponents/internal/responses"
)

type rawParseHandler struct {
	parser    domain.HTMLParser
	formatter domain.Formatter
}

func (h rawParseHandler) Handle(ctx context.Context, in requests.Parse) (responses.Parse, error) {
	cNode, err := h.parser.FromBytes(in.Body)
	if err != nil {
		return responses.Parse{}, err
	}

	_, err = h.formatter.Format(cNode)
	if err != nil {
		return responses.Parse{}, err
	}

	return responses.Parse{Body: cNode.String()}, nil
}
