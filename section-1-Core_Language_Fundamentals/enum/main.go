package main

import "fmt"

const (
	Sunday=10
	Monday=iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type LogLevel int

const (
	LogError LogLevel =iota
	LogWarn
	LogInfo
	LogDebug
	LogFatal
)

func main(){
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)

	fmt.Print(LogError)
	fmt.Print(LogWarn)
	fmt.Print(LogInfo)
	fmt.Print(LogDebug)
	fmt.Print(LogFatal)
}