package main



import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)


type Teacher struct{
	id int
	name, address, gender string
}


func main(){
	 dbHost := "127.0.0.1"
	 dbPort := "5432"
	 dbName := "golang_db"
	 dbPassword := "Rahasia"
	 dbUser := "postgres"

	 dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost,dbPort, dbName)

	 db, err := sqlx.Connect("postgres", dbUrl)

	 if err!= nil {
		log.Fatalln(err)
	 }

	 defer func(){
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	 }()

	var teach Teacher

	result, err := db.Query("select teacher_id,name from m_teacher")
	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()


	for result.Next(){
		err := result.Scan(&teach.id, &teach.name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(teach)
	}

	// err = result.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	 
}


