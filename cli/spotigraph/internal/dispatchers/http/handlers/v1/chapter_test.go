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
	chapterv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/chapter/v1"
	systemv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/system/v1"
)

// -----------------------------------------------------------------------------
type TestCase struct {
	name           string
	requestMethod  string
	requestURL     string
	requestParams  func(*http.Request)
	requestBody    io.Reader
	prepare        func(context.Context, *mock.MockChapter)
	expectedStatus int
	expectedBody   []byte
}

func chapterTestSpec(tt *TestCase) func(t *testing.T) {
	return func(t *testing.T) {
		g := NewGomegaWithT(t)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Arm mocks
		ctx := context.Background()
		chapters := mock.NewMockChapter(ctrl)

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
			tt.prepare(ctx, chapters)
		}

		// Initialize the handler
		underTest := ChapterRoutes(chapters)

		// Do the request
		underTest.ServeHTTP(rr, req)

		// assert results expectations
		g.Expect(rr.Code).To(Equal(tt.expectedStatus), "Status code should be has expected")
		g.Expect(rr.Body).ToNot(BeNil(), "Request Body should not be nil")
		g.Expect(rr.Body.Bytes()).To(Equal(tt.expectedBody), "Request body should be as expected")
	}
}

// -----------------------------------------------------------------------------
func TestCreateChapterHandler(t *testing.T) {
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
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				res := &chapterv1.CreateResponse{
					Entity: &chapterv1.Chapter{
						Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
					},
				}
				chapters.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(res, nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"Chapter","@id":"/api/v1/chapters/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al","id":"0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"}`),
		},
		{
			name:          "principal conflict",
			requestMethod: "POST",
			requestURL:    "/",
			requestParams: func(r *http.Request) {
				r.Header.Set("Content-Type", "application/json")
			},
			requestBody: bytes.NewBuffer([]byte("{}")),
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				res := &chapterv1.CreateResponse{
					Error: &systemv1.Error{
						Code:    http.StatusConflict,
						Message: "Principal already used",
					},
				}
				chapters.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(res, nil)
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
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(&chapterv1.CreateResponse{}, db.ErrTooManyResults)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte(`{"@type":"Error","code":500,"message":"Oups, something goes wrong during request handling !"}`),
		},
	}

	// Subtests
	for _, tt := range testCases {
		t.Run(tt.name, chapterTestSpec(tt))
	}
}

func TestReadChapterHandler(t *testing.T) {
	// Testcase list
	testCases := []*TestCase{
		{
			name:          "valid payload",
			requestMethod: "GET",
			requestURL:    "/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			requestBody:   nil,
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				res := &chapterv1.GetResponse{
					Entity: &chapterv1.Chapter{
						Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
					},
				}
				chapters.EXPECT().Get(gomock.Any(), gomock.Any()).Times(1).Return(res, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"Chapter","@id":"/api/v1/chapters/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al","id":"0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"}`),
		},
		{
			name:          "service error",
			requestMethod: "GET",
			requestURL:    "/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			requestBody:   nil,
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Get(gomock.Any(), gomock.Any()).Times(1).Return(&chapterv1.GetResponse{}, db.ErrTooManyResults)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte(`{"@type":"Error","code":500,"message":"Oups, something goes wrong during request handling !"}`),
		},
		{
			name:          "entity not found",
			requestMethod: "GET",
			requestURL:    "/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			requestBody:   nil,
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				res := &chapterv1.GetResponse{
					Error: &systemv1.Error{
						Code:    http.StatusNotFound,
						Message: "Chapter not found !",
					},
				}
				chapters.EXPECT().Get(gomock.Any(), gomock.Any()).Times(1).Return(res, db.ErrNoResult)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   []byte(`{"@type":"Error","code":404,"message":"Chapter not found !"}`),
		},
	}

	// Subtests
	for _, tt := range testCases {
		t.Run(tt.name, chapterTestSpec(tt))
	}
}

func TestUpdateChapterHandler(t *testing.T) {
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
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				res := &chapterv1.UpdateResponse{
					Entity: &chapterv1.Chapter{
						Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
					},
				}
				chapters.EXPECT().Update(gomock.Any(), gomock.Any()).Times(1).Return(res, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"Chapter","@id":"/api/v1/chapters/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al","id":"0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"}`),
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
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Update(gomock.Any(), gomock.Any()).Times(1).Return(&chapterv1.UpdateResponse{}, db.ErrTooManyResults)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte(`{"@type":"Error","code":500,"message":"Oups, something goes wrong during request handling !"}`),
		},
		{
			name:          "entity not found",
			requestMethod: "POST",
			requestURL:    "/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			requestParams: func(r *http.Request) {
				r.Header.Set("Content-Type", "application/json")
			},
			requestBody: bytes.NewBuffer([]byte("{}")),
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				res := &chapterv1.UpdateResponse{
					Error: &systemv1.Error{
						Code:    http.StatusNotFound,
						Message: "Chapter not found !",
					},
				}
				chapters.EXPECT().Update(gomock.Any(), gomock.Any()).Times(1).Return(res, db.ErrNoResult)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   []byte(`{"@type":"Error","code":404,"message":"Chapter not found !"}`),
		},
	}

	// Subtests
	for _, tt := range testCases {
		t.Run(tt.name, chapterTestSpec(tt))
	}
}

func TestDeleteChapterHandler(t *testing.T) {
	// Testcase list
	testCases := []*TestCase{
		{
			name:          "valid payload",
			requestMethod: "DELETE",
			requestURL:    "/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			requestBody:   nil,
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Delete(gomock.Any(), gomock.Any()).Times(1).Return(&chapterv1.DeleteResponse{}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"Status","code":200,"message":"Chapter successfully deleted"}`),
		},
		{
			name:          "service error",
			requestMethod: "DELETE",
			requestURL:    "/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			requestBody:   nil,
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Delete(gomock.Any(), gomock.Any()).Times(1).Return(&chapterv1.DeleteResponse{}, db.ErrTooManyResults)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte(`{"@type":"Error","code":500,"message":"Oups, something goes wrong during request handling !"}`),
		},
		{
			name:          "entity not found",
			requestMethod: "DELETE",
			requestURL:    "/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			requestBody:   nil,
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				res := &chapterv1.DeleteResponse{
					Error: &systemv1.Error{
						Code:    http.StatusNotFound,
						Message: "Chapter not found !",
					},
				}
				chapters.EXPECT().Delete(gomock.Any(), gomock.Any()).Times(1).Return(res, db.ErrNoResult)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   []byte(`{"@type":"Error","code":404,"message":"Chapter not found !"}`),
		},
	}

	// Subtests
	for _, tt := range testCases {
		t.Run(tt.name, chapterTestSpec(tt))
	}
}

func TestSearchChapterHandler(t *testing.T) {
	// Testcase list
	testCases := []*TestCase{
		{
			name:          "empty collection",
			requestMethod: "GET",
			requestURL:    "/",
			requestBody:   nil,
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Search(gomock.Any(), gomock.Any()).Times(1).Return(&chapterv1.SearchResponse{
					Count:       0,
					CurrentPage: 0,
					Total:       0,
					PerPage:     25,
					Members:     make([]*chapterv1.Chapter, 0),
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"ChapterCollection","@id":"/","per_page":25}`),
		},
		{
			name:          "1 page collection",
			requestMethod: "GET",
			requestURL:    "/",
			requestBody:   nil,
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Search(gomock.Any(), gomock.Any()).Times(1).Return(&chapterv1.SearchResponse{
					Count:       1,
					CurrentPage: 1,
					Total:       3,
					PerPage:     25,
					Members: []*chapterv1.Chapter{
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
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"ChapterCollection","@id":"/","total":3,"per_page":25,"count":1,"current_page":1,"members":[{"id":"0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"},{"id":"0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"},{"id":"0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"}]}`),
		},
		{
			name:          "1 page collection with principal",
			requestMethod: "GET",
			requestURL:    "/",
			requestParams: func(req *http.Request) {
				q := req.URL.Query()
				q.Add("chapter_id", "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al")
				req.URL.RawQuery = q.Encode()
			},
			requestBody: nil,
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Search(gomock.Any(), gomock.Any()).Times(1).Return(&chapterv1.SearchResponse{
					Count:       1,
					CurrentPage: 1,
					Total:       3,
					PerPage:     25,
					Members: []*chapterv1.Chapter{
						{
							Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
						},
					},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"ChapterCollection","@id":"/?chapter_id=0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al","total":3,"per_page":25,"count":1,"current_page":1,"members":[{"id":"0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"}]}`),
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
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Search(gomock.Any(), gomock.Any()).Times(1).Return(&chapterv1.SearchResponse{
					Count:       1,
					CurrentPage: 1,
					Total:       1,
					PerPage:     25,
					Members: []*chapterv1.Chapter{
						{
							Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
						},
					},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"ChapterCollection","@id":"/?page=-4","total":1,"per_page":25,"count":1,"current_page":1,"members":[{"id":"0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"}]}`),
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
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Search(gomock.Any(), gomock.Any()).Times(1).Return(&chapterv1.SearchResponse{
					Count:       1,
					CurrentPage: 1,
					Total:       1,
					PerPage:     25,
					Members: []*chapterv1.Chapter{
						{
							Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
						},
					},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"ChapterCollection","@id":"/?page=2\u0026perPage=2","total":1,"per_page":25,"count":1,"current_page":1,"members":[{"id":"0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"}]}`),
		},
		{
			name:          "service error",
			requestMethod: "GET",
			requestURL:    "/",
			requestBody:   nil,
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Search(gomock.Any(), gomock.Any()).Times(1).Return(&chapterv1.SearchResponse{
					Error: &systemv1.Error{
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
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Search(gomock.Any(), gomock.Any()).Times(1).Return(&chapterv1.SearchResponse{}, db.ErrTooManyResults)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte(`{"@type":"Error","code":500,"message":"Oups, something goes wrong during request handling !"}`),
		},
	}

	// Subtests
	for _, tt := range testCases {
		t.Run(tt.name, chapterTestSpec(tt))
	}
}
