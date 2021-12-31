package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// these tests require that this data exists in the InfluxDB instance exactly as it is...for now
func TestServer(t *testing.T) {
	rule := "MEASUREMENT>region>app>FIELD"
	measurement := "test"
	t.Run("make tree from rule", func(t *testing.T) {
		gotTree := makeTreeFromRule(rule, measurement)
		wantTree := Tree{
			0: &Measurement{"test", 0},
			1: &Key{"region", 1, []string{"us-east", "us-west"}, nil, nil},
			2: &Key{"app", 2, []string{"cart", "home", "login"}, nil, nil},
			3: &Field{"FIELD", 3},
		}

		assertEqual(t, gotTree, wantTree)
	})
	t.Run("make tree no rule", func(t *testing.T) {
		gotTree := makeTreeNoRule(measurement)
		wantTree := Tree{
			0: &Measurement{"test", 0},
			1: &Key{"region", 1, []string{"us-east", "us-west"}, nil, nil},
			2: &Key{"app", 2, []string{"cart", "home", "login"}, nil, nil},
			3: &Field{"FIELD", 3},
		}

		assertEqual(t, gotTree, wantTree)
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("\nGot status: %d\nWant status: %d", got, want)
	}
}

func newGetReq(path string) *http.Request {
	return httptest.NewRequest(http.MethodGet, path, nil)
}
