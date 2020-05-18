package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/cbroglie/mustache"
)

// cfgSite is the definition of a site as defined in site.cfg.
type cfgSite struct {
	Pages []cfgPage `json:"pages"`
}

// cfgPage is the definition of a singular page as defined in site.cfg.
type cfgPage struct {
	Template string `json:"template"`
	Dest string `json:"dest"`
	Data map[string]string `json:"data"`
}

// RenderSite reads the site.cfg, and renders all the pages where they belong.
func RenderSite(projectDir, outputDir string) error {
	const cfgBasename = "site.cfg"
	cfgFP := filepath.Join(projectDir, cfgBasename)
	b, err := ioutil.ReadFile(cfgFP)
	if err != nil {
		return fmt.Errorf("reading site cfg file %q: %w", cfgFP, err)
	}
	var cfg *cfgSite
	if err = json.Unmarshal(b, &cfg); err != nil {
		return fmt.Errorf("unmarshaling site cfg from %q: %w", cfgFP, err)
	}
	for _, page := range cfg.Pages {
		if page.Template == "" {
			return fmt.Errorf("site cfg contained a page with empty template: %+v", page)
		}
		templateFP := filepath.Join(projectDir, page.Template)
		outputFP := filepath.Join(outputDir, page.Dest)
		if err = RenderTemplateToFile(templateFP, outputFP, page.Data); err != nil {
			return fmt.Errorf("rendering page %q from template %q with data %+v: %w", page.Dest, page.Template, page.Data, err)
		}
	}
	return nil
}

// RenderTemplateToFile takes a Mustache template file, renders it with the given data, and writes it to a destination file.
func RenderTemplateToFile(templateFP, destFP string, context ...interface{}) error {
	// Render the mustache file.
	contents, err := mustache.RenderFile(templateFP, context...)
	if err != nil {
		return err
	}

	// Write contents to the destination.
	err = os.Remove(destFP)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	f, err := os.Create(destFP)
	if err != nil {
		return fmt.Errorf("creating file %s: %w", destFP, err)
	}
	defer f.Close()
	f.WriteString(contents)
	return nil
}
