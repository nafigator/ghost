package app

import (
	"os"
	"strings"
	"text/template"

	"github.com/nafigator/ghost/internal/app/config"
)

const (
	dirStrictMode  = 0750
	fileStrictMode = 0640
)

// generate service code from templates.
func generate(c *config.Conf) error {
	var f *os.File
	var err error
	var tpl *template.Template

	vars := map[string]any{
		"GoModule":         c.ModuleName,
		"Name":             c.Name,
		"Description":      c.Description,
		"GoImage":          c.GoImage,
		"GovulncheckImage": c.GovulncheckImage,
		"LinterImage":      c.LinterImage,
		"WithREST":         c.WithREST,
	}

	fn := template.FuncMap{
		"lower": strings.ToLower,
		"upper": strings.ToUpper,
	}

	for name, t := range templates(c) {
		if err = createDir(t.dir); err != nil {
			return err
		}

		if tpl, err = template.New(name).Funcs(fn).Parse(t.src); err != nil {
			return err
		}

		for _, inc := range t.inc {
			if tpl, err = tpl.Parse(inc); err != nil {
				return err
			}
		}

		if f, err = os.OpenFile(t.file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(fileStrictMode)); err != nil {
			return err
		}

		if err = tpl.Execute(f, vars); err != nil {
			return err
		}
	}

	return nil
}

func createDir(d string) error {
	if d == "" {
		return nil
	}

	if err := os.MkdirAll(d, os.FileMode(dirStrictMode)); err != nil {
		return err
	}

	return nil
}
