package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"auto_verse/Modules/users/controllers"
)

func TestUsersController_GetHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	controller := controllers.NewUsersController()
	controller.GetHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expected := "Users Get Handler"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}
}
