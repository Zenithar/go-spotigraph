package v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/spotigraph/internal/services/test/mock"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

// -----------------------------------------------------------------------------
type TestCase struct {
	name           string
	requestMethod  string
	requestURL     string
	requestParams  func(*http.Request)
	requestBody    io.Reader
	prepare        func(context.Context, *mock.MockUser)
	expectedStatus int
	expectedBody   []byte
}

func userTestSpec(tt *TestCase) func(t *testing.T) {
	return func(t *testing.T) {
		g := NewGomegaWithT(t)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Arm mocks
		ctx := context.Background()
		users := mock.NewMockUser(ctrl)

		// Prepare the recorder
		rr := httptest.NewRecorder()

		// Prepare the request
		req, err := http.NewRequest(tt.requestMethod, tt.requestURL, tt.requestBody)
		g.Expect(err).To(BeNil(), "Request building error should be nil")

		// Update request
		if tt.requestParams != nil {
			tt.requestParams(req)
		}

		// Prepare the mocks:
		if tt.prepare != nil {
			tt.prepare(ctx, users)
		}

		// Initialize the handler
		underTest := UserRoutes(users)

		// Do the request
		underTest.ServeHTTP(rr, req)

		// assert results expectations
		g.Expect(rr.Code).To(Equal(tt.expectedStatus), "Status code should be has expected")
		g.Expect(rr.Body).ToNot(BeNil(), "Request Body should not be nil")
		g.Expect(rr.Body.Bytes()).To(Equal(tt.expectedBody), "Request body should be as expected")
	}
}

// -----------------------------------------------------------------------------
func TestCreateUserHandler(t *testing.T) {
	// Testcase list
	testCases := []*TestCase{
		{
			name:          "blank body request",
			requestMethod: "POST",
			requestURL:    "/",
			requestParams: func(r *http.Request) {
				r.Header.Set("Content-Type", "application/json")
			},
			requestBody:    bytes.NewBuffer([]byte("")),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   []byte(`{"@type":"Error","code":400,"message":"Unable to process this request"}`),
		},
		{
			name:          "invalid json request",
			requestMethod: "POST",
			requestURL:    "/",
			requestParams: func(r *http.Request) {
				r.Header.Set("Content-Type", "application/json")
			},
			requestBody:    bytes.NewBuffer([]byte("a]))")),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   []byte(`{"@type":"Error","code":400,"message":"Unable to process this request"}`),
		},
		{
			name:          "valid payload",
			requestMethod: "POST",
			requestURL:    "/",
			requestParams: func(r *http.Request) {
				r.Header.Set("Content-Type", "application/json")
			},
			requestBody: bytes.NewBuffer([]byte("{}")),
			prepare: func(ctx context.Context, users *mock.MockUser) {
				res := &spotigraph.SingleUserRes{
					Entity: &spotigraph.Domain_User{
						Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
					},
				}
				users.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(res, nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"User","@id":"/api/v1/users/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al","id":"0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"}`),
		},
		{
			name:          "principal conflict",
			requestMethod: "POST",
			requestURL:    "/",
			requestParams: func(r *http.Request) {
				r.Header.Set("Content-Type", "application/json")
			},
			requestBody: bytes.NewBuffer([]byte("{}")),
			prepare: func(ctx context.Context, users *mock.MockUser) {
				res := &spotigraph.SingleUserRes{
					Error: &spotigraph.Error{
						Code:    http.StatusConflict,
						Message: "Principal already used",
					},
				}
				users.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(res, nil)
			},
			expectedStatus: http.StatusConflict,
			expectedBody:   []byte(`{"@type":"Error","code":409,"message":"Principal already used"}`),
		}, {
			name:          "service error",
			requestMethod: "POST",
			requestURL:    "/",
			requestParams: func(r *http.Request) {
				r.Header.Set("Content-Type", "application/json")
			},
			requestBody: bytes.NewBuffer([]byte("{}")),
			prepare: func(ctx context.Context, users *mock.MockUser) {
				users.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(&spotigraph.SingleUserRes{}, db.ErrTooManyResults)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte(`{"@type":"Error","code":500,"message":"Unable to process this request"}`),
		},
	}

	// Subtests
	for _, tt := range testCases {
		t.Run(tt.name, userTestSpec(tt))
	}
}

func TestReadUserHandler(t *testing.T) {
	// Testcase list
	testCases := []*TestCase{
		{
			name:          "valid payload",
			requestMethod: "GET",
			requestURL:    "/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			requestBody:   nil,
			prepare: func(ctx context.Context, users *mock.MockUser) {
				res := &spotigraph.SingleUserRes{
					Entity: &spotigraph.Domain_User{
						Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
					},
				}
				users.EXPECT().Get(gomock.Any(), gomock.Any()).Times(1).Return(res, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"User","@id":"/api/v1/users/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al","id":"0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"}`),
		},
		{
			name:          "service error",
			requestMethod: "GET",
			requestURL:    "/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			requestBody:   nil,
			prepare: func(ctx context.Context, users *mock.MockUser) {
				users.EXPECT().Get(gomock.Any(), gomock.Any()).Times(1).Return(&spotigraph.SingleUserRes{}, db.ErrTooManyResults)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte(`{"@type":"Error","code":500,"message":"Unable to process this request"}`),
		},
		{
			name:          "entity not found",
			requestMethod: "GET",
			requestURL:    "/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			requestBody:   nil,
			prepare: func(ctx context.Context, users *mock.MockUser) {
				res := &spotigraph.SingleUserRes{
					Error: &spotigraph.Error{
						Code:    http.StatusNotFound,
						Message: "User not found !",
					},
				}
				users.EXPECT().Get(gomock.Any(), gomock.Any()).Times(1).Return(res, db.ErrNoResult)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   []byte(`{"@type":"Error","code":404,"message":"User not found !"}`),
		},
	}

	// Subtests
	for _, tt := range testCases {
		t.Run(tt.name, userTestSpec(tt))
	}
}

func TestUpdateUserHandler(t *testing.T) {
	// Testcase list
	testCases := []*TestCase{
		{
			name:          "valid payload",
			requestMethod: "POST",
			requestURL:    "/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			requestParams: func(r *http.Request) {
				r.Header.Set("Content-Type", "application/json")
			},
			requestBody: bytes.NewBuffer([]byte("{}")),
			prepare: func(ctx context.Context, users *mock.MockUser) {
				res := &spotigraph.SingleUserRes{
					Entity: &spotigraph.Domain_User{
						Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
					},
				}
				users.EXPECT().Update(gomock.Any(), gomock.Any()).Times(1).Return(res, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"User","@id":"/api/v1/users/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al","id":"0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"}`),
		},
		{
			name:          "invalid payload",
			requestMethod: "POST",
			requestURL:    "/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			requestParams: func(r *http.Request) {
				r.Header.Set("Content-Type", "application/json")
			},
			requestBody:    bytes.NewBuffer([]byte("{aa")),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   []byte(`{"@type":"Error","code":400,"message":"Unable to process this request"}`),
		},
		{
			name:          "service error",
			requestMethod: "POST",
			requestURL:    "/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			requestParams: func(r *http.Request) {
				r.Header.Set("Content-Type", "application/json")
			},
			requestBody: bytes.NewBuffer([]byte("{}")),
			prepare: func(ctx context.Context, users *mock.MockUser) {
				users.EXPECT().Update(gomock.Any(), gomock.Any()).Times(1).Return(&spotigraph.SingleUserRes{}, db.ErrTooManyResults)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte(`{"@type":"Error","code":500,"message":"Unable to process this request"}`),
		},
		{
			name:          "entity not found",
			requestMethod: "POST",
			requestURL:    "/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			requestParams: func(r *http.Request) {
				r.Header.Set("Content-Type", "application/json")
			},
			requestBody: bytes.NewBuffer([]byte("{}")),
			prepare: func(ctx context.Context, users *mock.MockUser) {
				res := &spotigraph.SingleUserRes{
					Error: &spotigraph.Error{
						Code:    http.StatusNotFound,
						Message: "User not found !",
					},
				}
				users.EXPECT().Update(gomock.Any(), gomock.Any()).Times(1).Return(res, db.ErrNoResult)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   []byte(`{"@type":"Error","code":404,"message":"User not found !"}`),
		},
	}

	// Subtests
	for _, tt := range testCases {
		t.Run(tt.name, userTestSpec(tt))
	}
}

func TestDeleteUserHandler(t *testing.T) {
	// Testcase list
	testCases := []*TestCase{
		{
			name:          "valid payload",
			requestMethod: "DELETE",
			requestURL:    "/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			requestBody:   nil,
			prepare: func(ctx context.Context, users *mock.MockUser) {
				users.EXPECT().Delete(gomock.Any(), gomock.Any()).Times(1).Return(&spotigraph.EmptyRes{}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"Status","code":200,"message":"User successfully deleted"}`),
		},
		{
			name:          "service error",
			requestMethod: "DELETE",
			requestURL:    "/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			requestBody:   nil,
			prepare: func(ctx context.Context, users *mock.MockUser) {
				users.EXPECT().Delete(gomock.Any(), gomock.Any()).Times(1).Return(&spotigraph.EmptyRes{}, db.ErrTooManyResults)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte(`{"@type":"Error","code":500,"message":"Unable to process this request"}`),
		},
		{
			name:          "entity not found",
			requestMethod: "DELETE",
			requestURL:    "/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			requestBody:   nil,
			prepare: func(ctx context.Context, users *mock.MockUser) {
				res := &spotigraph.EmptyRes{
					Error: &spotigraph.Error{
						Code:    http.StatusNotFound,
						Message: "User not found !",
					},
				}
				users.EXPECT().Delete(gomock.Any(), gomock.Any()).Times(1).Return(res, db.ErrNoResult)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   []byte(`{"@type":"Error","code":404,"message":"User not found !"}`),
		},
	}

	// Subtests
	for _, tt := range testCases {
		t.Run(tt.name, userTestSpec(tt))
	}
}

func TestSearchUserHandler(t *testing.T) {
	// Testcase list
	testCases := []*TestCase{
		{
			name:          "empty collection",
			requestMethod: "GET",
			requestURL:    "/",
			requestBody:   nil,
			prepare: func(ctx context.Context, users *mock.MockUser) {
				users.EXPECT().Search(gomock.Any(), gomock.Any()).Times(1).Return(&spotigraph.PaginatedUserRes{
					Count:       0,
					CurrentPage: 0,
					Total:       0,
					PerPage:     25,
					Members:     make([]*spotigraph.Domain_User, 0),
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"UserCollection","@id":"/","per_page":25}`),
		},
		{
			name:          "1 page collection",
			requestMethod: "GET",
			requestURL:    "/",
			requestBody:   nil,
			prepare: func(ctx context.Context, users *mock.MockUser) {
				users.EXPECT().Search(gomock.Any(), gomock.Any()).Times(1).Return(&spotigraph.PaginatedUserRes{
					Count:       1,
					CurrentPage: 1,
					Total:       3,
					PerPage:     25,
					Members: []*spotigraph.Domain_User{
						{
							Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
						},
						{
							Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
						},
						{
							Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
						},
					},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"UserCollection","@id":"/","total":3,"per_page":25,"count":1,"current_page":1,"members":[{"id":"0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"},{"id":"0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"},{"id":"0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"}]}`),
		},
		{
			name:          "1 page collection with principal",
			requestMethod: "GET",
			requestURL:    "/",
			requestParams: func(req *http.Request) {
				q := req.URL.Query()
				q.Add("principal", "toto@foo.org")
				req.URL.RawQuery = q.Encode()
			},
			requestBody: nil,
			prepare: func(ctx context.Context, users *mock.MockUser) {
				users.EXPECT().Search(gomock.Any(), gomock.Any()).Times(1).Return(&spotigraph.PaginatedUserRes{
					Count:       1,
					CurrentPage: 1,
					Total:       3,
					PerPage:     25,
					Members: []*spotigraph.Domain_User{
						{
							Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
						},
					},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"UserCollection","@id":"/?principal=toto%40foo.org","total":3,"per_page":25,"count":1,"current_page":1,"members":[{"id":"0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"}]}`),
		},
		{
			name:          "invalid page number",
			requestMethod: "GET",
			requestURL:    "/",
			requestParams: func(req *http.Request) {
				q := req.URL.Query()
				q.Add("page", "-4")
				req.URL.RawQuery = q.Encode()
			},
			requestBody: nil,
			prepare: func(ctx context.Context, users *mock.MockUser) {
				users.EXPECT().Search(gomock.Any(), gomock.Any()).Times(1).Return(&spotigraph.PaginatedUserRes{
					Count:       1,
					CurrentPage: 1,
					Total:       1,
					PerPage:     25,
					Members: []*spotigraph.Domain_User{
						{
							Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
						},
					},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"UserCollection","@id":"/?page=-4","total":1,"per_page":25,"count":1,"current_page":1,"members":[{"id":"0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"}]}`),
		}, {
			name:          "with page number and page size",
			requestMethod: "GET",
			requestURL:    "/",
			requestParams: func(req *http.Request) {
				q := req.URL.Query()
				q.Add("page", "2")
				q.Add("perPage", "2")
				req.URL.RawQuery = q.Encode()
			},
			requestBody: nil,
			prepare: func(ctx context.Context, users *mock.MockUser) {
				users.EXPECT().Search(gomock.Any(), gomock.Any()).Times(1).Return(&spotigraph.PaginatedUserRes{
					Count:       1,
					CurrentPage: 1,
					Total:       1,
					PerPage:     25,
					Members: []*spotigraph.Domain_User{
						{
							Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
						},
					},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"UserCollection","@id":"/?page=2\u0026perPage=2","total":1,"per_page":25,"count":1,"current_page":1,"members":[{"id":"0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"}]}`),
		},
		{
			name:          "service error",
			requestMethod: "GET",
			requestURL:    "/",
			requestBody:   nil,
			prepare: func(ctx context.Context, users *mock.MockUser) {
				users.EXPECT().Search(gomock.Any(), gomock.Any()).Times(1).Return(&spotigraph.PaginatedUserRes{
					Error: &spotigraph.Error{
						Code:    http.StatusNotFound,
						Message: "No results",
					},
				}, nil)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   []byte(`{"@type":"Error","code":404,"message":"No results"}`),
		},
		{
			name:          "service error",
			requestMethod: "GET",
			requestURL:    "/",
			requestBody:   nil,
			prepare: func(ctx context.Context, users *mock.MockUser) {
				users.EXPECT().Search(gomock.Any(), gomock.Any()).Times(1).Return(&spotigraph.PaginatedUserRes{}, db.ErrTooManyResults)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte(`{"@type":"Error","code":500,"message":"Unable to process this request"}`),
		},
	}

	// Subtests
	for _, tt := range testCases {
		t.Run(tt.name, userTestSpec(tt))
	}
}
