package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type chapterCtrl struct {
	chapters services.Chapter
}

// -----------------------------------------------------------------------------

// ChapterRoutes returns chapter management related API
func ChapterRoutes(chapters services.Chapter) chi.Router {
	r := chi.NewRouter()

	// Initialize controller
	ctrl := &chapterCtrl{
		chapters: chapters,
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

func (c *chapterCtrl) create() http.HandlerFunc {
	// Request type
	var request spotigraph.ChapterCreateReq

	// Response type
	type response struct {
		Context string                     `json:"@context"`
		Type    string                     `json:"@type"`
		ID      string                     `json:"@id"`
		Entity  *spotigraph.Domain_Chapter `json:",inline"`
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
		res, err := c.chapters.Create(ctx, &request)
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSON(ctx, w, &response{
			Context: jsonldContext,
			Type:    "Chapter",
			ID:      fmt.Sprintf("/chapters/%s", res.Entity.Id),
			Entity:  res.Entity,
		})
	}
}

func (c *chapterCtrl) read() http.HandlerFunc {
	// Response type
	type response struct {
		Context string                     `json:"@context"`
		Type    string                     `json:"@type"`
		ID      string                     `json:"@id"`
		Entity  *spotigraph.Domain_Chapter `json:",inline"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Delegate to service
		res, err := c.chapters.Get(ctx, &spotigraph.ChapterGetReq{
			Id: chi.URLParamFromCtx(ctx, "id"),
		})
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSON(ctx, w, &response{
			Context: jsonldContext,
			Type:    "Chapter",
			ID:      fmt.Sprintf("/chapters/%s", res.Entity.Id),
			Entity:  res.Entity,
		})
	}
}

func (c *chapterCtrl) update() http.HandlerFunc {
	// Request type
	var request spotigraph.ChapterUpdateReq

	// Response type
	type response struct {
		Context string                     `json:"@context"`
		Type    string                     `json:"@type"`
		ID      string                     `json:"@id"`
		Entity  *spotigraph.Domain_Chapter `json:",inline"`
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
		res, err := c.chapters.Update(ctx, &request)
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSON(ctx, w, &response{
			Context: jsonldContext,
			Type:    "Chapter",
			ID:      fmt.Sprintf("/chapters/%s", res.Entity.Id),
			Entity:  res.Entity,
		})
	}
}

func (c *chapterCtrl) delete() http.HandlerFunc {
	// Response type
	type response struct {
		Context string                     `json:"@context"`
		Type    string                     `json:"@type"`
		ID      string                     `json:"@id"`
		Entity  *spotigraph.Domain_Chapter `json:",inline"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Delegate to service
		res, err := c.chapters.Delete(ctx, &spotigraph.ChapterGetReq{
			Id: chi.URLParamFromCtx(ctx, "id"),
		})
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSONStatus(ctx, w, http.StatusOK, "Chapter successfully deleted.")
	}
}

func (c *chapterCtrl) search() http.HandlerFunc {
	// Request type
	var request spotigraph.ChapterSearchReq

	// Response type
	type response struct {
		Context string                          `json:"@context"`
		Type    string                          `json:"@type"`
		ID      string                          `json:"@id"`
		Page    *spotigraph.PaginatedChapterRes `json:",inline"`
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
		res, err := c.chapters.Search(ctx, &request)
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
