package v1

import (
	"net/http"
	"strconv"

	"go.uber.org/zap"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/pkg/web/respond"
)

func toUint32(value string, fallback uint32) uint32 {
	v, err := strconv.ParseUint(value, 10, 32)
	if v <= 0 || err != nil {
		return fallback
	}
	return uint32(v)
}

// -----------------------------------------------------------------------------

// Error describes the spotigraph error getter
type Error interface {
	GetError() *spotigraph.Error
}

func publicError(w http.ResponseWriter, r *http.Request, res Error, err error) bool {
	log.For(r.Context()).Error("Unable to handle the request", zap.Error(err))

	if res != nil && res.GetError() != nil {
		respond.WithError(w, r, int(res.GetError().Code), res.GetError().Message)
		return true
	}
	if err != nil {
		respond.WithError(w, r, http.StatusInternalServerError, "Oups, something goes wrong during request handling !")
		return true
	}

	return false
}
