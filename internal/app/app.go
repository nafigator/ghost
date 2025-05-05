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

var build = "develop" //nolint:gochecknoglobals //build-flag

// Run runs application.
func Run(log *zap.SugaredLogger) error {
	var c *config.Conf
	var err error

	if c, err = config.Init(build, log); err != nil {
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
