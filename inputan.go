package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var angka string 
	fmt.Print("Masukkan angkka : ")
	
	fmt.Scan(&angka)
	
	fmt.Println("Angka nya yang kamu masukan adalah : ", angka)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("masukkan kalimat : ")
	scanner.Scan()
	fmt.Println("kalimatnya adlaah "+scanner.Text()+"'")
	fmt.Println("kalimatnya adlaah "+scanner.Text()+"'")
}