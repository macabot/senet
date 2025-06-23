package molecule

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
)

func SignalingError(signalingError *state.SignalingError) *hypp.VNode {
	if signalingError == nil {
		return nil
	}

	return html.Div(
		nil,
		html.P(
			hypp.HProps{"class": "error"},
			hypp.Text(signalingError.Summary),
		),
		html.Details(
			hypp.HProps{"class": "error"},
			html.Summary(nil, hypp.Text("Details")),
			html.Pre(nil, hypp.Textf(
				"Description:\n%s\n\nError:\n%s",
				signalingError.Description,
				signalingError.Error(),
			)),
		),
	)
}
