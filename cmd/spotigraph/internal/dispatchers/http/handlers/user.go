package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type userCtrl struct {
	users services.User
}

// -----------------------------------------------------------------------------

// UserRoutes returns user management related API
func UserRoutes(users services.User) chi.Router {
	r := chi.NewRouter()

	// Initialize controller
	ctrl := &userCtrl{
		users: users,
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

func (c *userCtrl) create() http.HandlerFunc {
	// Request type
	var request spotigraph.UserCreateReq

	// Response type
	type response struct {
		Context                 string `json:"@context"`
		Type                    string `json:"@type"`
		ID                      string `json:"@id"`
		*spotigraph.Domain_User `json:",omitempty"`
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
		res, err := c.users.Create(ctx, &request)
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSON(ctx, w, &response{
			Context:     jsonldContext,
			Type:        "User",
			ID:          fmt.Sprintf("/users/%s", res.Entity.Id),
			Domain_User: res.Entity,
		})
	}
}

func (c *userCtrl) read() http.HandlerFunc {
	// Response type
	type response struct {
		Context string                  `json:"@context"`
		Type    string                  `json:"@type"`
		ID      string                  `json:"@id"`
		Entity  *spotigraph.Domain_User `json:",inline"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Delegate to service
		res, err := c.users.Get(ctx, &spotigraph.UserGetReq{
			Id: chi.URLParamFromCtx(ctx, "id"),
		})
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSON(ctx, w, &response{
			Context: jsonldContext,
			Type:    "User",
			ID:      fmt.Sprintf("/users/%s", res.Entity.Id),
			Entity:  res.Entity,
		})
	}
}

func (c *userCtrl) update() http.HandlerFunc {
	// Request type
	var request spotigraph.UserUpdateReq

	// Response type
	type response struct {
		Context string                  `json:"@context"`
		Type    string                  `json:"@type"`
		ID      string                  `json:"@id"`
		Entity  *spotigraph.Domain_User `json:",inline"`
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
		res, err := c.users.Update(ctx, &request)
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSON(ctx, w, &response{
			Context: jsonldContext,
			Type:    "User",
			ID:      fmt.Sprintf("/users/%s", res.Entity.Id),
			Entity:  res.Entity,
		})
	}
}

func (c *userCtrl) delete() http.HandlerFunc {
	// Response type
	type response struct {
		Context string                  `json:"@context"`
		Type    string                  `json:"@type"`
		ID      string                  `json:"@id"`
		Entity  *spotigraph.Domain_User `json:",inline"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Delegate to service
		res, err := c.users.Delete(ctx, &spotigraph.UserGetReq{
			Id: chi.URLParamFromCtx(ctx, "id"),
		})
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSONStatus(ctx, w, http.StatusOK, "User successfully deleted.")
	}
}

func (c *userCtrl) search() http.HandlerFunc {
	// Request type
	var request spotigraph.UserSearchReq

	// Response type
	type response struct {
		Context string                       `json:"@context"`
		Type    string                       `json:"@type"`
		ID      string                       `json:"@id"`
		Page    *spotigraph.PaginatedUserRes `json:",inline"`
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
		res, err := c.users.Search(ctx, &request)
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
