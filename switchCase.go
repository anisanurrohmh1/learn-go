 package main
 
  import (
	"fmt"
  )


  func main(){

	 var jam int
	 fmt.Println("masukkan jam anda datang bekerja : ")

	 fmt.Scan(&jam)

	 switch  {
	 case jam<9:
		fmt.Println("ontime")
	 case jam <10:
		fmt.Println("late")
		default :
		fmt.Println("Tidak disiplin!")
	 }

  }
