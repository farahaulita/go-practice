package main

import "fmt"

func main() {
	// This is a Comment
	
	/* This is a
	Multiline Comment */
	const (
		nextName string = "Aulita"
		familyName string = "Adinata"
	)
	

	var firstName,lastName string

	firstName,lastName = "Farah", "Sausan"
	//var firstName string = "Farah"

	//lastName := "Sausan"

	fmt.Println("Hello World")
	fmt.Printf("Halo %s %s %s %s!\n", firstName,lastName, nextName,familyName)

	_ = "Study" //unused variable
	name,_ := "John", "Wick"
	fmt.Printf("Halo %s %s!\n", name,name)

	// name := new(string) str pointer, prints out memory address, use *name to get value

	// gunakan make utk beberapa tipe data instead of var, e.g channel, slice, map

	var num1 uint8 = 1
	var num2 int8 = -2
	var decimal float32 = 2.62
	var value = (((2 + 6) % 3) * 4 - 2) / 3

	fmt.Println(num1,num2, decimal, value)

	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	var point = 10

	if point==10{
		fmt.Println("Nilai Sempurna!")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Tidak Lulus")
	}

	switch {
	case (point==10):
		fmt.Println("Amazing")
		fallthrough

	case (point > 5):
		fmt.Println("Pass")
		fallthrough

	case (point % 2 == 0):
		fmt.Println("Even")
	default:
		fmt.Println("Nope")

	}

	loop()

	
}

func loop(){
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	   }
	   
}