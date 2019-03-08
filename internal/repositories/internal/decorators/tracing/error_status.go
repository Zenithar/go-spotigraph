package tracing

import (
	"go.opencensus.io/trace"

	"go.zenithar.org/pkg/db"
)

// spanStatus sets the span status
func spanStatus(span *trace.Span, err error) {
	// Set span status
	if err != nil {
		switch err {
		case db.ErrNoResult:
			span.SetStatus(trace.Status{Code: trace.StatusCodeNotFound, Message: err.Error()})
		default:
			span.SetStatus(trace.Status{Code: trace.StatusCodeAborted, Message: err.Error()})
		}
	} else {
		span.SetStatus(trace.Status{Code: trace.StatusCodeOK})
	}
}
