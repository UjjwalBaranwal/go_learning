package main

import (
	"fmt"
	"strings"
)

var productPrices = map[string]float64{
	"TSHIRT": 20.00,
	"MUG":    12.50,
	"HAT":    18.00,
	"BOOK":   25.99,
}

func calcProductPrice(itemCode string) (float64, bool) {
	val, isFound := productPrices[itemCode]
	if !isFound {
		if strings.HasSuffix(itemCode,"_SALE"){
			originalItemCode:=strings.TrimSuffix(itemCode,"_SALE")
			val,isFound:=productPrices[originalItemCode]
			if isFound{
				salePrice:=val*0.90
				fmt.Printf("Original Product : %s , BASE price : %.2f, sale Price : %.2f\n",
			originalItemCode,val,salePrice)
				return salePrice,true
			}
		}
		fmt.Printf("Product: %s not in store\n",itemCode)
		return 0.0,false;
		
	}
	// if found case``
	return val, true
}

func main() {
fmt.Println("-------------- Simple Sales Order Processor------------")

	orderItems := []string{
		"TSHIRT", "MUG_SALE", "HAT", "BOOK",
	}

	var subtotal float64
	fmt.Println("-------Processing Order Items:")

	for _, item := range orderItems {
		price, found := calcProductPrice(item)
		if found {
			subtotal += price
		}
	}

	fmt.Printf("Subtotal Price: %.2f\n", subtotal)
}