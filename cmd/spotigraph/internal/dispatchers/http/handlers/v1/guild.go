package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type guildCtrl struct {
	guilds services.Guild
}

// -----------------------------------------------------------------------------

// GuildRoutes returns guild management related API
func GuildRoutes(guilds services.Guild) http.Handler {
	r := chi.NewRouter()

	// Initialize controller
	ctrl := &guildCtrl{
		guilds: guilds,
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

func (c *guildCtrl) create() http.HandlerFunc {
	// Request type
	var request spotigraph.GuildCreateReq

	// Response type
	type response struct {
		Context string                   `json:"@context"`
		Type    string                   `json:"@type"`
		ID      string                   `json:"@id"`
		Entity  *spotigraph.Domain_Guild `json:",inline"`
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
		res, err := c.guilds.Create(ctx, &request)
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSON(ctx, w, &response{
			Context: jsonldContext,
			Type:    "Guild",
			ID:      fmt.Sprintf("/guilds/%s", res.Entity.Id),
			Entity:  res.Entity,
		})
	}
}

func (c *guildCtrl) read() http.HandlerFunc {
	// Response type
	type response struct {
		Context string                   `json:"@context"`
		Type    string                   `json:"@type"`
		ID      string                   `json:"@id"`
		Entity  *spotigraph.Domain_Guild `json:",inline"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Delegate to service
		res, err := c.guilds.Get(ctx, &spotigraph.GuildGetReq{
			Id: chi.URLParamFromCtx(ctx, "id"),
		})
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSON(ctx, w, &response{
			Context: jsonldContext,
			Type:    "Guild",
			ID:      fmt.Sprintf("/guilds/%s", res.Entity.Id),
			Entity:  res.Entity,
		})
	}
}

func (c *guildCtrl) update() http.HandlerFunc {
	// Request type
	var request spotigraph.GuildUpdateReq

	// Response type
	type response struct {
		Context string                   `json:"@context"`
		Type    string                   `json:"@type"`
		ID      string                   `json:"@id"`
		Entity  *spotigraph.Domain_Guild `json:",inline"`
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
		res, err := c.guilds.Update(ctx, &request)
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSON(ctx, w, &response{
			Context: jsonldContext,
			Type:    "Guild",
			ID:      fmt.Sprintf("/guilds/%s", res.Entity.Id),
			Entity:  res.Entity,
		})
	}
}

func (c *guildCtrl) delete() http.HandlerFunc {
	// Response type
	type response struct {
		Context string                   `json:"@context"`
		Type    string                   `json:"@type"`
		ID      string                   `json:"@id"`
		Entity  *spotigraph.Domain_Guild `json:",inline"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Delegate to service
		res, err := c.guilds.Delete(ctx, &spotigraph.GuildGetReq{
			Id: chi.URLParamFromCtx(ctx, "id"),
		})
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSONStatus(ctx, w, http.StatusOK, "Guild successfully deleted.")
	}
}

func (c *guildCtrl) search() http.HandlerFunc {
	// Request type
	var request spotigraph.GuildSearchReq

	// Response type
	type response struct {
		Context string                        `json:"@context"`
		Type    string                        `json:"@type"`
		ID      string                        `json:"@id"`
		Page    *spotigraph.PaginatedGuildRes `json:",inline"`
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
		res, err := c.guilds.Search(ctx, &request)
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
