package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	var count int
	defer f.Close()
	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(rec[0])

		fmt.Println("Enter the answer for above:")
		var input int
		_, error := fmt.Scanln(&input)
		if err != nil {
			log.Fatal(error)
			return
		}
		value, err := strconv.Atoi(rec[1])
		if err != nil {
			log.Fatal(err)
		}
		if input == value {
			count += 1
		}
	}
	fmt.Println("final count is :", count)
}
