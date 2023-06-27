package main

import (
	"fmt"
	"vskills/factoryPattern/appliances"
)

func main() {

	// Use the class factory to create an appliance of the requested type
	// Request the user to enter the appliance type
	fmt.Println("Appliance type")
	fmt.Println("0 - Stove")
	fmt.Println("1 - Fridge")
	fmt.Println("2 - Microwave")

	var myType int
	fmt.Scan(&myType)

	myAppliance, err := appliances.CreateAppliance(myType)

	if err == nil {
		myAppliance.Start()
		fmt.Println(myAppliance.GetPurpose())
	} else {
		fmt.Println(err)
	}

}
