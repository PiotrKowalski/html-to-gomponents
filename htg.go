// Package  exposes
package main

import (
	"golang.org/x/net/context"
	"html-to-gomponents/internal/app"
	"html-to-gomponents/internal/requests"
)

var (
	localApp = app.New(app.Config{})
)

func Parse(ctx context.Context, in []byte) (string, error) {
	res, err := localApp.ParseHandler.Handle(ctx, requests.Parse{Body: in})
	if err != nil {
		return "", err
	}
	return res.Body, nil
}
