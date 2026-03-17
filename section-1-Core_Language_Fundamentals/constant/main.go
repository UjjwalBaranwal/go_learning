package main
import "fmt"

const HOST string = "localhost"

const (
	Host = "127.0.0.1"
	Port=":8080"
	User = "root"
)

func main(){
	fmt.Println(HOST);
	const AppName string ="Go";
	fmt.Println(AppName)
	const pi float64 = 3.1415926
	fmt.Println(pi)
	print(Host)
}