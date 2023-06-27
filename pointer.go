package main


import ("fmt")

type Orang struct {
	hobby string
}

func (man *Orang) Married(){
	man.hobby = " Painting and " + man.hobby
	fmt.Println(man.hobby)
}

func main(){
	isi := Orang{ hobby : "swimmming"}
	isi.Married()

	fmt.Println(isi.hobby)
}