package app

import (
	_ "embed"
)

type tp struct {
	name string
	dir  string
	file string
	src  string
}

type tps []tp

var (
	//go:embed templates/cmd/main.gotmpl
	mainSrc string

	//go:embed templates/gomod.gotmpl
	gomodSrc string

	//go:embed templates/golangci.gotmpl
	golangciSrc string

	//go:embed templates/gitignore.gotmpl
	gitignoreSrc string

	//go:embed templates/makefile.gotmpl
	makefileSrc string

	//go:embed templates/compose.gotmpl
	composeSrc string

	//go:embed templates/zap.gotmpl
	zapSrc string

	//go:embed templates/internal/app/app.gotmpl
	appSrc string

	//go:embed templates/internal/app/shutdown.gotmpl
	shutdownSrc string

	//go:embed templates/internal/app/config/config.gotmpl
	configSrc string

	//go:embed templates/internal/app/container/container.gotmpl
	containerSrc string

	//go:embed templates/internal/app/http/init.gotmpl
	httpInitSrc string

	//go:embed templates/internal/app/http/mux.gotmpl
	httpMuxSrc string

	//go:embed templates/internal/app/http/handlers/support/build.gotmpl
	buildSrc string

	//go:embed templates/internal/app/http/handlers/support/health.gotmpl
	healthSrc string

	//go:embed templates/internal/app/http/handlers/support/log.gotmpl
	logSrc string

	//go:embed templates/internal/app/http/handlers/support/responses.gotmpl
	responsesSrc string

	//go:embed templates/internal/app/http/handlers/support/version.gotmpl
	versionSrc string

	//go:embed templates/internal/app/http/validators/validators.gotmpl
	validatorsSrc string

	//go:embed templates/internal/app/http/errors/errors.gotmpl
	errorsSrc string

	//go:embed templates/internal/sdk/pointer/pointer.gotmpl
	pointerSrc string

	//go:embed templates/internal/sdk/http/mux/middleware.gotmpl
	middlewareSrc string

	//go:embed templates/internal/sdk/http/mux/mux.gotmpl
	sdkMuxSrc string

	//go:embed templates/internal/sdk/http/mux/respond.gotmpl
	respondSrc string

	//go:embed templates/internal/sdk/http/mux/validator.gotmpl
	validatorSrc string
)

func templates() tps { //nolint:funlen  // This function supposed to be longer than check limit.
	return tps{
		{
			name: "golangci",
			file: ".golangci.yml",
			src:  golangciSrc,
		},
		{
			name: "gitignore",
			file: ".gitignore",
			src:  gitignoreSrc,
		},
		{
			name: "gomod",
			file: "go.mod",
			src:  gomodSrc,
		},
		{
			name: "makefile",
			file: "Makefile",
			src:  makefileSrc,
		},
		{
			name: "compose",
			file: "docker-compose.yml",
			src:  composeSrc,
		},
		{
			name: "zap",
			file: "config.yml.orig",
			src:  zapSrc,
		},
		{
			name: "main",
			dir:  "cmd",
			file: "cmd/main.go",
			src:  mainSrc,
		},
		{
			name: "app",
			dir:  "internal/app",
			file: "internal/app/app.go",
			src:  appSrc,
		},
		{
			name: "shutdown",
			dir:  "internal/app",
			file: "internal/app/shutdown.go",
			src:  shutdownSrc,
		},
		{
			name: "config",
			dir:  "internal/app/config",
			file: "internal/app/config/config.go",
			src:  configSrc,
		},
		{
			name: "container",
			dir:  "internal/app/container",
			file: "internal/app/container/container.go",
			src:  containerSrc,
		},
		{
			name: "init",
			dir:  "internal/app/http",
			file: "internal/app/http/init.go",
			src:  httpInitSrc,
		},
		{
			name: "mux",
			dir:  "internal/app/http",
			file: "internal/app/http/mux.go",
			src:  httpMuxSrc,
		},
		{
			name: "build",
			dir:  "internal/app/http/handlers/support",
			file: "internal/app/http/handlers/support/build.go",
			src:  buildSrc,
		},
		{
			name: "health",
			dir:  "internal/app/http/handlers/support",
			file: "internal/app/http/handlers/support/health.go",
			src:  healthSrc,
		},
		{
			name: "log",
			dir:  "internal/app/http/handlers/support",
			file: "internal/app/http/handlers/support/log.go",
			src:  logSrc,
		},
		{
			name: "responses",
			dir:  "internal/app/http/handlers/support",
			file: "internal/app/http/handlers/support/responses.go",
			src:  responsesSrc,
		},
		{
			name: "version",
			dir:  "internal/app/http/handlers/support",
			file: "internal/app/http/handlers/support/version.go",
			src:  versionSrc,
		},
		{
			name: "validators",
			dir:  "internal/app/http/validators",
			file: "internal/app/http/validators/validators.go",
			src:  validatorsSrc,
		},
		{
			name: "errors",
			dir:  "internal/app/http/errors",
			file: "internal/app/http/errors/errors.go",
			src:  errorsSrc,
		},
		{
			name: "pointer",
			dir:  "internal/sdk/pointer",
			file: "internal/sdk/pointer/pointer.go",
			src:  pointerSrc,
		},
		{
			name: "middleware",
			dir:  "internal/sdk/http/mux",
			file: "internal/sdk/http/mux/middleware.go",
			src:  middlewareSrc,
		},
		{
			name: "sdkmux",
			dir:  "internal/sdk/http/mux",
			file: "internal/sdk/http/mux/mux.go",
			src:  sdkMuxSrc,
		},
		{
			name: "respond",
			dir:  "internal/sdk/http/mux",
			file: "internal/sdk/http/mux/respond.go",
			src:  respondSrc,
		},
		{
			name: "validator",
			dir:  "internal/sdk/http/mux",
			file: "internal/sdk/http/mux/validator.go",
			src:  validatorSrc,
		},
	}
}
