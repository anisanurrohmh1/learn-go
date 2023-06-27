package main

import("fmt")

func sumAll(data ...int) int {
	total := 0
 
	for _, number := range data{
	 total += number
	}
 
	return total
 }

func main(){
  fmt.Println(sumAll(2,3,4,5))
}


