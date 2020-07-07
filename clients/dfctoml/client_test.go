package dfctoml

import (
	"net/http"
	"strings"
	"testing"

	"github.com/dfc/go/support/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientURL(t *testing.T) {
	//HACK:  we're testing an internal method rather than setting up a http client
	//mock.

	c := &Client{UseHTTP: false}
	assert.Equal(t, "https://dfc.org/.well-known/dfc.toml", c.url("dfc.org"))

	c = &Client{UseHTTP: true}
	assert.Equal(t, "http://dfc.org/.well-known/dfc.toml", c.url("dfc.org"))
}

func TestClient(t *testing.T) {
	h := httptest.NewClient()
	c := &Client{HTTP: h}

	// happy path
	h.
		On("GET", "https://dfc.org/.well-known/dfc.toml").
		ReturnString(http.StatusOK,
			`FEDERATION_SERVER="https://localhost/federation"`,
		)
	stoml, err := c.GetStellarToml("dfc.org")
	require.NoError(t, err)
	assert.Equal(t, "https://localhost/federation", stoml.FederationServer)

	// dfc.toml exceeds limit
	h.
		On("GET", "https://toobig.org/.well-known/dfc.toml").
		ReturnString(http.StatusOK,
			`FEDERATION_SERVER="https://localhost/federation`+strings.Repeat("0", StellarTomlMaxSize)+`"`,
		)
	stoml, err = c.GetStellarToml("toobig.org")
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "dfc.toml response exceeds")
	}

	// not found
	h.
		On("GET", "https://missing.org/.well-known/dfc.toml").
		ReturnNotFound()
	stoml, err = c.GetStellarToml("missing.org")
	assert.EqualError(t, err, "http request failed with non-200 status code")

	// invalid toml
	h.
		On("GET", "https://json.org/.well-known/dfc.toml").
		ReturnJSON(http.StatusOK, map[string]string{"hello": "world"})
	stoml, err = c.GetStellarToml("json.org")

	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "toml decode failed")
	}
}
