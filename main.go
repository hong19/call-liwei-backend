package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	parseCsvToDatabase()

}

func parseCsvToDatabase() {
	fmt.Printf("parse start\n")

	//connect to db
	db := connectToDatabase()

	dat, err := ioutil.ReadFile("legislator_info.csv")
	check(err)

	inputCsv := csv.NewReader(strings.NewReader(string(dat)))

	//read out header line
	inputCsv.Read()

	i := 0
	for i = 0; i < 1; i++ {

		record := readOneRow(inputCsv)

		j := 0
		for j = 0; j < len(record); j++ {
			fmt.Printf("%d  %s\n", j, record[j])
		}

		sql := "INSERT INTO legislators (name, gender) VALUES ('hong', 'M')"

		fmt.Print(sql)

		rows, err := db.Query("INSERT INTO legislators (name, gender) VALUES ('hong', 'M')")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(rows)

	}

	fmt.Printf("parse end\n")
}

func readOneRow(inputCsv *csv.Reader) []string {
	record, err := inputCsv.Read()

	if err != nil {
		log.Fatal(err)
	}

	return record
}

func connectToDatabase() *sql.DB {
	fmt.Printf("connect db\n")

	db, err := sql.Open("postgres", "user=legislator_call dbname=legislator_call")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
