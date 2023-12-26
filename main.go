package main

import (
	"fmt"
	store "myapp/storepackage"
)

func main() {

	/*
		units := store.Units()

		fmt.Println(units)

		bill := store.NewBill()
		fmt.Println(bill)
	*/
	/*
		bill := store.NewBill()
		units := store.Units()

		ok := store.AddItem(bill, units, "carrot", "dozen")

		fmt.Println(ok)
	*/

	bill := map[string]int{"carrot": 12, "grapes": 3}
	qty, ok := store.GetItem(bill, "carrot")
	fmt.Println(qty)

	fmt.Println(ok)
}
