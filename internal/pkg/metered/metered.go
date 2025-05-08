package metered

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/macabot/hypp/js"
	"github.com/macabot/senet/internal/pkg/promise"
	"github.com/macabot/senet/internal/pkg/webrtc"
)

// Set these variables when building using the -ldflags.
var (
	METERED_APP_NAME string
	METERED_API_KEY  string
)

var (
	FetchedICEServers []webrtc.ICEServer
	FetchErr          error
)

func init() {
	FetchedICEServers, FetchErr = fetchICEServers()
}

// fetchICEServers fetches ICE servers from your Metered app.
//
// See https://www.metered.ca/tools/openrelay/#-how-to-use
func fetchICEServers() ([]webrtc.ICEServer, error) {
	if METERED_APP_NAME == "" {
		return nil, errors.New("embedded variable METERED_APP_NAME must be set")
	}
	if METERED_API_KEY == "" {
		return nil, errors.New("embedded variable METERED_API_KEY must be set")
	}

	url := fmt.Sprintf("https://%s.metered.live/api/v1/turn/credentials?apiKey=%s", METERED_APP_NAME, METERED_API_KEY)
	responsePromise := js.Global().Call("fetch", url)
	response, err := promise.Await(responsePromise)
	if err != nil {
		return nil, err
	}
	textPromise := response.Call("text")
	text, err := promise.Await(textPromise)
	if err != nil {
		return nil, err
	}
	var iceServers []webrtc.ICEServer
	if err := json.Unmarshal([]byte(text.String()), &iceServers); err != nil {
		return nil, err
	}
	return iceServers, nil
}
