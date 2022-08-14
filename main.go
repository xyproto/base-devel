package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jessevdk/go-flags"
	"github.com/xyproto/env"
)

const versionString = "base-devel 0.0.1"

func main() {
	var opts struct {
		NotBin  bool `short:"n" long:"notbin" description:"Only list directories in the path that are not \"bin\" or \".bin\""`
		Version bool `short:"v" long:"version" description:"Show version innfo"`
	}
	if _, err := flags.ParseArgs(&opts, os.Args); err != nil {
		// if --help is passed, usage information has already been printed out by ParseArgs
		return
	}
	if opts.Version {
		fmt.Println(versionString)
		return
	}
	for _, directoryName := range env.Path() {
		baseFilename := filepath.Base(directoryName)
		if baseFilename == ".bin" || baseFilename == "bin" {
			if !opts.NotBin {
				fmt.Println(directoryName)
			}
		} else if opts.NotBin {
			fmt.Println(directoryName)
		}
	}
}
