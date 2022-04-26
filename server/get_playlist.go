package server

import (
	"errors"
	"net/http"

	"github.com/dfroberg/m3ufilter/logger"
	"github.com/dfroberg/m3ufilter/writer"
)

func getPlaylist(state *httpState, w http.ResponseWriter, r *http.Request) error {
	if r.Method != "HEAD" && r.Method != "GET" {
		logger.Get().Errorf("Method %s is not allowed", r.Method)
		err := errors.New(http.StatusText(http.StatusMethodNotAllowed))
		return StatusError{Code: http.StatusMethodNotAllowed, Err: err}
	}

	w.Header().Set("Content-Type", "audio/mpegurl")

	writer.WriteOutput(state.appConfig.Core.Output, w, *state.playlists)
	return nil
}
