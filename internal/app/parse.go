package app

import (
	"golang.org/x/net/context"
	"html-to-gomponents/internal/domain"
	"html-to-gomponents/internal/requests"
	"html-to-gomponents/internal/responses"
	"log"
)

type parseHandler struct {
	parser    domain.HTMLParser
	formatter domain.Formatter
}

func (h parseHandler) Handle(ctx context.Context, in requests.Parse) (responses.Parse, error) {

	cNode, err := h.parser.FromBytes(in.Body)
	if err != nil {
		return responses.Parse{}, err
	}
	log.Println(11)

	format, err := h.formatter.Format(cNode)
	if err != nil {
		return responses.Parse{}, err
	}

	return responses.Parse{Body: format}, nil
}
