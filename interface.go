package main

import ("fmt")

type Manusia interface{
   getName() string
}
func callName(manusia Manusia){
	fmt.Println("hai ", manusia.getName())
}


type Person struct{
	fullName string
}
func (person Person) getName() string{
	return person.fullName
}

// func main(){
// 	people := Person{ fullName : "Anisa Nur ROhkmah "}
// 	callName(people)

// }


// interface kosong 
func tesIn(angka int) interface{}{
	if angka == 1 {
		return 1
	}	else if angka <1 {
		return "false"
	}	else {
		return "true"
	}
}

func main(){
	 var dataa interface{} = tesIn(3)
	 fmt.Println(dataa)
}