package app

import (
	serviceformatter "html-to-gomponents/internal/adapters/services/formatter"
	serviceparser "html-to-gomponents/internal/adapters/services/parser"
)

type Config struct {
}

type Application struct {
	ParseHandler parseHandler
}

func New(cfg Config) Application {
	parser := serviceparser.New()
	formatter := serviceformatter.New()

	return Application{ParseHandler: parseHandler{parser: parser, formatter: formatter}}
}
