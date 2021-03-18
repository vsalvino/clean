package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const Version = "1.2"

const HelpText = `NAME
  clean - The quick little project cleaner.

SYNOPSIS
  clean [options] [dir]

DESCRIPTION
  Recursively cleans compiled output files within a directory. Specify language
  modes using [options]; if omitted all language modes are cleaned. Specify a
  directory [dir]; if omitted the current directory is used.

  For sake of speed, certain directories are never cleaned such as Python
  virtual envs, node_modules, or any hidden/dot directories.

  Source code at: https://github.com/vsalvino/clean

OPTIONS
`

func main() {

	// Parse command line options.
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), HelpText)
		flag.PrintDefaults()
	}
	var c_a bool = false
	var c_cc = flag.Bool(
		"cc", false, "Clean compiled output (.exe, .dll, .o, .out, .so) etc.")
	var c_ed = flag.Bool(
		"ed", false, "Clean editor cruft (Emacs backups, etc.)")
	var c_py = flag.Bool(
		"py", false, "Clean Python __pycache__ directories.")
	var ver = flag.Bool(
		"version", false, "Print version and exit.")
	flag.Parse()
	var dir string = flag.Arg(0)
	var err error

	// Print version if specified.
	if *ver {
		fmt.Printf("clean version %s", Version)
		os.Exit(0)
	}

	// If no flags provided, use all.
	c_a = !(*c_cc || *c_ed || *c_py)

	// Process provided directory.
	if dir == "" {
		// If dir is not provided, get working directory.
		dir, err = os.Getwd()
	} else {
		// Get absolute path of what the user provided.
		dir, err = filepath.Abs(dir)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Walk the path and clean files.
	err = filepath.WalkDir(
		dir,
		func(path string, d fs.DirEntry, err error) error {

			if err != nil {
				return err
			}

			// Check if dir/file needs cleaned.
			n := d.Name()
			var delete bool = false
			if *c_py || c_a {
				// Python.
				if n == "__pycache__" ||
					strings.HasSuffix(n, ".pyc") {
					delete = true
				}
			}
			if *c_ed || c_a {
				// Emacs backups.
				if strings.HasSuffix(n, "~") ||
					(strings.HasPrefix(n, "#") && strings.HasSuffix(n, "#")) {
					delete = true
				}
			}
			if *c_cc || c_a {
				// Compiled output.
				if strings.HasSuffix(n, ".app") ||
					strings.HasSuffix(n, ".dll") ||
					strings.HasSuffix(n, ".dylib") ||
					strings.HasSuffix(n, ".exe") ||
					strings.HasSuffix(n, ".idb") ||
					strings.HasSuffix(n, ".ko") ||
					strings.HasSuffix(n, ".o") ||
					strings.HasSuffix(n, ".obj") ||
					strings.HasSuffix(n, ".out") ||
					strings.HasSuffix(n, ".pdb") ||
					strings.HasSuffix(n, ".so") ||
					strings.HasSuffix(n, ".test") {
					delete = true
				}
			}

			// Remove the file if so.
			if delete {
				fmt.Printf("Remove %s\n", path)
				err = os.RemoveAll(path)
				if err != nil {
					fmt.Printf("  Error: %s\n", err)
				}
				// If we just deleted a directory, do not try to walk it.
				if d.IsDir() {
					return filepath.SkipDir
				}
			}

			// Do not walk dirs we don't care about.
			if d.IsDir() {
				if strings.HasPrefix(n, ".") ||
					n == "env" ||
					n == "venv" ||
					n == "pyenv" ||
					n == "node_modules" {
					return filepath.SkipDir
				}
			}

			return nil
		},
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
