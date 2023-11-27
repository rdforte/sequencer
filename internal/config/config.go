package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	apiPortKey            = "API_PORT"
	debugPortKey          = "DEBUG_PORT"
	shutdownTimeoutSecKey = "SHUTDOWN_TIMEOUT_SEC"
	environmentKey        = "ENVIRONMENT"
)

func CreateAPIConfig() (*config, error) {
	apiPort, err := strconv.Atoi(os.Getenv(apiPortKey))
	if err != nil {
		return nil, fmt.Errorf("env var %s not a valid api port %w", apiPortKey, err)
	}

	debugPort, err := strconv.Atoi(os.Getenv(debugPortKey))
	if err != nil {
		return nil, fmt.Errorf("env var %s not a valid debug port %w", debugPortKey, err)
	}

	if debugPort == apiPort {
		return nil, fmt.Errorf("debug port %d should not be the same as api port %d", debugPort, apiPort)
	}

	shutdownTimeoutSec, err := strconv.Atoi(os.Getenv("SHUTDOWN_TIMEOUT_SEC"))
	if err != nil {
		return nil, fmt.Errorf("env var %s not a valid shutdown timeout %w", shutdownTimeoutSecKey, err)
	}

	env := os.Getenv(environmentKey)
	if len(env) == 0 {
		return nil, fmt.Errorf("no environment set for env %s", environmentKey)
	}

	return &config{
		apiPort:            apiPort,
		debugPort:          debugPort,
		shutdownTimeoutSec: time.Second * time.Duration(shutdownTimeoutSec),
		env:                env,
	}, nil
}

type config struct {
	apiPort            int
	debugPort          int
	shutdownTimeoutSec time.Duration
	env                string
}

func (c *config) ApiPort() int {
	return c.apiPort
}

func (c *config) DebugPort() int {
	return c.debugPort
}

func (c *config) ShutdownTimeout() time.Duration {
	return c.shutdownTimeoutSec
}

func (c *config) Env() string {
	return c.env
}
