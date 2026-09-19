package app

import (
	"bytes"
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
	var err error
	var tt tps

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

	tt, err = templates(c)
	if err != nil {
		return err
	}

	for _, t := range tt {
		if err = write(t, fn, vars); err != nil {
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

func write(t tp, fn template.FuncMap, vars map[string]any) error {
	var err error
	var f *os.File
	var tpl *template.Template

	if err = createDir(t.dir); err != nil {
		return err
	}

	tpl, err = template.New(t.file).Funcs(fn).Parse(t.src)
	if err != nil {
		return err
	}

	f, err = os.OpenFile(t.file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(fileStrictMode))
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	// For atomic operation write to buffer at first.
	var buf bytes.Buffer
	if err = tpl.Execute(&buf, vars); err != nil {
		return err
	}

	return os.WriteFile(t.file, buf.Bytes(), os.FileMode(fileStrictMode))
}
