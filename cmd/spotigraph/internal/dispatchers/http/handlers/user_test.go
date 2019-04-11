package handlers

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"

	"go.zenithar.org/spotigraph/internal/services/test/mock"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

func TestCreateUserHandler(t *testing.T) {
	// Testcase list
	testCases := []struct {
		name           string
		requestMethod  string
		requestURL     string
		requestBody    io.Reader
		prepare        func(context.Context, *mock.MockUser)
		wantErr        bool
		expectedStatus int
		expectedBody   []byte
	}{
		{
			name:           "blank body request",
			requestMethod:  "POST",
			requestURL:     "/",
			requestBody:    bytes.NewBuffer([]byte("")),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"Error","code":400,"message":"Unable to handle this request"}`),
		},
		{
			name:           "invalid json request",
			requestMethod:  "POST",
			requestURL:     "/",
			requestBody:    bytes.NewBuffer([]byte("a]))")),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"Error","code":400,"message":"Unable to handle this request"}`),
		},
		{
			name:          "valid payload",
			requestMethod: "POST",
			requestURL:    "/",
			requestBody:   bytes.NewBuffer([]byte("{}")),
			prepare: func(ctx context.Context, users *mock.MockUser) {
				res := &spotigraph.SingleUserRes{
					Entity: &spotigraph.Domain_User{
						Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
					},
				}
				users.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(res, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"User","@id":"/users/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al","id":"0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"}`),
		},
		{
			name:          "service error",
			requestMethod: "POST",
			requestURL:    "/",
			requestBody:   bytes.NewBuffer([]byte("{}")),
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
			expectedBody:   []byte(`{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"Error","code":409,"message":"Principal already used"}`),
		},
	}

	// Subtests
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
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
		})
	}
}
