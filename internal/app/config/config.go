// Package config provides app configuration.
package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/ardanlabs/conf/v3"
)

const (
	appName   = "GHOST (Go High-level Open Service Templater)"
	cfgPrefix = "GHOST"
)

type logger interface {
	Infof(format string, args ...any)
}

type Conf struct { //nolint:govet // Acknowledged
	conf.Args
	conf.Version

	OutputDir        string        `conf:"-"`
	Name             string        `conf:"default:test,short:n,help:Project short name"`
	Description      string        `conf:"default:Go microservice,short:d,help:Project short description"`
	ModuleName       string        `conf:"default:github.com/test/test,short:m,help:Go module name"`
	GoImage          string        `conf:"default:nafigat0r/go:1.26.0,short:g,help:Go docker image"`
	LinterImage      string        `conf:"default:nafigat0r/golangci-lint:2.9.0,short:l,help:Linter docker image"`
	GovulncheckImage string        `conf:"default:nafigat0r/govulncheck:1.1.4,short:c,help:Govulncheck docker image"`
	ShutdownTimeout  time.Duration `conf:"default:10s,short:t,help:Timeout for graceful shutdown"`
	WithREST         bool          `conf:"default:false,short:r,help:Add HTTP server with REST API functionality"`
}

// ErrInfo is returned when the user requested informational output
// (help, version) and the application should exit without error.
var ErrInfo = errors.New("info wanted")

func Init(build string, log logger) (*Conf, error) {
	var c string
	var err error

	cfg := Conf{
		Version: conf.Version{
			Build: build,
			Desc:  appName,
		},
	}

	if c, err = conf.Parse(cfgPrefix, &cfg); err != nil {
		if errors.Is(err, conf.ErrHelpWanted) || errors.Is(err, conf.ErrVersionWanted) {
			fmt.Println(c) //nolint:forbidigo // Need for raw output of help message

			return nil, ErrInfo
		}

		return nil, err
	}

	if err = resolveOutputDir(&cfg); err != nil {
		return nil, err
	}

	if c, err = conf.String(&cfg); err != nil {
		return nil, err
	}

	log.Infof("Initial config:\n%s", c)

	return &cfg, nil
}

func resolveOutputDir(cfg *Conf) error {
	cfg.OutputDir = cfg.Num(0)
	if cfg.OutputDir == "" {
		return errors.New("output directory not specified")
	}

	if d, err := os.Stat(cfg.OutputDir); err == nil {
		if !d.IsDir() {
			return fmt.Errorf("output dir %q exists and is not a directory", cfg.OutputDir)
		}
	} else if !errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("stat %s: %w", cfg.OutputDir, err)
	}

	return nil
}
