package main
import "fmt"

func main(){
	tmp := 25
	if tmp > 30{
		fmt.Println("Temp is greater than 30")
	}else if tmp > 25{
		fmt.Println("Temp is greater than 25 but less than or qual to 30")
	}else{
		fmt.Println("temp is less than 25")
	}

	score:=85
	if score >= 90{
		fmt.Println("Grade : A")
	}else if score >=80{
		fmt.Println("Grade B")
	}else if score >=70{
		fmt.Println("Grade C")
	}else{
		fmt.Println("Failed")
	}

	userAccess:=map[string]bool{// map of string --> bool
		"Ujjwal":true,
		"Pro":false,
	}
	if hasAccess,ok := userAccess["Ujjwal"];ok && hasAccess{
		fmt.Println("Ujjwal has access to it")
	}else{
		fmt.Println("Ujjwal dont have access")
	}
	// userAccess["Ujjwal"] --> give two value [value , isKeyExisted]
}