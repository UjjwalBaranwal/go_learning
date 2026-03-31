package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// filePath := "text.txt"
	// data := "Welcome to the Go Programmnig Language!"
	// dir, _ := os.Getwd()
	// fmt.Println("Current Dir:", dir)
	// err := os.WriteFile(filePath, []byte(data), 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("done writing")
	// content, err := os.ReadFile(filePath)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(content))
	// file, err := os.Create("file-via-create.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// _, err = file.Write([]byte("Hi Gooooooooooo !!!!!"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	filename := "file-via-create.txt"
	// printContent(filename)
	newFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
	_, _ = newFile.WriteString(fmt.Sprintf("- C\n"))
	_, _ = newFile.WriteString(fmt.Sprintf("- Ruby\n"))
	_, _ = newFile.WriteString(fmt.Sprintf("- Fortan\n"))
	_, _ = newFile.WriteString(fmt.Sprintf("- Ada\n"))
	_, _ = newFile.WriteString(fmt.Sprintf("- Rust\n"))
}

func printContent(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		fmt.Println(lineNumber, scanner.Text())
		lineNumber++
	}
	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}

}
