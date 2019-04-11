package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type tribeCtrl struct {
	tribes services.Tribe
}

// -----------------------------------------------------------------------------

// TribeRoutes returns tribe management related API
func TribeRoutes(tribes services.Tribe) chi.Router {
	r := chi.NewRouter()

	// Initialize controller
	ctrl := &tribeCtrl{
		tribes: tribes,
	}

	// Map routes
	r.Get("/", ctrl.search())
	r.Post("/", ctrl.create())

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", ctrl.read())
		r.Put("/", ctrl.update())
		r.Delete("/", ctrl.delete())
	})

	// Return router
	return r
}

// -----------------------------------------------------------------------------

func (c *tribeCtrl) create() http.HandlerFunc {
	// Request type
	var request spotigraph.TribeCreateReq

	// Response type
	type response struct {
		Context string                   `json:"@context"`
		Type    string                   `json:"@type"`
		ID      string                   `json:"@id"`
		Entity  *spotigraph.Domain_Tribe `json:",inline"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Decode request as json
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			asJSONError(ctx, w, err)
			return
		}

		// Delegate to service
		res, err := c.tribes.Create(ctx, &request)
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSON(ctx, w, &response{
			Context: jsonldContext,
			Type:    "Tribe",
			ID:      fmt.Sprintf("/tribes/%s", res.Entity.Id),
			Entity:  res.Entity,
		})
	}
}

func (c *tribeCtrl) read() http.HandlerFunc {
	// Response type
	type response struct {
		Context string                   `json:"@context"`
		Type    string                   `json:"@type"`
		ID      string                   `json:"@id"`
		Entity  *spotigraph.Domain_Tribe `json:",inline"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Delegate to service
		res, err := c.tribes.Get(ctx, &spotigraph.TribeGetReq{
			Id: chi.URLParamFromCtx(ctx, "id"),
		})
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSON(ctx, w, &response{
			Context: jsonldContext,
			Type:    "Tribe",
			ID:      fmt.Sprintf("/tribes/%s", res.Entity.Id),
			Entity:  res.Entity,
		})
	}
}

func (c *tribeCtrl) update() http.HandlerFunc {
	// Request type
	var request spotigraph.TribeUpdateReq

	// Response type
	type response struct {
		Context string                   `json:"@context"`
		Type    string                   `json:"@type"`
		ID      string                   `json:"@id"`
		Entity  *spotigraph.Domain_Tribe `json:",inline"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Decode request as json
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			asJSONError(ctx, w, err)
			return
		}

		// Delegate to service
		res, err := c.tribes.Update(ctx, &request)
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSON(ctx, w, &response{
			Context: jsonldContext,
			Type:    "Tribe",
			ID:      fmt.Sprintf("/tribes/%s", res.Entity.Id),
			Entity:  res.Entity,
		})
	}
}

func (c *tribeCtrl) delete() http.HandlerFunc {
	// Response type
	type response struct {
		Context string                   `json:"@context"`
		Type    string                   `json:"@type"`
		ID      string                   `json:"@id"`
		Entity  *spotigraph.Domain_Tribe `json:",inline"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Delegate to service
		res, err := c.tribes.Delete(ctx, &spotigraph.TribeGetReq{
			Id: chi.URLParamFromCtx(ctx, "id"),
		})
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSONStatus(ctx, w, http.StatusOK, "Tribe successfully deleted.")
	}
}

func (c *tribeCtrl) search() http.HandlerFunc {
	// Request type
	var request spotigraph.TribeSearchReq

	// Response type
	type response struct {
		Context string                        `json:"@context"`
		Type    string                        `json:"@type"`
		ID      string                        `json:"@id"`
		Page    *spotigraph.PaginatedTribeRes `json:",inline"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Decode request as json
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			asJSONError(ctx, w, err)
			return
		}

		// Delegate to service
		res, err := c.tribes.Search(ctx, &request)
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSON(ctx, w, &response{
			Context: jsonldContext,
			Type:    "Collection",
			ID:      r.RequestURI,
			Page:    res,
		})
	}
}
