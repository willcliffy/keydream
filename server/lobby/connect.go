package lobby

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/willcliffy/keydream-server/common"
	lobby_models "github.com/willcliffy/keydream-server/lobby/models"
)

func (l *LobbyHandler) ConnectHandler(w http.ResponseWriter, r *http.Request) {
	// TODO - add a dependency to the repo here to try out that golang http client
	bodyBytes := new(bytes.Buffer)
	if _, err := bodyBytes.ReadFrom(r.Body); err != nil {
		http.Error(w, "unable to read request body", http.StatusBadRequest)
		return
	}

	var connectRequest lobby_models.ConnectRequest
	if err := json.Unmarshal(bodyBytes.Bytes(), &connectRequest); err != nil {
		http.Error(w, "unable to unmarshal request body", http.StatusBadRequest)
		return
	}

	log.Printf("User %s connecting\n", connectRequest.Name)

	var connectResponse lobby_models.ConnectResponse
	for _, server := range l.Worlds {
		connectResponse.Worlds = append(connectResponse.Worlds, server)
	}

	if len(connectResponse.Worlds) == 0 {
		http.Error(w, "no worlds available", http.StatusNotFound)
		return
	}

	common.SendHTTPResponse(w, connectResponse, http.StatusOK)
}

func (l *LobbyHandler) JoinHandler(w http.ResponseWriter, r *http.Request) {
	// TODO - add a dependency to the repo here to try out that golang http client
	var bodyBytes bytes.Buffer
	if _, err := bodyBytes.ReadFrom(r.Body); err != nil {
		http.Error(w, "unable to read request body", http.StatusBadRequest)
		return
	}

	var joinRequest lobby_models.JoinRequest
	if err := json.Unmarshal(bodyBytes.Bytes(), &joinRequest); err != nil {
		http.Error(w, "unable to unmarshal request body", http.StatusBadRequest)
		return
	}

	// TODO - allow more than one world here
	if joinRequest.WorldID != 1 {
		http.Error(w, "we only support world 1 right now", http.StatusBadRequest)
		return
	}

	var joinResponse lobby_models.JoinResponse

	// TODO - Set up a world queue here that issues tokens that the game servers can
	// somehow verify on join to avoid overflow and unwanted players joining
	joinResponse.Token = "token"

	common.SendHTTPResponse(w, joinResponse, http.StatusOK)
}
