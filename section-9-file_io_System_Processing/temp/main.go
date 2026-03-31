package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	tempfile, err := os.CreateTemp("", "logs.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Temp File name %s\n", tempfile.Name())
	defer func() {
		fmt.Println("Removing Tempfile", tempfile.Name())
		_ = os.RemoveAll(tempfile.Name())
	}()

	_, err = tempfile.Write([]byte("Hello World"))
	if err != nil {
		tempfile.Close()
		log.Fatal(err)
		return
	}
	// temp dir
	tempDir, err := os.MkdirTemp("", "my_app_logs")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println("Removing tempDir", tempDir)
		_ = os.Remove(tempDir)
	}()
}
