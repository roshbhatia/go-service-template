package api

import (
	"encoding/json"
	"fmt"
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
	threadWorkDone := make(chan bool)

	go func() {
		var echoReq echoRequest

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&echoReq)
		if err != nil {
			if strings.Contains(err.Error(), "json: unknown field") {
				w.WriteHeader(http.StatusBadRequest)
				h.Logger.Err(fmt.Sprintf("client %s sent malformed request body\n", r.RemoteAddr))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				h.Logger.Err(fmt.Sprintf("failed to decode request json from client %s\n", r.RemoteAddr))
			}
			return
		}

		echoResp := echo{
			TimeStamp: time.Now().UTC().String(),
			EchoStr:   echoReq.EchoStr,
		}
		h.Logger.Info(fmt.Sprintf("client %s wants '%s' echoed", r.RemoteAddr, echoReq.EchoStr))

		responseBody, err := json.Marshal(echoResp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			h.Logger.Err(fmt.Sprintf("failed to marshal health check object"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBody)
		threadWorkDone <- true
	}()

	// Not really safe per-say, but if the contex is done, we'll just abort the above thread
	select {
	case <-h.Ctx.Done():
		return
	case <-threadWorkDone:
		return
	}
}
