package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type squadCtrl struct {
	squads services.Squad
}

// -----------------------------------------------------------------------------

// SquadRoutes returns squad management related API
func SquadRoutes(squads services.Squad) chi.Router {
	r := chi.NewRouter()

	// Initialize controller
	ctrl := &squadCtrl{
		squads: squads,
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

func (c *squadCtrl) create() http.HandlerFunc {
	// Request type
	var request spotigraph.SquadCreateReq

	// Response type
	type response struct {
		Context string                   `json:"@context"`
		Type    string                   `json:"@type"`
		ID      string                   `json:"@id"`
		Entity  *spotigraph.Domain_Squad `json:",inline"`
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
		res, err := c.squads.Create(ctx, &request)
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSON(ctx, w, &response{
			Context: jsonldContext,
			Type:    "Squad",
			ID:      fmt.Sprintf("/squads/%s", res.Entity.Id),
			Entity:  res.Entity,
		})
	}
}

func (c *squadCtrl) read() http.HandlerFunc {
	// Response type
	type response struct {
		Context string                   `json:"@context"`
		Type    string                   `json:"@type"`
		ID      string                   `json:"@id"`
		Entity  *spotigraph.Domain_Squad `json:",inline"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Delegate to service
		res, err := c.squads.Get(ctx, &spotigraph.SquadGetReq{
			Id: chi.URLParamFromCtx(ctx, "id"),
		})
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSON(ctx, w, &response{
			Context: jsonldContext,
			Type:    "Squad",
			ID:      fmt.Sprintf("/squads/%s", res.Entity.Id),
			Entity:  res.Entity,
		})
	}
}

func (c *squadCtrl) update() http.HandlerFunc {
	// Request type
	var request spotigraph.SquadUpdateReq

	// Response type
	type response struct {
		Context string                   `json:"@context"`
		Type    string                   `json:"@type"`
		ID      string                   `json:"@id"`
		Entity  *spotigraph.Domain_Squad `json:",inline"`
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
		res, err := c.squads.Update(ctx, &request)
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSON(ctx, w, &response{
			Context: jsonldContext,
			Type:    "Squad",
			ID:      fmt.Sprintf("/squads/%s", res.Entity.Id),
			Entity:  res.Entity,
		})
	}
}

func (c *squadCtrl) delete() http.HandlerFunc {
	// Response type
	type response struct {
		Context string                   `json:"@context"`
		Type    string                   `json:"@type"`
		ID      string                   `json:"@id"`
		Entity  *spotigraph.Domain_Squad `json:",inline"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Delegate to service
		res, err := c.squads.Delete(ctx, &spotigraph.SquadGetReq{
			Id: chi.URLParamFromCtx(ctx, "id"),
		})
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSONStatus(ctx, w, http.StatusOK, "Squad successfully deleted.")
	}
}

func (c *squadCtrl) search() http.HandlerFunc {
	// Request type
	var request spotigraph.SquadSearchReq

	// Response type
	type response struct {
		Context string                        `json:"@context"`
		Type    string                        `json:"@type"`
		ID      string                        `json:"@id"`
		Page    *spotigraph.PaginatedSquadRes `json:",inline"`
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
		res, err := c.squads.Search(ctx, &request)
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
