//this contain simple logger
package main

import "fmt"

type LogLevel int

var levelName=[]string{"Trace","Debug","Info","Warning","Error"}

const (
	LevelTrace LogLevel=iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
)

func (l LogLevel) String() string{
	if l < LevelTrace || l > LevelError{
		return "Unknown Level"
	}  
	return levelName[l]
}

func printLogLevel(l LogLevel){
	fmt.Printf("Log Level : %d %s\n",l,l.String())
}

func main(){
	printLogLevel(LevelTrace)
	printLogLevel(LevelDebug)
	printLogLevel(LevelInfo)
	printLogLevel(LevelWarn)
	printLogLevel(LevelError)
	printLogLevel(10)
}

// let break down
/*
func (l LogLevel) String() string{
* --> this is not a function , this is method 
1) this is like we are creating a method for the LogLevel
* --> func --> defining the function (in this case method)
* --> (1 LogLevel) --> Receiver (method belong to this type)
* --> String() --> method name
* --> string --> return type
*/