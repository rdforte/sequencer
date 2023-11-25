package handlers

import (
	"expvar"
	"github.com/rdforte/sequencer/internal/handlers/health"
	v1 "github.com/rdforte/sequencer/internal/handlers/v1"
	"net/http"
	"net/http/pprof"
)

func CreateAPIMux(build string) http.Handler {
	mux := http.NewServeMux()

	v1Handler := v1.Handler{
		Build: build,
	}

	mux.HandleFunc("/v1/sequence", v1Handler.Sequencer)

	return mux
}

func CreateDebugMux(build string) http.Handler {
	mux := debugStandardLibraryMux()

	cgh := health.Handler{
		Build: build,
	}

	mux.HandleFunc("/debug/readiness", cgh.Readiness)
	mux.HandleFunc("/debug/liveness", cgh.Liveness)

	return mux
}

func debugStandardLibraryMux() *http.ServeMux {
	mux := http.NewServeMux()

	// Register all the standard library debug endpoints.
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/vars", expvar.Handler())

	return mux
}
