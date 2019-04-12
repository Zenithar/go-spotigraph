package v1

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/gogo/protobuf/types"

	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
	"go.zenithar.org/spotigraph/pkg/request"
	"go.zenithar.org/spotigraph/pkg/respond"
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
	var req spotigraph.UserCreateReq

	// Response type
	type response struct {
		*respond.Resource
		*spotigraph.Domain_User `json:",omitempty"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Decode request
		if err := request.Parse(r, &req); err != nil {
			respond.WithError(w, r, http.StatusBadRequest, err)
			return
		}

		// Delegate to service
		res, err := c.users.Create(ctx, &req)
		if publicError(w, r, res, err) {
			return
		}

		// Marshal response
		respond.With(w, r, http.StatusCreated, &response{
			Resource: &respond.Resource{
				Context: jsonldContext,
				Type:    "User",
				ID:      fmt.Sprintf("/api/v1/users/%s", res.Entity.Id),
			},
			Domain_User: res.Entity,
		})
	}
}

func (c *userCtrl) read() http.HandlerFunc {
	// Response type
	type response struct {
		*respond.Resource
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
		if publicError(w, r, res, err) {
			return
		}

		// Marshal response
		respond.With(w, r, http.StatusOK, &response{
			Resource: &respond.Resource{
				Context: jsonldContext,
				Type:    "User",
				ID:      fmt.Sprintf("/api/v1/users/%s", res.Entity.Id),
			},
			Domain_User: res.Entity,
		})
	}
}

func (c *userCtrl) update() http.HandlerFunc {
	// Request type
	var req spotigraph.UserUpdateReq

	// Response type
	type response struct {
		*respond.Resource
		*spotigraph.Domain_User `json:",omitempty"`
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Decode request as json
		if err := request.Parse(r, &req); err != nil {
			respond.WithError(w, r, http.StatusBadRequest, err)
			return
		}

		// Delegate to service
		res, err := c.users.Update(ctx, &req)
		if publicError(w, r, res, err) {
			return
		}

		// Marshal response
		respond.With(w, r, http.StatusOK, &response{
			Resource: &respond.Resource{
				Context: jsonldContext,
				Type:    "User",
				ID:      fmt.Sprintf("/api/v1/users/%s", res.Entity.Id),
			},
			Domain_User: res.Entity,
		})
	}
}

func (c *userCtrl) delete() http.HandlerFunc {
	// Response type
	type response struct {
	}

	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Delegate to service
		res, err := c.users.Delete(ctx, &spotigraph.UserGetReq{
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
			Message: "User successfully deleted",
		})
	}
}

func (c *userCtrl) search() http.HandlerFunc {
	// Response type
	type response struct {
		*respond.Resource
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
		if publicError(w, r, res, err) {
			return
		}

		// Marshal response
		respond.With(w, r, http.StatusOK, &response{
			Resource: &respond.Resource{
				Context: jsonldContext,
				Type:    "UserCollection",
				ID:      r.URL.RequestURI(),
			},
			PaginatedUserRes: res,
		})
	}
}
