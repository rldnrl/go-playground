package main

import "fmt"

func main() {
  // var language string = "typescript"
  // var language = "typescript"
  language := "typescript"
  fmt.Println(language)

  var age int8 = 25
  fmt.Println(age)

  var price float32 = 3.25
  fmt.Println(price)

  // number array
  numbers := []int{10, 11, 12, 13, 14, 15}
  fmt.Println(numbers)

  // string array
  lang := []string{"ts", "js", "go", "c", "cpp", "java"}
  fmt.Println(lang)

  // array append
  numbers = append(numbers, 30)
  lang = append(lang, "md")

  fmt.Println(numbers)
  fmt.Println(lang)

  // for loop
  // i := 0
  // for i < 5 {
  //   fmt.Println("i =", i)
  //   i++
  // }

  for i := 0; i < 5; i++ {
    fmt.Println("i =", i)
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

	fmt.Println(myBill.format())
}