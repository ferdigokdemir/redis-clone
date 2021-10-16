package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetItemHandler(t *testing.T) {
	t.Log("Testing GetItemHandler")
	LoadDatabase()
	data["test"] = "test"

	t.Run("returns KeyNotFound error", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/api/v1/getItem", strings.NewReader(`{"key": "this_key_is_just_for_test"}`))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(GetItemHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// expected body
		expected := `{"error":"KeyNotFound","success":false}`

		if strings.TrimSpace(rr.Body.String()) != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("returns success result for 'test' key", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/api/v1/getItem", strings.NewReader(`{"key": "test"}`))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(GetItemHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// expected body
		expected := `{"data":{"key":"test","value":"test"},"success":true}`

		if strings.TrimSpace(rr.Body.String()) != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

}

func TestSetItemHandler(t *testing.T) {
	t.Log("Testing SetItemHandler")

	t.Run("returns success result for 'test' key", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/api/v1/setItem", strings.NewReader(`{"key": "test", "value": "test"}`))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(SetItemHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// expected body
		expected := `{"data":null,"success":true}`

		if strings.TrimSpace(rr.Body.String()) != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("returns error result for invalid key or value", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/api/v1/setItem", strings.NewReader(`{"key2": "test", "value2": "test"}`))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(SetItemHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// expected body
		expected := `{"error":"SetItemError","success":false}`

		if strings.TrimSpace(rr.Body.String()) != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}

func TestFlushItemsHandler(t *testing.T) {
	t.Log("Testing FlushItemsHandler")

	t.Run("returns success result", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/api/v1/flushItems", strings.NewReader(`{}`))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(FlushItemsHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// expected body
		expected := `{"data":null,"success":true}`

		if strings.TrimSpace(rr.Body.String()) != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}
