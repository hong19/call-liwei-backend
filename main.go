package main

import (
	"encoding/csv"
	"fmt"
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
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
