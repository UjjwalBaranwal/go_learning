package main

import (
	"embed"
	"fmt"
	"log"
	"os"
)

//go:embed public
var public embed.FS

func main() {
	data, err := os.ReadFile("public/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
