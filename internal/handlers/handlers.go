package handlers

import (
	"github.com/rdforte/sequencer/internal/atomicSequence"
	"github.com/rdforte/sequencer/internal/handlers/health"
	v1 "github.com/rdforte/sequencer/internal/handlers/v1"
	"net/http"
	"net/http/pprof"
)

func CreateAPIMux() http.Handler {
	mux := http.NewServeMux()

	sequencer := atomicSequence.CreateAtomicSequencer()
	v1Handler := v1.CreateHandler(sequencer)

	mux.HandleFunc("/v1/sequence", v1Handler.Sequencer)

	return mux
}

func CreateDebugMux(buildEnv string) http.Handler {
	mux := debugStandardLibraryMux()

	healthHandler := health.CreateHandler(buildEnv)

	mux.HandleFunc("/debug/readiness", healthHandler.Readiness)
	mux.HandleFunc("/debug/liveness", healthHandler.Liveness)

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

	return mux
}
