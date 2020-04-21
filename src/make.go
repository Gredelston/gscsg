package main

import (
	"fmt"
	"flag"
	"os"
	"path/filepath"
)

var outputDir string
const outputDirFlag = "o"

func init() {
	flag.StringVar(&outputDir, outputDirFlag, "", "`Output directory` for the generated website")
	flag.Parse()
}

func fpExists(fp string) bool {
	_, err := os.Stat(fp)
	return !os.IsNotExist(err)
}

func main() {
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
	const fileBasename = "index.html"
	fp := filepath.Join(outputDir, fileBasename)
	err = os.Remove(fp)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	f, err := os.Create(fp)
	if err != nil {
		panic(fmt.Errorf("creating file %s: %w", fp, err))
	}
	defer f.Close()
	const contents = "<html>\n\t<body>\n\t\tHello, world!\n\t</body>\n</html>"
	f.WriteString(contents)
}
