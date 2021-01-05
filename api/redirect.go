package api

import (
	"fmt"
	"net/http"
)

func (h *Api) Redirect(w http.ResponseWriter, r *http.Request) {
	newUrl := fmt.Sprintf("https://%s%s", r.Host, r.URL.Path)
	h.Logger.Info(fmt.Sprintf("redirecting %s to %s", r.RequestURI, newUrl))
	http.Redirect(w, r, newUrl, http.StatusPermanentRedirect)
}
