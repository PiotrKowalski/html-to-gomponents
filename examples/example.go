package example

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func example() g.Node {
	return Body(
		Div(Class("overflow-x-auto"),
			Table(Class("min-w-full divide-y-2 divide-gray-200 bg-white text-sm"),
				THead(Class("ltr:text-left rtl:text-right"),
					Tr(
						Th(Class("px-4 py-2"),
							Label(For("SelectAll"), Class("sr-only")),
							Input(Type("checkbox"), ID("SelectAll"), Class("size-5 rounded border-gray-300")),
						),
						Th(Class("whitespace-nowrap px-4 py-2 font-medium text-gray-900")),
						Th(Class("whitespace-nowrap px-4 py-2 font-medium text-gray-900")),
						Th(Class("whitespace-nowrap px-4 py-2 font-medium text-gray-900")),
						Th(Class("whitespace-nowrap px-4 py-2 font-medium text-gray-900")),
					),
				),
				TBody(Class("divide-y divide-gray-200"),
					Tr(
						Td(Class("px-4 py-2"),
							Label(Class("sr-only"), For("Row1")),
							Input(Class("size-5 rounded border-gray-300"), Type("checkbox"), ID("Row1")),
						),
						Td(Class("whitespace-nowrap px-4 py-2 font-medium text-gray-900")),
						Td(Class("whitespace-nowrap px-4 py-2 text-gray-700")),
						Td(Class("whitespace-nowrap px-4 py-2 text-gray-700")),
						Td(Class("whitespace-nowrap px-4 py-2 text-gray-700")),
					),
					Tr(
						Td(Class("px-4 py-2"),
							Label(For("Row2"), Class("sr-only")),
							Input(Type("checkbox"), ID("Row2"), Class("size-5 rounded border-gray-300")),
						),
						Td(Class("whitespace-nowrap px-4 py-2 font-medium text-gray-900")),
						Td(Class("whitespace-nowrap px-4 py-2 text-gray-700")),
						Td(Class("whitespace-nowrap px-4 py-2 text-gray-700")),
						Td(Class("whitespace-nowrap px-4 py-2 text-gray-700")),
					),
					Tr(
						Td(Class("px-4 py-2"),
							Label(Class("sr-only"), For("Row3")),
							Input(ID("Row3"), Class("size-5 rounded border-gray-300"), Type("checkbox")),
						),
						Td(Class("whitespace-nowrap px-4 py-2 font-medium text-gray-900")),
						Td(Class("whitespace-nowrap px-4 py-2 text-gray-700")),
						Td(Class("whitespace-nowrap px-4 py-2 text-gray-700")),
						Td(Class("whitespace-nowrap px-4 py-2 text-gray-700")),
					),
				),
			),
		),
	)
}
