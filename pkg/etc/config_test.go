package etc

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetAPIConfig(t *testing.T) {
	testCases := []struct {
		name           string
		expectedConfig APIConfig
	}{
		{
			name: "Should return default config",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			config, err := GetAPIConfig()
			require.NoError(t, err)
			assert.Equal(t, APIConfig{
				Addr:         ":8080",
				ReadTimeout:  15 * time.Second,
				WriteTimeout: 15 * time.Second,
			}, config)
		})
	}
}
