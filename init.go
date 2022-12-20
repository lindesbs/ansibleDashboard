package main

import (
	"flag"
	"os"
)

func checkInit() bool {

	flag.StringVar(&settings.importPath, "importPath", "./data/", "the importPath")
	flag.StringVar(&settings.exportPath, "exportPath", "./public/", "the exportPath")

	flag.Parse()

	checkDirectoryStructure()

	return true
}

func checkDirectoryStructure() bool {
	if settings.importPath != "" {
		settings.importPath += "/"

		if _, err := os.Stat(settings.importPath); os.IsNotExist(err) {
			err := os.Mkdir(settings.exportPath, 0755)
			if err != nil {
				return false
			}
		}

		return true
	}

	return false
}
