package molecule

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
)

func SenetHeader() *hypp.VNode {
	return html.Header(
		hypp.HProps{"class": "senet-header"},
		html.H1(nil, hypp.Text("Senet")),
		html.P(
			nil,
			html.Span(hypp.HProps{"class": "by"}, hypp.Text("by ")),
			html.Span(hypp.HProps{"class": "macabot"}, hypp.Text("macabot")),
		),
	)
}
