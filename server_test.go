package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	Serve(res, req)

	got := res.Code
	want := 200

	assertStatus(t, got, want)

}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("\nGot status: %d\nWant status: %d", got, want)
	}
}
