package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

var (
	// PATH specifies the directory to clean
	PATH string
	// DRY specifies whether this is a dry run
	DRY bool
	// SAVED specifies the amount of bytes saved
	SAVED int64
	// DELETED specifies the number of files removed
	DELETED int64
	// FOUND specifies whether the GOPATH environment variable is set
	FOUND bool
	// DELETE specifies file matches to delete
	DELETE = []string{
		"LICENSE", ".appveyor.yml", ".travis.yml", "AUTHORS", "CONTRIBUTORS", "PATENTS",
	}
)

// Initialization
func init() {
	FOUND = false

	flag.BoolVar(&DRY, "dry-run", false, "no action, log files to be removed")
	flag.StringVar(&PATH, "path", goPath(), "directory to clean")

	flag.Parse()
}

// Main is the main entrypoint of the program
func main() {
	// Check to see if we were aimed at a directory
	if !FOUND && PATH == "" {
		log.Fatal("✗ GOPATH not found, please specify directory to clean using the -path flag")
	}

	// Clean
	if err := filepath.Walk(PATH, walkFn); err != nil {
		log.Fatal("✗ ", err)
	} else {
		log.Println("✓ Completed successfully")
		log.Println("🎉 Saved", SAVED, "bytes, deleted", DELETED, "files")
	}
}

// walkFn processes a file/directory
func walkFn(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	for _, match := range DELETE {
		if info.Name() == match {
			return del(path, info)
		}
	}

	return err
}

// del deletes a file
func del(path string, info os.FileInfo) error {
	SAVED += info.Size()
	DELETED++

	log.Println("[?] Deleting", path)
	if !DRY {
		return os.Remove(path)
	}

	return nil
}

// goPath determines the location of GOPATH
func goPath() string {
	path, FOUND := os.LookupEnv("GOPATH")

	if !FOUND {
		return ""
	} else {
		return path
	}
}
