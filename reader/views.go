package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	bytes, err := w.Write(response)
	if err != nil {
		log.Printf("writing response failed: %s", err)
	}
	log.Printf("response bytes %d", bytes)
}

func respondWithText(w http.ResponseWriter, code int, payload string) {
	response := []byte(fmt.Sprintf("<pre>%s</pre>", payload))
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(code)
	bytes, err := w.Write(response)
	if err != nil {
		log.Printf("writing response failed: %s", err)
	}
	log.Printf("response bytes %d", bytes)
}

func (a *App) getHealth(w http.ResponseWriter, _ *http.Request) {
	url := fmt.Sprintf("http://%s/health", os.Getenv("APP_PING_PONG_HOST"))
	response, err := http.Get(url)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if response.StatusCode == 200 {
		respondWithJSON(w, http.StatusOK, map[string]string{"message": "ok"})
		return
	} else {
		respondWithError(w, http.StatusInternalServerError, "pingpong service not responding")
		return
	}

}

func (a *App) getRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("getting request from %s", r.RemoteAddr)
	var uuids []string
	if fileExists(os.Getenv("APP_INPUT_FILE")) {
		uuids = readUUIDs(os.Getenv("APP_INPUT_FILE"))
	} else {
		log.Printf("opening input file failed: %s", os.Getenv("APP_INPUT_FILE"))
	}
	pingpong, err := readPingPongService()
	if err != nil {
		respondWithText(w, http.StatusOK, fmt.Sprintf("%s\n%s", os.Getenv("MESSAGE"), lastRow(uuids)))
	}

	respondWithText(w, http.StatusOK, fmt.Sprintf("%s\n%s%s", os.Getenv("MESSAGE"), lastRow(uuids), pingpong))
}
