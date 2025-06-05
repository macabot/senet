package dispatch

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/app/state"
)

func resetListeners() {
	onSetSpeechBubbleKind = map[state.SpeechBubbleKind]func(s *state.State, player int){}
	onUnsetSpeechBubbleKind = map[state.SpeechBubbleKind]func(s *state.State, player int){}
	onToggleSpeechBubbleByKind = map[state.SpeechBubbleKind]func(s *state.State, player int){}
	onMoveToSquare = []func(s, newState *state.State, from, to state.Position) []hypp.Effect{}
	onNoMove = []func(s, newState *state.State) []hypp.Effect{}
	onThrowSticks = []func(s, newState *state.State) []hypp.Effect{}
}

var beforeUnloadListenerID window.EventListenerID

// addBeforeUnloadListener triggers a browser-generated confirmation dialog that asks users to confirm if they really want to leave the page when they try to close or reload it, or navigate somewhere else.
// See https://developer.mozilla.org/en-US/docs/Web/API/Window/beforeunload_event
// It will do nothing if the event listener is already registered.
// Use removeBeforeUnloadListener to remove the event listener.
func addBeforeUnloadListener() {
	if !beforeUnloadListenerID.IsUndefined() {
		return
	}
	beforeUnloadListenerID = window.AddEventListener("beforeunload", func(e window.Event) {
		e.PreventDefault()
	})
}

// removeBeforeUnloadListener removes the event listener set by addBeforeUnloadListener.
// It will do nothing if no event listener is registered.
func removeBeforeUnloadListener() {
	if beforeUnloadListenerID.IsUndefined() {
		return
	}
	window.RemoveEventListener("beforeunload", beforeUnloadListenerID)
	beforeUnloadListenerID = window.EventListenerID{}
}

func GoToTutorial(s *state.State, _ hypp.Payload) hypp.Dispatchable {
	newState := s.Clone()
	newState.Screen = state.GameScreen
	newState.Game = state.NewGame()
	newState.Game.Players[0].Name = "You"
	newState.Game.Players[1].Name = "Tutor"
	newState.Game.Turn = 1
	newState.Game.TurnMode = state.IsPlayer0
	newState.Game.Players[1].SpeechBubble = &state.SpeechBubble{
		Kind: state.TutorialStart,
	}
	newState.Game.Sticks.GeneratorKind = state.TutorialSticksGeneratorKind
	newState.TutorialIndex = 0
	resetListeners()
	RegisterTutorial()
	addBeforeUnloadListener()
	return newState
}

func GoToLocalPlayerVsPlayer(s *state.State, _ hypp.Payload) hypp.Dispatchable {
	newState := s.Clone()
	newState.Screen = state.GameScreen
	newState.Game = state.NewGame()
	newState.Game.TurnMode = state.IsBothPlayers
	addBeforeUnloadListener()
	return newState
}

func GoToStartPage(_ *state.State, _ hypp.Payload) hypp.Dispatchable {
	newState := &state.State{
		Screen: state.StartScreen,
	}
	resetListeners()
	resetSignaling(newState)
	removeBeforeUnloadListener()
	return newState
}

func GoToSignalingPage(_ *state.State, _ hypp.Payload) hypp.Dispatchable {
	newState := &state.State{
		Screen: state.SignalingScreen,
	}
	resetSignaling(newState)
	initSignaling(newState)
	addBeforeUnloadListener()
	return newState
}

func GoToOnlinePlayerVsPlayer(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	isPlayer0 := payload.(bool)
	newState := s.Clone()
	newState.Screen = state.GameScreen
	newState.Game = state.NewGame()
	if isPlayer0 {
		newState.Game.TurnMode = state.IsPlayer0
		newState.Game.Players[0].Name = "You"
		newState.Game.Players[1].Name = "Opponent"
	} else {
		newState.Game.TurnMode = state.IsPlayer1
		newState.Game.Players[0].Name = "Opponent"
		newState.Game.Players[1].Name = "You"
	}
	newState.Game.Sticks.GeneratorKind = state.CommitmentSchemeGeneratorKind
	isCaller := !newState.Game.HasTurn()
	effects := sendIsReady(newState, isCaller)
	addBeforeUnloadListener()
	return hypp.StateAndEffects[*state.State]{
		State:   newState,
		Effects: effects,
	}
}

func GoToWhoGoesFirstPage(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	isCaller := payload.(bool)
	connectionState := ""
	readyState := ""
	loading := false
	if s.Signaling != nil {
		connectionState = s.Signaling.ConnectionState
		readyState = s.Signaling.ReadyState
		loading = s.Signaling.Loading
	}
	if connectionState != "connected" || readyState != "open" || loading {
		return s
	}

	newState := s.Clone()
	newState.Screen = state.WhoGoesFirstScreen
	newState.CommitmentScheme.IsCaller = isCaller
	resetListeners()
	registerCommitmentScheme()
	effects := sendIsReady(newState, isCaller)
	addBeforeUnloadListener()
	return hypp.StateAndEffects[*state.State]{
		State:   newState,
		Effects: effects,
	}
}
