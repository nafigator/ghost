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
	//go:embed templates/default/cmd/main.gotmpl
	mainSrc string

	//go:embed templates/default/gomod.gotmpl
	gomodSrc string

	//go:embed templates/default/gosum.gotmpl
	gosumSrc string

	//go:embed templates/default/golangci.gotmpl
	golangciSrc string

	//go:embed templates/default/gitignore.gotmpl
	gitignoreSrc string

	//go:embed templates/default/makefile.gotmpl
	makefileSrc string

	//go:embed templates/rest/default/makefile.gotmpl
	restMakefileSrc string

	//go:embed templates/default/compose.gotmpl
	composeSrc string

	//go:embed templates/default/build/compose.override.gotmpl
	composeOverrideSrc string

	//go:embed templates/rest/default/build/compose.override.gotmpl
	restComposeOverrideSrc string

	//go:embed templates/default/build/zapper.gotmpl
	zapperSrc string

	//go:embed templates/default/internal/app/app.gotmpl
	appSrc string

	//go:embed templates/default/internal/app/shutdown.gotmpl
	shutdownSrc string

	//go:embed templates/default/internal/app/config/config.gotmpl
	configSrc string

	//go:embed templates/rest/default/internal/app/config/config.gotmpl
	restConfigSrc string

	//go:embed templates/default/internal/app/container/container.gotmpl
	containerSrc string

	//go:embed templates/default/internal/app/http/init.gotmpl
	httpInitSrc string

	//go:embed templates/rest/default/internal/app/http/init.gotmpl
	restHTTPInitSrc string

	//go:embed templates/default/internal/app/http/mux.gotmpl
	httpMuxSrc string

	//go:embed templates/rest/default/internal/app/http/mux.gotmpl
	restHTTPMuxSrc string

	//go:embed templates/default/internal/app/http/handlers/support/build.gotmpl
	buildSrc string

	//go:embed templates/default/internal/app/http/handlers/support/health.gotmpl
	healthSrc string

	//go:embed templates/default/internal/app/http/handlers/support/log.gotmpl
	logSrc string

	//go:embed templates/default/internal/app/http/handlers/support/response/response.gotmpl
	responseSrc string

	//go:embed templates/default/internal/app/http/handlers/support/version.gotmpl
	versionSrc string

	//go:embed templates/rest/default/internal/app/http/handlers/api/index.gotmpl
	indexAPISrc string

	//go:embed templates/default/internal/app/http/validators/validators.gotmpl
	validatorsSrc string

	//go:embed templates/default/internal/app/readiness/readiness.gotmpl
	readinessSrc string

	//go:embed templates/rest/default/internal/app/readiness/readiness.gotmpl
	restReadinessSrc string

	//go:embed templates/default/internal/app/http/errors/errors.gotmpl
	errorsSrc string

	//go:embed templates/default/internal/sdk/readiness/socket/socket.gotmpl
	readinessSocketSrc string

	//go:embed templates/default/internal/sdk/http/mux/middleware.gotmpl
	middlewareSrc string

	//go:embed templates/default/internal/sdk/http/mux/mux.gotmpl
	sdkMuxSrc string

	//go:embed templates/default/internal/sdk/http/mux/respond.gotmpl
	respondSrc string

	//go:embed templates/default/internal/sdk/http/mux/validator.gotmpl
	validatorSrc string
)

func templates(c *config.Conf) tps {
	t := common()

	if c.WithREST {
		t["index"] = tp{
			dir:  "internal/app/http/handlers/api",
			file: "internal/app/http/handlers/api/index.go",
			src:  indexAPISrc,
		}

		t["config"] = tp{
			dir:  "internal/app/config",
			file: "internal/app/config/config.go",
			src:  restConfigSrc,
		}

		t["makefile"] = tp{
			file: "Makefile",
			src:  restMakefileSrc,
		}

		t["compose-override"] = tp{
			file: "docker-compose.override.yml",
			src:  restComposeOverrideSrc,
		}

		t["init"] = tp{
			dir:  "internal/app/http", //nolint:goconst	// Don't change for readability.
			file: "internal/app/http/init.go",
			src:  restHTTPInitSrc,
		}

		t["mux"] = tp{
			dir:  "internal/app/http",
			file: "internal/app/http/mux.go",
			src:  restHTTPMuxSrc,
		}

		t["readiness"] = tp{
			dir:  "internal/app/readiness",
			file: "internal/app/readiness/readiness.go",
			src:  restReadinessSrc,
		}

		t["compose-override-dist"] = tp{
			dir:  "build", //nolint:goconst	// Don't change for readability.
			file: "build/docker-compose.override.dist.yml",
			src:  restComposeOverrideSrc,
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
		"gosum": {
			file: "go.sum",
			src:  gosumSrc,
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
			file: "docker-compose.override.yml",
			src:  composeOverrideSrc,
		},
		"compose-override-dist": {
			dir:  "build",
			file: "build/docker-compose.override.dist.yml",
			src:  composeOverrideSrc,
		},
		"zapper": {
			dir:  "bin",
			file: "bin/zapper.yml",
			src:  zapperSrc,
		},
		"zapper-dist": {
			dir:  "build",
			file: "build/zapper.dist.yml",
			src:  zapperSrc,
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
			dir:  "internal/app/http/handlers/support", //nolint:goconst	// Don't change for readability.
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
		"socket": {
			dir:  "internal/sdk/readiness/socket",
			file: "internal/sdk/readiness/socket/socket.go",
			src:  readinessSocketSrc,
		},
		"middleware": {
			dir:  "internal/sdk/http/mux", //nolint:goconst	// Don't change for readability.
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
