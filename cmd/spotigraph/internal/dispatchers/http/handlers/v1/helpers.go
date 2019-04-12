package v1

import (
	"net/http"
	"strconv"

	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"

	"go.zenithar.org/spotigraph/pkg/respond"
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
	if res != nil && res.GetError() != nil {
		respond.WithError(w, r, int(res.GetError().Code), res.GetError().Message)
		return true
	}
	if err != nil {
		respond.WithError(w, r, http.StatusInternalServerError, err)
		return true
	}

	return false
}
