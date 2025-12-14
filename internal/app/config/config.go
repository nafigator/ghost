// Package config provides app configuration.
package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/ardanlabs/conf/v3"
	"go.uber.org/zap"
)

const (
	appName   = "GHOST (Go High-level Open Service Templater)"
	cfgPrefix = "GHOST"
)

type Conf struct {
	conf.Version
	Name             string        `conf:"default:test,short:n,help:Project short name"`
	Description      string        `conf:"default:Go microservice,short:d,help:Project short description"`
	ModuleName       string        `conf:"default:github.com/test/test,short:m,help:Go module name"`
	GoImage          string        `conf:"default:nafigat0r/go:1.25.5,short:g,help:Go docker image"`
	LinterImage      string        `conf:"default:nafigat0r/golangci-lint:2.5.0,short:l,help:Linter docker image"`
	GovulncheckImage string        `conf:"default:nafigat0r/govulncheck:1.1.4,short:c,help:Govulncheck docker image"`
	ShutdownTimeout  time.Duration `conf:"default:10s,short:t,help:Timeout for graceful shutdown"`
	WithREST         bool          `conf:"default:false,short:r,help:Add HTTP server with REST API functionality"`
}

func Init(build string, log *zap.SugaredLogger) (*Conf, error) {
	var c string
	var err error

	cfg := Conf{
		Version: conf.Version{
			Build: build,
			Desc:  appName,
		},
	}

	if c, err = conf.Parse(cfgPrefix, &cfg); err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(c) //nolint:forbidigo // Need for raw output of help message
		}

		return nil, err
	}

	if c, err = conf.String(&cfg); err != nil {
		return nil, err
	}

	log.Infof("Initial config:\n%s", c)

	return &cfg, nil
}
