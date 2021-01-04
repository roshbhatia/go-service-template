package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type echoRequest struct {
	EchoStr string
}

type echo struct {
	TimeStamp, EchoStr string
}

func (h *Api) EchoHandler(w http.ResponseWriter, r *http.Request) {
	var echoReq echoRequest

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&echoReq)
	if err != nil {
		if strings.Contains(err.Error(), "json: unknown field") {
			w.WriteHeader(http.StatusBadRequest)
			h.Logger.Err.Printf("client %s sent malformed request body\n", r.RemoteAddr)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			h.Logger.Err.Printf("failed to decode request json from client %s\n", r.RemoteAddr)
		}
		return
	}

	echoResp := echo{
		TimeStamp: time.Now().UTC().String(),
		EchoStr:   echoReq.EchoStr,
	}
	h.Logger.Info.Printf("client %s wants '%s' echoed", r.RemoteAddr, echoReq.EchoStr)

	responseBody, err := json.Marshal(echoResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.Logger.Err.Println("failed to marshal health check object")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
