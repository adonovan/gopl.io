package main

import (
	"fmt"
)

func main() {

	accountName := make(map[int]string)  // Map to hold account names. Ints as keys and strings as values
	accountPhone := make(map[int]string) // Map to hold account Phone #. ints as keys and strings as values

	accountName[1] = "Douglas Will"
	accountName[2] = "Sam Spade"
	accountName[3] = "Joe Schmoe"
	accountName[4] = "Onshore Flo"

	accountPhone[1] = "303-808-1304"
	accountPhone[2] = "123-123-1234"
	accountPhone[3] = "321-321-4321"
	accountPhone[4] = "098-089-0987"

	for account, name := range accountName {

		fmt.Printf("Account %d Name: %s\n", account, name)
		fmt.Printf("         Phone: %s\n", accountPhone[account])
		fmt.Printf("\n")

	}

	selectAccount := 2

	if name, ok := accountName[selectAccount]; !ok {
		fmt.Printf("Account not found %s\n", name)
	} else {
		fmt.Printf("Account %d Name: %s\n", selectAccount, name)
	}

}
