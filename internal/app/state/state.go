package state

import (
	"github.com/macabot/senet/internal/pkg/scaledrone"
	"github.com/macabot/senet/internal/pkg/webrtc"
)

type State struct {
	PanicStackTrace    string
	Game               *Game
	Screen             Screen
	ShowMenu           bool
	HideOrientationTip bool

	SignalingError *SignalingError
	RoomName       string
	PeerConnection webrtc.PeerConnection
	DataChannel    webrtc.DataChannel
	Scaledrone     *scaledrone.Scaledrone

	CommitmentScheme CommitmentScheme

	TutorialIndex int
}

func (s *State) Clone() *State {
	if s == nil {
		return nil
	}
	return &State{
		PanicStackTrace:    s.PanicStackTrace,
		Game:               s.Game.Clone(),
		Screen:             s.Screen,
		ShowMenu:           s.ShowMenu,
		HideOrientationTip: s.HideOrientationTip,

		SignalingError: s.SignalingError,
		RoomName:       s.RoomName,
		PeerConnection: s.PeerConnection,
		DataChannel:    s.DataChannel,
		Scaledrone:     s.Scaledrone,

		CommitmentScheme: s.CommitmentScheme.Clone(),

		TutorialIndex: s.TutorialIndex,
	}
}
