package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var answer string
	csvfile, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println(err)
	}
	r := csv.NewReader(csvfile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("What is %v= ?\n", record[0])
		fmt.Scanln(&answer)

		if answer == record[1] {
			fmt.Println("Correct")
		} else {
			fmt.Println("Incorrect")
		}
	}
}
