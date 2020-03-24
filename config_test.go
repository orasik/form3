package form3_test

import (
	"github.com/orasik/form3"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func Test_Config_Missing_Required(t *testing.T) {
	_, err := form3.ParseConfig()

	if err.Error() != "required environment variable \"API_BASEURL\" is not set. required environment variable \"ACCOUNT_ENDPOINT\" is not set" {
		t.Errorf("expected error for missing API_BASEURL and ACCOUNT_ENDPOINT, got %s", err.Error())
	}
}

func Test_Config_Missing_Defaults(t *testing.T) {
	os.Setenv("API_BASEURL", "baseurl")
	os.Setenv("ACCOUNT_ENDPOINT", "endpoint")

	cfg, err := form3.ParseConfig()
	assert.Equal(t, cfg.AccountBaseURL, "baseurl")
	assert.Equal(t, cfg.AccountEndPoint, "endpoint")
	assert.Equal(t, cfg.Port, 2400)
	assert.Equal(t, cfg.LogFormat, "json")
	assert.Equal(t, cfg.LogLevel, "debug")
	assert.Equal(t, cfg.Timeout, 3*time.Second)
	assert.Equal(t, err, nil)
}

func Test_Config_With_Port(t *testing.T) {
	os.Setenv("API_BASEURL", "baseurl")
	os.Setenv("ACCOUNT_ENDPOINT", "endpoint")
	os.Setenv("PORT", "5000")

	cfg, err := form3.ParseConfig()
	assert.Equal(t, cfg.AccountBaseURL, "baseurl")
	assert.Equal(t, cfg.AccountEndPoint, "endpoint")
	assert.Equal(t, cfg.Port, 5000)
	assert.Equal(t, cfg.LogFormat, "json")
	assert.Equal(t, cfg.LogLevel, "debug")
	assert.Equal(t, cfg.Timeout, 3*time.Second)
	assert.Equal(t, err, nil)
}

func Test_Config_With_Wrong_Log_Level(t *testing.T) {
	os.Setenv("API_BASEURL", "baseurl")
	os.Setenv("ACCOUNT_ENDPOINT", "endpoint")
	os.Setenv("PORT", "3000")
	os.Setenv("LOG_LEVEL", "blah")

	cfg, err := form3.ParseConfig()
	assert.Equal(t, cfg.AccountBaseURL, "baseurl")
	assert.Equal(t, cfg.AccountEndPoint, "endpoint")
	assert.Equal(t, cfg.Port, 3000)
	assert.Equal(t, cfg.LogFormat, "json")
	assert.Equal(t, cfg.LogLevel, "blah")
	assert.Equal(t, cfg.Timeout, 3*time.Second)
	assert.Equal(t, err, nil)
	assert.Equal(t, log.GetLevel(), log.DebugLevel)
}

func Test_Config_With_Fatal_Log_Level(t *testing.T) {
	os.Setenv("API_BASEURL", "baseurl")
	os.Setenv("ACCOUNT_ENDPOINT", "endpoint")
	os.Setenv("PORT", "3000")
	os.Setenv("LOG_LEVEL", "fatal")

	cfg, err := form3.ParseConfig()
	assert.Equal(t, cfg.AccountBaseURL, "baseurl")
	assert.Equal(t, cfg.AccountEndPoint, "endpoint")
	assert.Equal(t, cfg.Port, 3000)
	assert.Equal(t, cfg.LogFormat, "json")
	assert.Equal(t, cfg.LogLevel, "fatal")
	assert.Equal(t, cfg.Timeout, 3*time.Second)
	assert.Equal(t, err, nil)
	assert.Equal(t, log.GetLevel(), log.FatalLevel)
}

func Test_Config_With_Error_Log_Level(t *testing.T) {
	os.Setenv("API_BASEURL", "baseurl")
	os.Setenv("ACCOUNT_ENDPOINT", "endpoint")
	os.Setenv("PORT", "3000")
	os.Setenv("LOG_LEVEL", "error")

	cfg, err := form3.ParseConfig()
	assert.Equal(t, cfg.AccountBaseURL, "baseurl")
	assert.Equal(t, cfg.AccountEndPoint, "endpoint")
	assert.Equal(t, cfg.Port, 3000)
	assert.Equal(t, cfg.LogFormat, "json")
	assert.Equal(t, cfg.LogLevel, "error")
	assert.Equal(t, cfg.Timeout, 3*time.Second)
	assert.Equal(t, err, nil)
	assert.Equal(t, log.GetLevel(), log.ErrorLevel)
}

func Test_Config_With_Warn_Log_Level(t *testing.T) {
	os.Setenv("API_BASEURL", "baseurl")
	os.Setenv("ACCOUNT_ENDPOINT", "endpoint")
	os.Setenv("PORT", "3000")
	os.Setenv("LOG_LEVEL", "warn")

	cfg, err := form3.ParseConfig()
	assert.Equal(t, cfg.AccountBaseURL, "baseurl")
	assert.Equal(t, cfg.AccountEndPoint, "endpoint")
	assert.Equal(t, cfg.Port, 3000)
	assert.Equal(t, cfg.LogFormat, "json")
	assert.Equal(t, cfg.LogLevel, "warn")
	assert.Equal(t, cfg.Timeout, 3*time.Second)
	assert.Equal(t, err, nil)
	assert.Equal(t, log.GetLevel(), log.WarnLevel)
}

func Test_Config_With_Info_Log_Level(t *testing.T) {
	os.Setenv("API_BASEURL", "baseurl")
	os.Setenv("ACCOUNT_ENDPOINT", "endpoint")
	os.Setenv("PORT", "3000")
	os.Setenv("LOG_LEVEL", "info")

	cfg, err := form3.ParseConfig()
	assert.Equal(t, cfg.AccountBaseURL, "baseurl")
	assert.Equal(t, cfg.AccountEndPoint, "endpoint")
	assert.Equal(t, cfg.Port, 3000)
	assert.Equal(t, cfg.LogFormat, "json")
	assert.Equal(t, cfg.LogLevel, "info")
	assert.Equal(t, cfg.Timeout, 3*time.Second)
	assert.Equal(t, err, nil)
	assert.Equal(t, log.GetLevel(), log.InfoLevel)
}
