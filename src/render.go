package main

import (
	"fmt"
	"os"

	"github.com/cbroglie/mustache"
)

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
