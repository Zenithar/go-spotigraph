package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"go.zenithar.org/spotigraph/internal/services/test/mock"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

func TestListUsersHandler(t *testing.T) {
	g := gomock.NewController(t)
	defer g.Finish()

	// Forge request
	req, err := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer([]byte("{}")))
	if err != nil {
		t.Fatal(err)
	}

	// Prepare the recorder
	rr := httptest.NewRecorder()

	// Mock service calls
	srv := mock.NewMockUser(g)
	srv.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(&spotigraph.SingleUserRes{
		Entity: &spotigraph.Domain_User{
			Id:        "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			Principal: "",
		},
	}, nil)

	// Initialize the handler
	ctrl := &userCtrl{
		users: srv,
	}
	handler := ctrl.create()
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"@context":"https://go.zenithar.org/spotigraph/v1","@type":"User","@id":"/users/0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al","id":"0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
