package delivery

import (
	"github.com/encasol/tipsterchat/service"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type HttpTipHandler struct {
	TipService service.AbstractTipService
	Json       JsonProxy
}

func (h HttpTipHandler) ListenAndServe(host string, port int) error {
	r := mux.NewRouter()

	r.HandleFunc("/tip", h.ListTipHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return srv.ListenAndServe()
}

func (h HttpTipHandler) ListTipHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		tip, _ := h.Json.DecodeJson(r.Body)
		h.TipService.AddTip(tip)
	}

	tips, err := h.TipService.ListTips()
	if err != nil {
		panic(err)
	}

	js, err := h.Json.EncodeJson(tips)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
