package v1

import (
	"encoding/json"
	"github.com/rdforte/sequencer/internal/atomicSequence"
	"log"
	"net/http"
)

func CreateHandler(sequencer atomicSequence.Sequencer) handler {
	return handler{
		sequencer: sequencer,
	}
}

type handler struct {
	sequencer atomicSequence.Sequencer
}

func (h handler) Sequencer(w http.ResponseWriter, r *http.Request) {
	status := "ok"
	statusCode := http.StatusOK

	seqNumber := h.sequencer.Increment()

	data := struct {
		Status   string   `json:"status"`
		Sequence sequence `json:"sequence"`
	}{
		Status: status,
		Sequence: sequence{
			Number: seqNumber,
		},
	}

	if err := response(w, statusCode, data); err != nil {
		log.Printf("readiness ERROR %v", err)
	}
}

type sequence struct {
	Number int64 `json:"number"`
}

func response(w http.ResponseWriter, statusCode int, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}
