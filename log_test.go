package bulog

import (
	"os"
	"testing"
)

func TestNewLog(t *testing.T) {
	log := New(os.Stdout)

	log.Info().Msg("hello world")
	log.Debug().Msg("hello world")
}
