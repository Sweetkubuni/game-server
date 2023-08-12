package api

import (
	"encoding/json"
	"fmt"
	"game-server/internal"
	"net/http"
	"time"

	"github.com/pion/webrtc/v3"

	"github.com/julienschmidt/httprouter"
)

var peerConnection *webrtc.PeerConnection //nolint

func ListPlayer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func GeneratePlayer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error

	if peerConnection == nil {
		if peerConnection, err = webrtc.NewPeerConnection(webrtc.Configuration{}); err != nil {
			panic(err)
		}

		// Set the handler for ICE connection state
		// This will notify you when the peer has connected/disconnected
		peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
			fmt.Printf("ICE Connection State has changed: %s\n", connectionState.String())
		})

		// Send the current time via a DataChannel to the remote peer every 3 seconds
		peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
			d.OnOpen(func() {
				for range time.Tick(time.Second * 3) {
					if err = d.SendText(time.Now().String()); err != nil {
						panic(err)
					}
				}
			})
		})
	}

	var player PlayerRequestParam
	if err = json.NewDecoder(r.Body).Decode(&player); err != nil {
		panic(err)
	}

	if err = peerConnection.SetRemoteDescription(player.OfferString); err != nil {
		panic(err)
	}

	offerstr, err := json.Marshal(player.OfferString)
    if err != nil {
        fmt.Printf("Error: %s", err)
        return;
    }
	internal.SetRedis(player.PlayerName, string(offerstr))

	// Create channel that is blocked until ICE Gathering is complete
	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	} else if err = peerConnection.SetLocalDescription(answer); err != nil {
		panic(err)
	}

	// Block until ICE Gathering is complete, disabling trickle ICE
	// we do this because we only can exchange one signaling message
	// in a production application you should exchange ICE Candidates via OnICECandidate
	<-gatherComplete

	response, err := json.Marshal(*peerConnection.LocalDescription())
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(response); err != nil {
		panic(err)
	}

}
