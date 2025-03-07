package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"auto_verse/Modules/auth/controllers"
)

func TestAuthController_GetHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/auth", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	controller := controllers.NewAuthController()
	controller.GetHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expected := "Auth Get Handler"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}
}
