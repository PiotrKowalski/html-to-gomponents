package view

import (
	"github.com/labstack/echo/v4"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	hxhttp "github.com/maragudk/gomponents-htmx/http"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
	"html-to-gomponents/internal/app"
	"html-to-gomponents/internal/requests"
	"log"
)

func createIndexPageHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		title, body := indexPage()
		return Page(title, c.Request().URL.Path, body).Render(c.Response())
	}
}

func createParseHandler(application app.Application) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println(hxhttp.GetTrigger(c.Request().Header) == "htmlText")

		if hxhttp.GetTrigger(c.Request().Header) == "htmlText" {
			text := c.FormValue("htmlText")
			handle, err := application.ParseHandler.Handle(c.Request().Context(), requests.Parse{Body: []byte(text)})
			if err != nil {
				return err
			}

			return result(handle.Body).Render(c.Response())
		}
		return nil
	}
}

func indexPage() (string, g.Node) {
	return "Welcome!", Div(Class("flex flex-row justify-center  h-screen grow"),

		Div(Class("basis-5/12 bg-gray-200 flex flex-col"),
			Textarea(ID("htmlText"), Class("grow  w-full  border-gray-300 bg-gray-100 align-top shadow-sm sm:text-sm"), Name("htmlText"), hx.Boost("true"), hx.Trigger("input from:#htmlText"), hx.Post("/parse"), hx.Target("#result"), Placeholder("Write HTML HERE")),
		),
		Div(Class("basis-1/12 bg-gray-200")),
		Div(Class("basis-5/12 bg-gray-200 flex flex-col"),
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
			Script(Src("https://unpkg.com/htmx.org")),
		},
		Body: []g.Node{
			Class("h-screen bg-gray-200"),
			Container(
				Div(Class("flex flex-col"),
					body,
				),
			),
		},
	})
}

func Container(children ...g.Node) g.Node {
	return Div(Class("max-w-7xl mx-auto px-2 sm:px-6 lg:px-8"), g.Group(children))
}
