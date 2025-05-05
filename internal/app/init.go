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

	//go:embed templates/internal/app/app.gotmpl
	appSrc string

	//go:embed templates/internal/app/config/config.gotmpl
	configSrc string
)

func templates() tps {
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
			name: "config",
			dir:  "internal/app/config",
			file: "internal/app/config/config.go",
			src:  configSrc,
		},
	}
}
