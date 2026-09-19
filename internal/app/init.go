package app

import (
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/nafigator/ghost/internal/app/config"
)

type tp struct {
	dir  string
	path string
	src  string
}

type tps map[string]tp

func transformName(name string) string {
	if v, ok := rename[name]; ok {
		name = v
	} else if before, found := strings.CutSuffix(name, ".gotmpl"); found {
		name = before + ".go"
	}

	return name
}

var rename = map[string]string{ //nolint:gochecknoglobals // Acknowledged
	"makefile.gotmpl":               "Makefile",
	"gomod.gotmpl":                  "go.mod",
	"gosum.gotmpl":                  "go.sum",
	"golangci.gotmpl":               ".golangci.yml",
	"gitignore.gotmpl":              ".gitignore",
	"compose.gotmpl":                "docker-compose.yml",
	"build/compose.override.gotmpl": "build/docker-compose.override.dist.yml",
	"build/zapper.gotmpl":           "build/zapper.dist.yml",
}

var (
	//go:embed templates/common
	commonFS embed.FS

	//go:embed templates/rest/common
	restFS embed.FS
)

func templates(c *config.Conf) (tps, error) {
	tt := make(tps)

	tt, err := sources(tt, commonFS, "templates/common")
	if err != nil {
		return nil, err
	}

	if c.WithREST {
		tt, err = sources(tt, restFS, "templates/rest/common")
		if err != nil {
			return nil, err
		}
	}

	return tt, nil
}

func sources(tt tps, f embed.FS, trimPath string) (tps, error) {
	sfs, err := fs.Sub(f, trimPath)
	if err != nil {
		return nil, err
	}

	return load(tt, sfs)
}

func load(tt tps, rfs fs.FS) (tps, error) {
	err := fs.WalkDir(rfs, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		b, rErr := fs.ReadFile(rfs, p)
		if rErr != nil {
			return fmt.Errorf("read %s: %w", p, rErr)
		}

		out := filepath.FromSlash(transformName(p))

		t := tp{
			dir:  filepath.Dir(out),
			path: out,
			src:  string(b),
		}

		tt[out] = t

		return nil
	})
	if err != nil {
		return nil, err
	}

	return tt, nil
}
