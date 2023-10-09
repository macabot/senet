package dispatch

import (
	"github.com/macabot/senet/internal/app/state"
	"github.com/macabot/senet/internal/pkg/webrtc"
)

// TODO should this be a subscription in the AppProps?
func CreatePeerConnectionOffer() {
	state.PeerConnection.SetLocalDescription(state.PeerConnection.CreateOffer())
	state.PeerConnection.SetOnICECandidate(func(pci webrtc.PeerConnectionICEEvent) {
		if pci.Truthy() {
			return
		}
		/*offer := */ state.PeerConnection.LocalDescription().SDP()
	})
}
