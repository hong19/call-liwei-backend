package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	fmt.Printf("parse start\n")

	dat, err := ioutil.ReadFile("legislator_info.csv")
	check(err)

	r := csv.NewReader(strings.NewReader(string(dat)))

	//read out header line
	r.Read()

	i := 0
	for i = 0; i < 2; i++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}

	fmt.Printf("parse end\n")

	fmt.Printf("connect db\n")

	//db, err := sql.Open("postgres", "postgres://legislator_call@localhost/legislator_call?sslmode=verify-full")
	db, err := sql.Open("postgres", "user=legislator_call dbname=legislator_call")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("INSERT INTO legislators (id, name, gender) VALUES (2, 'hong', 'M')")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(rows)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
