package main

import("fmt")

func hellow(name string) string{
	return "hellow" + name
}


// return multiple value
func hellow2(name string) (string, int, string){
	return "hellow" + name, 2, " version "
}

//named return value
func hellow3(name string) (fullname string, address string){
	fullname = name 
	address = "Jakarta"
	return 
}

//function as parameter
func params(name string, fungsi func(string) (string, string)){
	fmt.Println(fungsi(name))
}

//anonimous function
type Blacklist func(string) bool

func registUser(name string, black Blacklist){
	if black(name){
		fmt.Println("you are blocked ", name)
	} else {
		fmt.Println("welcome ", name)
	}
}

// func blackImpl(name string) boo

func main(){
   blackVar := func(name string) bool {
	return name == "berbie"
   }

   registUser("anis", blackVar)


   registUser("berbie", func(name string) bool{ return name =="berbie"})
// 	name := "anisa"
//    fmt.Println(hellow("anis"))
//    nama, angka, bahasa := hellow2(name)
//    fmt.Println(nama, angka, bahasa)

//    fmt.Println(hellow3("Anisa"))

   params("Nisa", hellow3)
}