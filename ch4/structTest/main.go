// structTest

package main

import (
	"fmt"
)

func main() {

	type Employee struct{
		ID	int 
		Name		string
		Phone 		string
		Address 	string
	}

	var doug Employee


	doug.ID:=1
	doug.Name:= "Douglas Will"
	doug.Phone:= "303-808-1304"
	doug.Address:= "Somewhere"



	fmt.Printf("Employee name: %s\n
				Employee Phone: %s\n
				Employee Address: %s\n
				Employee ID: %d\n",
				doug.Name, doug.Phone, doug.Address, doug.ID)
}