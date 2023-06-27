//slice ada potongan dari array

// terdiri dari pointer, lenght(panjang slice) dan capasitas(kapasitas dari lenght)
// lenght tidak boleh lebih dari capacity

package main

import(
	"fmt"
)

func main(){

	names := [...]string{"ani","sa","rohkmah","nur","nr","ria"}

	slice := names[4:6]
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	//fmt.Println(slice[2]) // out of range

}