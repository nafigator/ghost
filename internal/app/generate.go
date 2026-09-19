package app

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
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
	vars := map[string]any{
		"GoModule":         c.ModuleName,
		"Name":             c.Name,
		"Description":      c.Description,
		"GoImage":          c.GoImage,
		"GovulncheckImage": c.GovulncheckImage,
		"LinterImage":      c.LinterImage,
	}

	fn := template.FuncMap{
		"lower": strings.ToLower,
		"upper": strings.ToUpper,
	}

	tt, err := templates(c)
	if err != nil {
		return err
	}

	for _, t := range tt {
		if err = write(c.OutputDir, t, fn, vars); err != nil {
			return err
		}
	}

	return nil
}

func createDir(d string) error {
	if d == "" {
		return nil
	}

	if err := os.MkdirAll(d, dirStrictMode); err != nil {
		return err
	}

	return nil
}

func write(target string, t tp, fn template.FuncMap, vars map[string]any) error {
	if err := createDir(filepath.Join(target, t.dir)); err != nil {
		return err
	}

	tpl, err := template.New(t.path).Funcs(fn).Parse(t.src)
	if err != nil {
		return fmt.Errorf("parse %s: %w", t.path, err)
	}

	var buf bytes.Buffer
	if err = tpl.Execute(&buf, vars); err != nil {
		return fmt.Errorf("execute %s: %w", t.path, err)
	}

	if err = os.WriteFile(filepath.Join(target, t.path), buf.Bytes(), fileStrictMode); err != nil {
		return fmt.Errorf("write %s: %w", t.path, err)
	}

	return nil
}
