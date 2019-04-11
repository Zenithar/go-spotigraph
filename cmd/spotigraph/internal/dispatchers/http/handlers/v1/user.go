package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/gogo/protobuf/types"

	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type userCtrl struct {
	users services.User
}

// -----------------------------------------------------------------------------

// UserRoutes returns user management related API
func UserRoutes(users services.User) http.Handler {
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
		r.Post("/", ctrl.update())
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
		if err != nil || res.Error != nil {
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
		Context                 string `json:"@context"`
		Type                    string `json:"@type"`
		ID                      string `json:"@id"`
		*spotigraph.Domain_User `json:",omitempty"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Delegate to service
		res, err := c.users.Get(ctx, &spotigraph.UserGetReq{
			Id: chi.URLParamFromCtx(ctx, "id"),
		})
		if err != nil || res != nil && res.Error != nil {
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

func (c *userCtrl) update() http.HandlerFunc {
	// Request type
	var request spotigraph.UserUpdateReq

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
		res, err := c.users.Update(ctx, &request)
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

func (c *userCtrl) delete() http.HandlerFunc {
	// Response type
	type response struct {
		Context string `json:"@context"`
		Type    string `json:"@type"`
		ID      string `json:"@id"`
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
	// Response type
	type response struct {
		Context                      string `json:"@context"`
		Type                         string `json:"@type"`
		ID                           string `json:"@id"`
		*spotigraph.PaginatedUserRes `json:",inline"`
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
			principal = q.Get("principal")
		)

		// Prepare request filter
		req := &spotigraph.UserSearchReq{
			Page:    toUint32(page, 1),
			PerPage: toUint32(perPage, 25),
			Sorts:   sorts,
		}
		if principal != "" {
			req.Principal = &types.StringValue{Value: principal}
		}

		// Delegate to service
		res, err := c.users.Search(ctx, req)
		if err != nil {
			asJSONResultError(ctx, w, res.Error, err)
			return
		}

		// Marshal response
		asJSON(ctx, w, &response{
			Context:          jsonldContext,
			Type:             "UserCollection",
			ID:               r.URL.RequestURI(),
			PaginatedUserRes: res,
		})
	}
}
