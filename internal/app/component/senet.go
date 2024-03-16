package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
)

func Senet(s *state.State) *hypp.VNode {
	return html.Html(
		hypp.HProps{"lang": "en"},
		head(),
		body(s),
	)
}

func head() *hypp.VNode {
	return html.Head(
		nil,
		html.Link(hypp.HProps{
			"rel":  "stylesheet",
			"href": "/senet.css",
		}),
		html.Meta(hypp.HProps{"charset": "utf-8"}),
		html.Meta(hypp.HProps{
			"name":    "viewport",
			"content": "width=device-width, initial-scale=1.0",
		}),
		html.Title(nil, hypp.Text("Senet")),
	)
}

func body(s *state.State) *hypp.VNode {
	if s.PanicStackTrace != nil {
		return PanicModal(s)
	}

	var page *hypp.VNode
	switch s.Page {
	case state.StartPage:
		page = StartPage(s)
	case state.SignalingPage:
		page = SignalingPage(s)
	case state.WhoGoesFirstPage:
		page = WhoGoesFirstPage(s)
	case state.GamePage:
		page = GamePage(s)
	case state.RulesPage:
		page = RulesPage()
	default:
		panic(fmt.Errorf("component not implemented for page %d", s.Page))
	}

	var menu *hypp.VNode
	if s.ShowMenu {
		menu = Menu()
	}

	return html.Body(nil, page, menu)
}
