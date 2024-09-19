package view

import (
	"errors"
	"github.com/labstack/echo/v4"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	hxhttp "github.com/maragudk/gomponents-htmx/http"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
	"html-to-gomponents/internal/adapters/services/parser"
	"html-to-gomponents/internal/app"
	"html-to-gomponents/internal/requests"
	"log"
	"net/http"
)

func createIndexPageHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		title, body := indexPage()
		return Page(title, c.Request().URL.Path, body).Render(c.Response())
	}
}

func createParseHandler(application app.Application) echo.HandlerFunc {
	return func(c echo.Context) error {
		if hxhttp.GetTrigger(c.Request().Header) == "htmlText" {
			text := c.FormValue("htmlText")
			handle, err := application.ParseHandler.Handle(c.Request().Context(), requests.Parse{Body: []byte(text)})
			if err != nil {
				if errors.Is(err, parser.ParseErr) {
					c.Logger().Error(err)
					//c.Response().Status = http.StatusBadRequest
					log.Println(err.Error())
					return result(err.Error()).Render(c.Response())
				}
				c.Logger().Error(err)
				return err
			}
			return result(handle.Body).Render(c.Response())
		}
		handle, err := application.ParseHandler.Handle(c.Request().Context(), requests.Parse{Body: []byte("")})

		if err != nil {
			if errors.Is(err, parser.ParseErr) {
				c.Logger().Error(err)
				c.Response().Status = http.StatusBadRequest
				return result(err.Error()).Render(c.Response())
			}
			c.Logger().Error(err)
			return err
		}
		return result(handle.Body).Render(c.Response())
	}
}

func indexPage() (string, g.Node) {
	return "HTML To Gomponents", Div(Class("flex flex-row justify-between grow sm:gap-0 md:gap-2 lg:gap-4"),

		Div(Class("basis-1/2 bg-gray-200 flex flex-col"), hx.Boost("true"), hx.Trigger("load"), hx.Post("/parse"), hx.Target("#result"),
			Textarea(ID("htmlText"), Class("grow  w-full  border-gray-300 bg-gray-100 align-top shadow-sm sm:text-sm"), Name("htmlText"), hx.Boost("true"), hx.Trigger("input from:#htmlText"), hx.Post("/parse"), hx.Target("#result"), Placeholder("WRITE HTML HERE")),
		),
		Div(Class("basis-1/2 bg-gray-200 flex flex-col"),
			Textarea(ID("result"), Class("grow  w-full  border-gray-300 bg-gray-100 align-top shadow-sm sm:text-sm "), ReadOnly(), Name("result"), ID("result"), hx.Swap("innerHTML"), result("")),
		),
	)
}

func result(str string) g.Node {
	return g.Text(str)
}

func Page(title, path string, body g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:    title,
		Language: "en",
		Head: []g.Node{
			Script(Src("https://cdn.tailwindcss.com?plugins=typography,forms")),
			Script(Src("htmx.min.js")),
			Link(Rel("icon"), Href("favicon.ico")),
		},
		Body: []g.Node{
			Div(Class("bg-gray-200 w-screen rounded-none"),
				Div(Class("mx-auto h-screen flex flex-col max-w-7xl sm:px-0 md:px-4 lg:px-8"),
					body,
					PageFooter(),
				),
			),
		},
	})
}

func PageFooter() g.Node {
	return Footer(Class("prose prose-sm prose-indigo max-w-none"),
		P(Class("text-center"),
			g.Text("© by "),
			A(Href("https://github.com/piotrkowalski"), g.Text("Piotr Kowalski")),
			g.Text(" 2024"),
			g.Text(". Please report bugs on "),
			A(Href("https://github.com/PiotrKowalski/html-to-gomponents"), g.Text("Github Issues")),
		))
}
