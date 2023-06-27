package main

import ("fmt")


func main(){
	person := map[string]string{
		"name":"Anisa",
		"address":"jakarta",
		"job":"trainer",
	}

	var book map[string]string = make(map[string]string)
	book["title"] = " doraeman"

	fmt.Println(len(person))
 delete(person, "job")
	fmt.Println(person)
	fmt.Println(person["name"])

	
}