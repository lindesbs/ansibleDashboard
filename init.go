package main

import (
	"flag"
	"fmt"
	"os"
)

func checkInit() bool {

	flag.StringVar(&settings.importPath, "importPath", "./", "the importPath")
	flag.StringVar(&settings.exportPath, "exportPath", "./public", "the exportPath")

	flag.Parse()

	checkDirectoryStructure()

	fmt.Println("tail:", flag.Args())

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

		fmt.Println("Importing from", settings.importPath)

		return true
	}

	return false
}
