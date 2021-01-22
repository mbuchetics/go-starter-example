package common_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"go-starter-example/internal/api"
	"go-starter-example/internal/test"
)

func TestSwaggerYAMLRetrieval(t *testing.T) {
	t.Parallel()

	test.WithTestServer(t, func(s *api.Server) {
		res := test.PerformRequest(t, s, "GET", "/swagger.yml", nil, nil)
		require.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
}
