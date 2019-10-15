package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestHandler(t *testing.T) {
	app := NewAPIHandler(BuildInfo{
		Version: "v0.0.1",
		Commit:  "sha1",
		Date:    "15 Oct 2019 13:45",
	})

	ts := httptest.NewServer(app)
	defer ts.Close()

	t.Run("GET /api/info", func(t *testing.T) {
		rs, err := ts.Client().Get(ts.URL + "/api/info")

		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, rs.StatusCode)
		assert.Equal(t, "application/json", rs.Header.Get("Content-Type"))

		bodyBytes, err := ioutil.ReadAll(rs.Body)
		require.NoError(t, err)

		assert.JSONEq(t, `{
  "version": "v0.0.1",
  "commit": "sha1",
  "date": "15 Oct 2019 13:45"
}`, string(bodyBytes))
	})

	t.Run("GET /api/health", func(t *testing.T) {
		rs, err := ts.Client().Get(ts.URL + "/api/health")

		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, rs.StatusCode)
	})

}
