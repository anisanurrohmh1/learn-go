package main 
// defer function adalah function yang di eksekusi setelah function selesai dieksuksi dan akan selalu di eksekusi walaupun sebuah function itu terjadi error
import (
	"fmt"
)

func main() {
	fmt.Println("hitung")
	for i := 0; i<10; i++{
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func 


//SELECT mst_customer.customer_name, sum(mst_menu.price) as price FROM mst_customer JOIN trx_bill ON mst_customer.id = trx_bill.customer_id JOIN trx_bill_detail ON trx_bill.id = trx_bill_detail.bill_id JOIN mst_menu ON mst_menu.id = trx_bill_detail.menu_id GROUP BY mst_customer.id ORDER BY sum(mst_menu.price) DESC  LIMIT 1;