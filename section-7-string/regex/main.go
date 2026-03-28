package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	text1 := "Ujjwal is learning Go"

	// reggo := regexp.MustCompile("[a-z") // it will panic as the given regex is false
	// use mustCompile only when you are sure that your regex is corrrect
	// safe way use Compile
	reggo, err := regexp.Compile("Go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Text : %s, matches Go : %t\n", text1, reggo.MatchString(text1))
	text2 := "Products codes: P123, X342, P789"
	rProductP := regexp.MustCompile(`P\d+`)
	firstProduct := rProductP.FindString(text2)
	fmt.Println(firstProduct)

	allPProducts := rProductP.FindAllString(text2, -1)
	fmt.Printf("%+v\n", allPProducts)
}
