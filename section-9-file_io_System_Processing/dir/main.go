package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir := "Downloads/static/image"
	err := os.MkdirAll(filepath.Clean(dir), 0755)
	if err != nil {
		log.Fatal(err)
	}
	err = os.RemoveAll("Downloads")
	if err != nil {
		log.Fatal(err)
	}

}
