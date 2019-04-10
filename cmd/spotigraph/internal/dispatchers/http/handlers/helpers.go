package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

func asJSON(ctx context.Context, w http.ResponseWriter, response interface{}) {
	// Marshal response as json
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set conteent type header
	w.Header().Set("Content-Type", "application/json")

	// Write response
	_, err = w.Write(js)
	log.CheckErrCtx(ctx, "Unable to write response", err)
}

func asJSONStatus(ctx context.Context, w http.ResponseWriter, code int, message string) {
	// Write HTTP status
	w.WriteHeader(code)

	// Encode payload
	asJSON(ctx, w, map[string]interface{}{
		"@context": jsonldContext,
		"@type":    "Status",
		"code":     code,
		"message":  message,
	})
}

func asJSONResultError(ctx context.Context, w http.ResponseWriter, publicErr *spotigraph.Error, internalErr error) {
	if publicErr != nil {
		asJSONError(ctx, w, publicErr)
	}
}

func asJSONError(ctx context.Context, w http.ResponseWriter, errObj interface{}) {
	switch err := errObj.(type) {
	case *spotigraph.Error:
		// Write HTTP status
		w.WriteHeader(int(err.Code))

		// Encode payload
		asJSON(ctx, w, map[string]interface{}{
			"@context": jsonldContext,
			"@type":    "Error",
			"code":     err.Code,
			"message":  err.Message,
		})
	default:
		// Write HTTP status
		w.WriteHeader(http.StatusInternalServerError)

		// Return public error
		asJSON(ctx, w, map[string]interface{}{
			"@context": jsonldContext,
			"@type":    "Error",
			"code":     http.StatusInternalServerError,
			"message":  "Unable to handle this request",
		})
	}

}
