package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func authRoute(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		sendResponse(w, "invalid body", http.StatusBadRequest)
		return
	}

	var payload map[string]string
	err = json.Unmarshal(data, &payload)
	if err != nil {
		log.Println(err)
		sendResponse(w, "invalid json", http.StatusBadRequest)
		return
	}

	askedSecret := payload["Secret"]
	askedToken := globalService.config.generateToken(askedSecret)

	storedToken := globalService.config.getToken()

	if askedToken != storedToken {
		sendResponse(w, "invalid secret", http.StatusUnauthorized)
		return
	}

	sendResponse(w, storedToken, http.StatusOK)

}
