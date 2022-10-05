package bulog

import (
	"os"
	"testing"
)

func TestNewLog(t *testing.T) {
	log := New(os.Stdout)

	log.Info().Msg("hello world")
	log.Debug().Msg("hello world")
	log.Warn().Msg("hello world")
	log.Error().Msg("hello world")
        log.Fatal().Msg("hello world")

	go infoLog(log)
	go debugLog(log)
	go warnLog(log)
	go errorLog(log)
        go fatalLog(log)
}

func infoLog(log Logger) {
	log.Info().Msg("hello world")
}

func debugLog(log Logger) {
	log.Debug().Msg("hello world")
}

func warnLog(log Logger) {
	log.Warn().Msg("hello world")
}

func errorLog(log Logger) {
	log.Error().Msg("hello world")
}

func fatalLog(log Logger) {
	log.Fatal().Msg("hello world")
}
