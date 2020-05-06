package main

import (
	"fmt"
	"flag"
	"os"
	"path/filepath"

	"github.com/cbroglie/mustache"
)

var outputDir string
const outputDirFlag = "o"

var projectDir string
const projectDirFlag = "r"
var defaultProjectDir = filepath.Join("/", "home", "gredelston", "dev", "gscsg", "src")

func init() {
	flag.StringVar(&outputDir, outputDirFlag, "", "`Output directory` for the generated website")
	flag.StringVar(&projectDir, projectDirFlag, "", "`Project path` containing sourcefiles used to generate website")
	flag.Parse()
}

func fpExists(fp string) bool {
	_, err := os.Stat(fp)
	return !os.IsNotExist(err)
}

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

func checkProjectDir() {
	if !fpExists(projectDirFlag) {
		projectDir = defaultProjectDir
	}
	for _, subdir := range []string{"templates"} {
		if !fpExists(filepath.Join(projectDir, subdir)) {
			panic(fmt.Errorf("project directory %s does not contain subdir %s", projectDir, subdir))
		}
	}
}

func main() {
	checkOutputDir()
	checkProjectDir()

	// Load hello.mustache and render with custom data
	const basenameNoExt = "hello"
	templateFP := filepath.Join(projectDir, "templates", fmt.Sprintf("%s.mustache", basenameNoExt))
	contents, err := mustache.RenderFile(templateFP, map[string]string{"name": "Jill"})
	if err != nil {
		panic(err)
	}

	// Write rendered contents to output file
	outputFP := filepath.Join(outputDir, fmt.Sprintf("%s.html", basenameNoExt))
	err = os.Remove(outputFP)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	f, err := os.Create(outputFP)
	if err != nil {
		panic(fmt.Errorf("creating file %s: %w", outputFP, err))
	}
	defer f.Close()
	f.WriteString(contents)
}
