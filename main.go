package main

import (
	"io/ioutil"
	"fmt"	
	//"encoding/csv"
)


func main() {

    fmt.Printf("program start\n")

    dat, err := ioutil.ReadFile("legislator_info.csv")
    check(err)
    fmt.Print(string(dat))
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}