package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveTelemetryEnabled(t *testing.T) {

	dir, err := os.MkdirTemp("", "test-config-save")
	assert.NoError(t, err)

	// create homedir config file
	s := Storage{Dir: dir}
	err = s.Init()
	assert.NoError(t, err)

	t.Run("check if telemetry system is enabled", func(t *testing.T) {
		// given / when
		d, err := s.Load()

		// then
		assert.NoError(t, err)
		assert.Equal(t, true, d.TelemetryEnabled)

	})

	t.Run("check if telemetry system is disabled", func(t *testing.T) {
		// given
		d, err := s.Load()
		assert.NoError(t, err)

		d.DisableAnalytics()
		err = s.Save(d)
		assert.NoError(t, err)

		// when
		d, err = s.Load()

		// then
		assert.NoError(t, err)
		assert.Equal(t, false, d.TelemetryEnabled)

	})

}
