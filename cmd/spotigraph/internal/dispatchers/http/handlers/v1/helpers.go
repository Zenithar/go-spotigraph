package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

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
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

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
	} else {
		asJSONError(ctx, w, internalErr)
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
		w.WriteHeader(http.StatusBadRequest)

		// Return public error
		asJSON(ctx, w, map[string]interface{}{
			"@context": jsonldContext,
			"@type":    "Error",
			"code":     http.StatusBadRequest,
			"message":  "Unable to handle this request",
		})
	}
}

func toUint32(value string, fallback uint32) uint32 {
	v, err := strconv.ParseUint(value, 10, 32)
	if v <= 0 || err != nil {
		return fallback
	}
	return uint32(v)
}
