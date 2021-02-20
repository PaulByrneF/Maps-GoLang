package main

import "fmt"

func main() {

	// Option 1: initialise, declare and assign colors to map[string]string
	// with initial key value pairs
	colors := map[string]string{
		"white": "#ffffff",
		"black": "#000000",
	}
	// Use format package: println with colors map as argument
	fmt.Println(colors)

	// Option 2: Initialise, declare variable colours as empty map
	var colours map[string]string
	fmt.Println(colours)

	//Option 3: Initialise, declare variable colorMap as empty map
	var colorMap = make(map[string]string)
	fmt.Println(colorMap)

	//Adding key value pair into existing map
	colorMap["white"] = "#ffffff"
	fmt.Println(colorMap)

	//Deleting a key value pair within a given map
	delete(colorMap, "white")
	fmt.Println(colorMap)

	//Declare variable of type map[int]string
	accounts := make(map[int]string)
	accounts[1] = "user1"
	accounts[2] = "user2"
	accounts[3] = "user3"
	accounts[4] = "user4"
	accounts[5] = "user5"

	//call function that prints accounts map
	printAccounts(accounts)
}

//function that takes in accounts: map[int]string argument
//loops over all accounts in map and prints username
func printAccounts(a map[int]string) {

	for _, user := range a {
		fmt.Println("Username: ", user)
	}

}
