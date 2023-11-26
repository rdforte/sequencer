package config_test

import (
	"fmt"
	"github.com/rdforte/sequencer/internal/config"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	t.Run("should return config along with values when all env variables are set", func(t *testing.T) {
		defer teardown()

		wantAPIPort := 3000
		wantDebugPort := 3001
		shutdownTimeoutSec := 20
		wantShutdownTimeout := time.Duration(shutdownTimeoutSec) * time.Second

		os.Setenv("API_PORT", strconv.Itoa(wantAPIPort))
		os.Setenv("DEBUG_PORT", strconv.Itoa(wantDebugPort))
		os.Setenv("SHUTDOWN_TIMEOUT_SEC", strconv.Itoa(shutdownTimeoutSec))

		cfg, err := config.CreateAPIConfig()
		assert.Nil(t, err, fmt.Sprintf("wanted error nil but got %v", err))

		assert.Equal(t, wantAPIPort, cfg.ApiPort())
		assert.Equal(t, wantDebugPort, cfg.DebugPort())
		assert.Equal(t, wantShutdownTimeout, cfg.ShutdownTimeout())
	})

	t.Run("should return an api port error when there is no API_PORT env variable", func(t *testing.T) {
		defer teardown()

		os.Setenv("DEBUG_PORT", "3001")
		os.Setenv("SHUTDOWN_TIMEOUT_SEC", "20")

		_, err := config.CreateAPIConfig()
		assert.Error(t, err, "wanted error but got nil")
	})

	t.Run("should return an debug port error when there is no DEBUG_PORT env variable", func(t *testing.T) {
		defer teardown()

		os.Setenv("API_PORT", "3001")
		os.Setenv("SHUTDOWN_TIMEOUT_SEC", "20")

		_, err := config.CreateAPIConfig()
		assert.Error(t, err, "wanted error but got nil")
	})

	t.Run("should return an shutdown timeout error when there is no SHUTDOWN_TIMEOUT_SEC env variable", func(t *testing.T) {
		defer teardown()

		os.Setenv("API_PORT", "3000")
		os.Setenv("DEBUG_PORT", "3001")

		_, err := config.CreateAPIConfig()
		assert.Error(t, err, "wanted error but got nil")
	})

	t.Run("should return an error when all env variables are set but api port is the same as debug port", func(t *testing.T) {
		defer teardown()

		os.Setenv("API_PORT", "3000")
		os.Setenv("DEBUG_PORT", "3000")
		os.Setenv("SHUTDOWN_TIMEOUT_SEC", "20")

		_, err := config.CreateAPIConfig()
		assert.Error(t, err, "wanted error but got nil")
	})
}

func teardown() {
	os.Unsetenv("API_PORT")
	os.Unsetenv("DEBUG_PORT")
	os.Unsetenv("SHUTDOWN_TIMEOUT_SEC")
}
