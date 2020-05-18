package main

import (
	"fmt"
	"flag"
	"os"
	"path/filepath"
)

// Subdirs for the expected project structure
const templatesDirName = "templates"

// Command-line flag for where to put the generated website.
var outputDir string
const outputDirFlag = "o"

// Command-line flag for the directory containing project data.
var projectDir string
const projectDirFlag = "r"
var defaultProjectDir = filepath.Join("/", "home", "gredelston", "dev", "gscsg", "src")

// init sets up command-line flags.
func init() {
	flag.StringVar(&outputDir, outputDirFlag, "", "`Output directory` for the generated website")
	flag.StringVar(&projectDir, projectDirFlag, "", "`Project path` containing sourcefiles used to generate website")
	flag.Parse()
}

// fpExists is a convenience function to check whether a directory exists.
func fpExists(fp string) bool {
	_, err := os.Stat(fp)
	return !os.IsNotExist(err)
}

// checkOutputDir verifies that outputDir is valid, and creates it if necessary.
func checkOutputDir() {
	// Ensure validity of output directory
	if outputDir == "" {
		panic(fmt.Errorf("must supply an output directory with -%s", outputDirFlag))
	}
	outputDir, err := filepath.Abs(outputDir)
	if err != nil {
		panic(err)
	}
	if !fpExists(outputDir) {
		parentDir := filepath.Dir(outputDir)
		if !fpExists(parentDir) {
			panic(fmt.Errorf("could not find parent directory %s of specified output directory %s", parentDir, outputDir))
		} else {
			err = os.Mkdir(outputDir, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	}
}

// checkProjectDir verifies that projectDir exists; or if none is provided, swaps in the default.
func checkProjectDir() {
	if !fpExists(projectDirFlag) {
		projectDir = defaultProjectDir
	}
	for _, subdir := range []string{templatesDirName} {
		if !fpExists(filepath.Join(projectDir, subdir)) {
			panic(fmt.Errorf("project directory %s does not contain subdir %s", projectDir, subdir))
		}
	}
}

func main() {
	checkOutputDir()
	checkProjectDir()

	// Load hello.mustache and render with custom data
	if err := RenderSite(projectDir, outputDir); err != nil {
		panic(fmt.Errorf("failed to render site: %w", err))
	} else {
		fmt.Printf("Site written to %s\n", outputDir)
	}
}
