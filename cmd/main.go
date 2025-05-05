package main

import (
	"errors"
	"io/fs"

	c "github.com/ardanlabs/conf/v3"
	"github.com/nafigator/zapper"
	"github.com/nafigator/zapper/conf"

	"github.com/nafigator/ghost/internal/app"
)

func main() {
	log := zapper.Must(conf.Must())
	defer func() {
		// https://github.com/uber-go/zap/issues/328
		var pathError *fs.PathError
		if err := log.Sync(); err != nil && !errors.As(err, &pathError) {
			log.Error(err)
		}
	}()

	if err := app.Run(log); err != nil {
		if errors.Is(err, c.ErrHelpWanted) {
			return
		}

		log.Fatal(err) //nolint:gocritic // Exceptional case for fatal errors. Do not run defer.
	}
}
