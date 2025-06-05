package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/component/screen"
	"github.com/macabot/senet/internal/app/state"
)

func Senet(s *state.State) *hypp.VNode {
	return html.Html(
		hypp.HProps{"lang": "en"},
		head(s.Screen),
		body(s),
	)
}

func head(page state.Screen) *hypp.VNode {
	var title string
	switch page {
	case state.HomePage:
		title = "Senet"
	case state.RulesPage:
		title = "Senet Rules"
	default:
		title = "Play Senet"
	}

	return html.Head(
		nil,
		html.Meta(hypp.HProps{"charset": "utf-8"}),
		html.Meta(hypp.HProps{
			"name":    "viewport",
			"content": "width=device-width, initial-scale=1.0",
		}),
		html.Title(nil, hypp.Text(title)),
		html.Link(hypp.HProps{
			"rel":  "stylesheet",
			"href": "/senet.css",
		}),
	)
}

func body(s *state.State) *hypp.VNode {
	if s.PanicStackTrace != nil {
		return PanicModal(s)
	}

	var page *hypp.VNode
	switch s.Screen {
	case state.StartScreen:
		page = StartPage(s)
	case state.OnlineScreen:
		page = screen.Online()
	case state.WhoGoesFirstScreen:
		page = WhoGoesFirstPage(s)
	case state.GameScreen:
		page = GamePage(s)
	case state.HomePage:
		page = HomePage()
	case state.RulesPage:
		page = RulesPage()
	default:
		panic(fmt.Errorf("component not implemented for page %d", s.Screen))
	}

	var menu *hypp.VNode
	if s.ShowMenu {
		menu = Menu()
	}

	return html.Body(nil, page, menu)
}
