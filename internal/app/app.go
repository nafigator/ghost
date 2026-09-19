// Package app that implement application.
package app

import (
	"runtime/debug"

	"go.uber.org/zap"

	"github.com/nafigator/ghost/internal/app/config"
)

const (
	buildMsg    = "Build info:\n%s"
	startMsg    = "ghost start"
	shutdownMsg = "ghost shutdown"
)

var (
	buildInfo = "develop" //nolint:gochecknoglobals //build-flag
	// ErrInfo signals that the application has output requested information
	// (help, version) and should exit successfully.
	ErrInfo = config.ErrInfo
)

// Run runs application.
func Run(log *zap.SugaredLogger) error {
	c, err := config.Init(buildInfo, log)
	if err != nil {
		return err
	}

	if bi, ok := debug.ReadBuildInfo(); ok {
		log.Debugf(buildMsg, bi.String())
	}

	log.Info(startMsg)
	defer log.Info(shutdownMsg)

	if err = generate(c); err != nil {
		return err
	}

	return nil
}
