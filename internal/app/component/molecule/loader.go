package molecule

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
)

func Loader(hidden bool) *hypp.VNode {
	return html.Div(
		hypp.HProps{
			"class": map[string]bool{
				"loader-wrapper": true,
				"hidden":         hidden,
			},
		},
		html.Div(hypp.HProps{"class": "loader"}),
	)
}
