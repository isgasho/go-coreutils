// +build ignore

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const path = `"github.com/EricLagerg/go-coreutils/`

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dir, err := os.Open(wd)
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	stats, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create("build_imports.go")
	if err != nil {
		panic(err)
	}

	var names []string
	for _, info := range stats {
		name := filepath.Base(info.Name())
		if name != "Godeps" &&
			info.Mode().IsDir() &&
			strings.Index(name, ".") == -1 {
			names = append(names, name)
		}
	}
	sort.Strings(names)

	outFile.WriteString("package main\n\nimport (\n")

	for _, name := range names {
		fmt.Fprintf(outFile, "\t_ %s%s\"\n", path, name)
	}
	outFile.WriteString(")")
}
