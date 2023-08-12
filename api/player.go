package api

import "github.com/pion/webrtc/v3"

type PlayerRequestParam struct {
	PlayerName  string
	OfferString webrtc.SessionDescription
}
