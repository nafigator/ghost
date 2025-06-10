package app

import (
	_ "embed"

	"github.com/nafigator/ghost/internal/app/config"
)

type tp struct {
	dir  string
	file string
	src  string
}

type tps map[string]tp

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

	//go:embed templates/makefile-api.gotmpl
	makefileAPISrc string

	//go:embed templates/compose.gotmpl
	composeSrc string

	//go:embed templates/compose.override.gotmpl
	composeOverrideSrc string

	//go:embed templates/compose-api.override.gotmpl
	composeOverrideAPISrc string

	//go:embed templates/zap.gotmpl
	zapSrc string

	//go:embed templates/internal/app/app.gotmpl
	appSrc string

	//go:embed templates/internal/app/shutdown.gotmpl
	shutdownSrc string

	//go:embed templates/internal/app/config/config.gotmpl
	configSrc string

	//go:embed templates/internal/app/config/config-api.gotmpl
	configAPISrc string

	//go:embed templates/internal/app/container/container.gotmpl
	containerSrc string

	//go:embed templates/internal/app/http/init.gotmpl
	httpInitSrc string

	//go:embed templates/internal/app/http/init-api.gotmpl
	httpInitAPISrc string

	//go:embed templates/internal/app/http/mux.gotmpl
	httpMuxSrc string

	//go:embed templates/internal/app/http/mux-api.gotmpl
	httpMuxAPISrc string

	//go:embed templates/internal/app/http/handlers/support/build.gotmpl
	buildSrc string

	//go:embed templates/internal/app/http/handlers/support/health.gotmpl
	healthSrc string

	//go:embed templates/internal/app/http/handlers/support/log.gotmpl
	logSrc string

	//go:embed templates/internal/app/http/handlers/support/response/response.gotmpl
	responseSrc string

	//go:embed templates/internal/app/http/handlers/support/version.gotmpl
	versionSrc string

	//go:embed templates/internal/app/http/handlers/api/index.gotmpl
	indexAPISrc string

	//go:embed templates/internal/app/http/validators/validators.gotmpl
	validatorsSrc string

	//go:embed templates/internal/app/readiness/readiness.gotmpl
	readinessSrc string

	//go:embed templates/internal/app/readiness/readiness-api.gotmpl
	readinessAPISrc string

	//go:embed templates/internal/app/http/errors/errors.gotmpl
	errorsSrc string

	//go:embed templates/internal/sdk/http/mux/middleware.gotmpl
	middlewareSrc string

	//go:embed templates/internal/sdk/http/mux/mux.gotmpl
	sdkMuxSrc string

	//go:embed templates/internal/sdk/http/mux/respond.gotmpl
	respondSrc string

	//go:embed templates/internal/sdk/http/mux/validator.gotmpl
	validatorSrc string
)

func templates(c *config.Conf) tps {
	t := common()

	if c.WithREST {
		t["makefile"] = tp{
			file: "Makefile",
			src:  makefileAPISrc,
		}

		t["compose-override"] = tp{
			file: "docker-compose.override.yml.dist",
			src:  composeOverrideAPISrc,
		}

		t["config"] = tp{
			dir:  "internal/app/config",
			file: "internal/app/config/config.go",
			src:  configAPISrc,
		}

		t["init"] = tp{
			dir:  "internal/app/http",
			file: "internal/app/http/init.go",
			src:  httpInitAPISrc,
		}

		t["mux"] = tp{
			dir:  "internal/app/http",
			file: "internal/app/http/mux.go",
			src:  httpMuxAPISrc,
		}

		t["index"] = tp{
			dir:  "internal/app/http/handlers/api",
			file: "internal/app/http/handlers/api/index.go",
			src:  indexAPISrc,
		}

		t["readiness"] = tp{
			dir:  "internal/app/readiness",
			file: "internal/app/readiness/readiness.go",
			src:  readinessAPISrc,
		}
	}

	return t
}

func common() tps { //nolint:funlen  // This function supposed to be longer than check limit.
	return tps{
		"golangci": {
			file: ".golangci.yml",
			src:  golangciSrc,
		},
		"gitignore": {
			file: ".gitignore",
			src:  gitignoreSrc,
		},
		"gomod": {
			file: "go.mod",
			src:  gomodSrc,
		},
		"makefile": {
			file: "Makefile",
			src:  makefileSrc,
		},
		"compose": {
			file: "docker-compose.yml",
			src:  composeSrc,
		},
		"compose-override": {
			file: "docker-compose.override.yml.dist",
			src:  composeOverrideSrc,
		},
		"zap": {
			file: "config.yml.dist",
			src:  zapSrc,
		},
		"main": {
			dir:  "cmd",
			file: "cmd/main.go",
			src:  mainSrc,
		},
		"app": {
			dir:  "internal/app",
			file: "internal/app/app.go",
			src:  appSrc,
		},
		"shutdown": {
			dir:  "internal/app",
			file: "internal/app/shutdown.go",
			src:  shutdownSrc,
		},
		"config": {
			dir:  "internal/app/config",
			file: "internal/app/config/config.go",
			src:  configSrc,
		},
		"container": {
			dir:  "internal/app/container",
			file: "internal/app/container/container.go",
			src:  containerSrc,
		},
		"init": {
			dir:  "internal/app/http",
			file: "internal/app/http/init.go",
			src:  httpInitSrc,
		},
		"mux": {
			dir:  "internal/app/http",
			file: "internal/app/http/mux.go",
			src:  httpMuxSrc,
		},
		"build": {
			dir:  "internal/app/http/handlers/support",
			file: "internal/app/http/handlers/support/build.go",
			src:  buildSrc,
		},
		"health": {
			dir:  "internal/app/http/handlers/support",
			file: "internal/app/http/handlers/support/health.go",
			src:  healthSrc,
		},
		"log": {
			dir:  "internal/app/http/handlers/support",
			file: "internal/app/http/handlers/support/log.go",
			src:  logSrc,
		},
		"responses": {
			dir:  "internal/app/http/handlers/support/response",
			file: "internal/app/http/handlers/support/response/response.go",
			src:  responseSrc,
		},
		"version": {
			dir:  "internal/app/http/handlers/support",
			file: "internal/app/http/handlers/support/version.go",
			src:  versionSrc,
		},
		"validators": {
			dir:  "internal/app/http/validators",
			file: "internal/app/http/validators/validators.go",
			src:  validatorsSrc,
		},
		"errors": {
			dir:  "internal/app/http/errors",
			file: "internal/app/http/errors/errors.go",
			src:  errorsSrc,
		},
		"readiness": {
			dir:  "internal/app/readiness",
			file: "internal/app/readiness/readiness.go",
			src:  readinessSrc,
		},
		"middleware": {
			dir:  "internal/sdk/http/mux",
			file: "internal/sdk/http/mux/middleware.go",
			src:  middlewareSrc,
		},
		"sdkmux": {
			dir:  "internal/sdk/http/mux",
			file: "internal/sdk/http/mux/mux.go",
			src:  sdkMuxSrc,
		},
		"respond": {
			dir:  "internal/sdk/http/mux",
			file: "internal/sdk/http/mux/respond.go",
			src:  respondSrc,
		},
		"validator": {
			dir:  "internal/sdk/http/mux",
			file: "internal/sdk/http/mux/validator.go",
			src:  validatorSrc,
		},
	}
}
