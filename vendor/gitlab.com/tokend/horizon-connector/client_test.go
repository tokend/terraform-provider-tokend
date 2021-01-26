package horizon

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestClient_Do(t *testing.T) {
	handler := func(status int, body []byte) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(status)
			w.Write(body)
		}
	}

	do := func(ts *httptest.Server) ([]byte, error) {
		endpoint, err := url.Parse(ts.URL)
		if err != nil {
			t.Fatal(err)
		}

		client := NewClient(ts.Client(), endpoint)
		request, err := http.NewRequest("GET", ts.URL, nil)
		if err != nil {
			t.Fatal(err)
		}
		return client.Do(request)
	}

	t.Run("404", func(t *testing.T) {
		ts := httptest.NewServer(handler(404, nil))
		defer ts.Close()

		got, err := do(ts)
		assert.NoError(t, err)
		assert.Nil(t, got)
	})

	t.Run("200", func(t *testing.T) {
		expected := []byte(`{"msg": "ok"}`)
		ts := httptest.NewServer(handler(200, expected))
		defer ts.Close()

		got, err := do(ts)
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("401", func(t *testing.T) {
		response := []byte(`{"msg": "unauthorized"}`)
		ts := httptest.NewServer(handler(401, response))
		defer ts.Close()

		got, err := do(ts)
		assert.Error(t, err)
		assert.Nil(t, got)

		cause := errors.Cause(err)
		herr, ok := cause.(Error)
		assert.True(t, ok)
		if ok {
			assert.EqualValues(t, 401, herr.Status())
			assert.Equal(t, response, herr.Body())
		}
	})

	t.Run("500", func(t *testing.T) {
		response := []byte(`{"msg": "internal server error"}`)
		ts := httptest.NewServer(handler(500, response))
		defer ts.Close()

		got, err := do(ts)
		assert.Error(t, err)
		assert.Nil(t, got)

		cause := errors.Cause(err)
		herr, ok := cause.(Error)
		assert.True(t, ok)
		if ok {
			assert.EqualValues(t, 500, herr.Status())
			assert.Equal(t, response, herr.Body())
		}
	})
}
