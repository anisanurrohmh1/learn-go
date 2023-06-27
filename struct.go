// kumpulan field

package main

import ( "fmt")

type Customer struct {
	name, address string
	age int
}

func (customer Customer) jadiMilikStruct(){
	fmt.Println("hello world dari ", customer.name)
}

func main(){
	var pembeli Customer
	pembeli.name ="Anisa"

	var pembeli2 = Customer{
		name :"Nisa",
		address :" jakarta",
		age : 17,
	}

	fmt.Println(pembeli)
	pembeli2.jadiMilikStruct()
	pembeli.jadiMilikStruct()
	fmt.Println(pembeli2)
}