package molecule

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
)

func SignalingErrors(signalingErrors []state.SignalingError) *hypp.VNode {
	if len(signalingErrors) == 0 {
		return nil
	}

	var children []*hypp.VNode
	for _, e := range signalingErrors {
		children = append(children, signalingError(e))
	}
	return html.Section(
		nil,
		children...,
	)
}

func signalingError(signalingError state.SignalingError) *hypp.VNode {
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
