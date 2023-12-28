package main

import (
	"fmt"
	// "golang.org/x/text/number"
)

// auth := 12345 this lines gives error because we are allowed to use volorous operater only inside any methods 

const Login_token = "sushil" //while declaring the public variable the first letter of the variable should always be capital.

func main() {
	// fmt.Println("variable")
	var username string = "sushil"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T \n", username)

	var is_LoggedIn bool = true
	fmt.Println(is_LoggedIn)
	fmt.Printf("The type of variable is: %T \n", is_LoggedIn)

	var age int = 23
	fmt.Println(age)
	fmt.Printf("The Type of variable is: %T \n", age)

	var income float64 = 10.00
	fmt.Println(income)
	fmt.Printf("The Type of variable is: %T \n", income)

	//implicit style
	var wesite = "sushilmakaju.com.np" //in this type we should not declare the type of variable it will denote the variable as per the data type
	fmt.Println(wesite)
	fmt.Printf("The Type of variable is: %T \n", wesite)

	// wesite = 3 now we cant declare the same name in another datatype

	//no var stype
	number_of_stds := 10000
	fmt.Println(number_of_stds)
	fmt.Printf("The type of variable is %T,\n", number_of_stds)

	fmt.Println(Login_token)
	fmt.Printf("The Type of Variable ia %T, \n", Login_token)
}
