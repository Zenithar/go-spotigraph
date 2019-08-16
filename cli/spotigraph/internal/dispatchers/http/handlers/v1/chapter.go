package v1

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/gogo/protobuf/types"

	"go.zenithar.org/pkg/web/request"
	"go.zenithar.org/pkg/web/respond"
	"go.zenithar.org/spotigraph/internal/services"
	chapterv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/chapter/v1"
)

type chapterCtrl struct {
	chapters services.Chapter
}

// -----------------------------------------------------------------------------

// ChapterRoutes returns chapter management related API
func ChapterRoutes(chapters services.Chapter) http.Handler {
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
		r.Post("/", ctrl.update())
		r.Delete("/", ctrl.delete())
	})

	// Return router
	return r
}

// -----------------------------------------------------------------------------

func (c *chapterCtrl) create() http.HandlerFunc {
	// Response type
	type response struct {
		*respond.Resource
		*chapterv1.Chapter `json:",omitempty"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Request type
		var req chapterv1.CreateRequest

		// Decode request
		if err := request.Parse(r, &req); err != nil {
			respond.WithError(w, r, http.StatusBadRequest, err)
			return
		}

		// Delegate to service
		res, err := c.chapters.Create(ctx, &req)
		if publicError(w, r, res, err) {
			return
		}

		// Marshal response
		respond.With(w, r, http.StatusCreated, &response{
			Resource: &respond.Resource{
				Context: jsonldContext,
				Type:    "Chapter",
				ID:      fmt.Sprintf("/api/v1/chapters/%s", res.Entity.Id),
			},
			Chapter: res.Entity,
		})
	}
}

func (c *chapterCtrl) read() http.HandlerFunc {
	// Response type
	type response struct {
		*respond.Resource
		*chapterv1.Chapter `json:",omitempty"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Delegate to service
		res, err := c.chapters.Get(ctx, &chapterv1.GetRequest{
			Id: chi.URLParamFromCtx(ctx, "id"),
		})
		if publicError(w, r, res, err) {
			return
		}

		// Marshal response
		respond.With(w, r, http.StatusOK, &response{
			Resource: &respond.Resource{
				Context: jsonldContext,
				Type:    "Chapter",
				ID:      fmt.Sprintf("/api/v1/chapters/%s", res.Entity.Id),
			},
			Chapter: res.Entity,
		})
	}
}

func (c *chapterCtrl) update() http.HandlerFunc {
	// Response type
	type response struct {
		*respond.Resource
		*chapterv1.Chapter `json:",omitempty"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Request type
		var req chapterv1.UpdateRequest

		// Decode request as json
		if err := request.Parse(r, &req); err != nil {
			respond.WithError(w, r, http.StatusBadRequest, err)
			return
		}

		// Delegate to service
		res, err := c.chapters.Update(ctx, &req)
		if publicError(w, r, res, err) {
			return
		}

		// Marshal response
		respond.With(w, r, http.StatusOK, &response{
			Resource: &respond.Resource{
				Context: jsonldContext,
				Type:    "Chapter",
				ID:      fmt.Sprintf("/api/v1/chapters/%s", res.Entity.Id),
			},
			Chapter: res.Entity,
		})
	}
}

func (c *chapterCtrl) delete() http.HandlerFunc {
	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Delegate to service
		res, err := c.chapters.Delete(ctx, &chapterv1.DeleteRequest{
			Id: chi.URLParamFromCtx(ctx, "id"),
		})
		if publicError(w, r, res, err) {
			return
		}

		// Marshal response
		respond.With(w, r, http.StatusOK, &respond.Status{
			Resource: &respond.Resource{
				Context: jsonldContext,
				Type:    "Status",
			},
			Code:    http.StatusOK,
			Message: "Chapter successfully deleted",
		})
	}
}

func (c *chapterCtrl) search() http.HandlerFunc {
	// Response type
	type response struct {
		*respond.Resource
		*chapterv1.SearchResponse `json:",inline"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		var (
			q         = r.URL.Query()
			page      = q.Get("page")
			perPage   = q.Get("perPage")
			sorts     = strings.Split(q.Get("sorts"), ",")
			chapterID = q.Get("chapter_id")
		)

		// Prepare request filter
		req := &chapterv1.SearchRequest{
			Page:    toUint32(page, 1),
			PerPage: toUint32(perPage, 25),
			Sorts:   sorts,
		}
		if chapterID != "" {
			req.ChapterId = &types.StringValue{Value: chapterID}
		}

		// Delegate to service
		res, err := c.chapters.Search(ctx, req)
		if publicError(w, r, res, err) {
			return
		}

		// Marshal response
		respond.With(w, r, http.StatusOK, &response{
			Resource: &respond.Resource{
				Context: jsonldContext,
				Type:    "ChapterCollection",
				ID:      r.URL.RequestURI(),
			},
			SearchResponse: res,
		})
	}
}
