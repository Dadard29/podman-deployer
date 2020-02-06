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
		log.Fatalln(err)
	}

	var payload map[string]string
	err = json.Unmarshal(data, &payload)
	if err != nil {
		log.Println(err)
	}

	askedSecret := payload["Secret"]
	askedToken := globalService.config.generateToken(askedSecret)

	storedToken := globalService.config.getToken()

	if askedToken != storedToken {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if _, err := w.Write([]byte(storedToken)); err != nil {
		log.Println(err)
	}

}
