package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func updateName(x *string) {
	*x = "David"
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("Create new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add, s - save bill, t - add tip): ", reader)

	fmt.Println(opt)
}

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, error := reader.ReadString('\n')

	return strings.TrimSpace(input), error
}

func main() {

	// var firstName = "David"
	// var age uint8 = 255
	firstName := "David"
	age := 27

	fmt.Printf("First Name is %v", firstName)
	fmt.Printf("Age is %v\n", age)

	// var i = 0
	// for i < 5 {
	// 	fmt.Println("i is", i)
	// 	i++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("i is", i)
	// }

	arrayOfNumber := []int16{100, 200, 300, 400, 500, 600}

	for i, v := range arrayOfNumber {
		fmt.Println(i, "-", v)
	}

	array := []string{"David", "Jay", "Chiho"}

	for i, v := range array {
		fmt.Println(i, "-", v)
	}

	// Map
	myMap := map[string]int8{
		"Google": 1,
		"Apple":  2,
		"Meta":   3,
	}

	fmt.Println(myMap)

	for key, value := range myMap {
		fmt.Println(key, "-", value)
	}

	// Pointer
	name := "netninja"

	// updateName(name)

	fmt.Println("memory address of name is", &name)

	m := &name
	fmt.Println("memory address:", m)
	fmt.Println("value at memory address:", *m)

	updateName(m)
	fmt.Println(name)

	// Struct Custom Types
	myBill := newBill("mario's bill")
	fmt.Println(myBill)

	// add item
	myBill.addItem("onion soup", 4.50)
	myBill.addItem("veg pie", 8.95)
	myBill.addItem("coffee pudding", 4.95)
	myBill.addItem("coffee", 3.25)

	// update tip
	myBill.updateTip(10)

	fmt.Println(myBill.format())

	// User Input
	myBill2 := createBill()
	promptOptions(myBill2)
}
