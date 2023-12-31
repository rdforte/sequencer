package health

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func CreateHandler(buildEnv string) handler {
	return handler{buildEnv}
}

type handler struct {
	BuildEnv string
}

func (h handler) Readiness(w http.ResponseWriter, r *http.Request) {
	status := "ok"
	statusCode := http.StatusOK

	data := struct {
		Status string `json:"status"`
	}{
		Status: status,
	}

	if err := response(w, statusCode, data); err != nil {
		log.Printf("readiness ERROR %v", err)
	}

	log.Printf("readiness: statusCode %v, method %v, path %v, remoteAddress %v", statusCode, r.Method, r.URL.Path, r.RemoteAddr)
}

func (h handler) Liveness(w http.ResponseWriter, r *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		host = "unavailable"
	}

	data := struct {
		Status string `json:"status,omitempty"`
		Build  string `json:"build,omitempty"`
		Host   string `json:"host,omitempty"`
	}{
		Status: "up",
		Build:  h.BuildEnv,
		Host:   host,
	}

	statusCode := http.StatusOK
	if err := response(w, statusCode, data); err != nil {
		log.Printf("liveness ERROR %v", err)
	}

	log.Printf("liveness: statusCode %v, method %v, path %v, remoteAddress %v", statusCode, r.Method, r.URL.Path, r.RemoteAddr)
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
