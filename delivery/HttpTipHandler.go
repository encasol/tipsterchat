package delivery

import (
	"encoding/json"
	"github.com/encasol/tipsterchat/model"
	"github.com/encasol/tipsterchat/service"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type HttpTipHandler struct {
	TipService service.ITipService
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
		decoder := json.NewDecoder(r.Body)
		var tip model.Tip
		err := decoder.Decode(&tip)
		if err != nil {
			panic(err)
		}

		h.TipService.AddTip(tip)
	}

	tips, err := h.TipService.ListTips()
	if err != nil {
		panic(err)
	}

	js, err := json.Marshal(tips)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
