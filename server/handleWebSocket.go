package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

var upgrader = websocket.Upgrader{}

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type incomingWebsocketMessage struct {
		Action string
		Domain string
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.logger.Error("WS Upgrade:", err)
		return
	}

	defer ws.Close()

	_, inMessageString, err := ws.ReadMessage()
	if err != nil {
		s.LogError(err.Error())
		return
	}

	var inMessage incomingWebsocketMessage
	err = json.Unmarshal(inMessageString, &inMessage)

	runScript(inMessage.Action, inMessage.Domain, ws)
}
