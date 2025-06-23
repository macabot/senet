package page

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/component/screen"
	"github.com/macabot/senet/internal/app/component/template"
	"github.com/macabot/senet/internal/app/state"
)

func Play(s *state.State) *hypp.VNode {
	return html.Html(
		hypp.HProps{"lang": "en"},
		template.Head("Play Senet"),
		html.Body(nil, playScreen(s)),
	)
}

func playScreen(s *state.State) *hypp.VNode {
	if s.PanicStackTrace != "" {
		return screen.Panic(s)
	}

	var screenNode *hypp.VNode
	switch s.Screen {
	case state.StartScreen:
		screenNode = screen.Start()
	case state.OnlineScreen:
		screenNode = screen.Online()
	case state.NewGameScreen:
		screenNode = screen.NewGame(s)
	case state.JoinGameScreen:
		screenNode = screen.JoinGame(s)
	case state.WhoGoesFirstScreen:
		screenNode = screen.WhoGoesFirst(s)
	case state.GameScreen:
		screenNode = screen.Game(s)
	default:
		panic(fmt.Errorf("component not implemented for screen %s", s.Screen))
	}

	return screenNode
}
