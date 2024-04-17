package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
)

func main() {
	// work on adding a timer
	csvReader, f := fileReader("problems.csv")
	finalCount := QuestionAnswer(csvReader)
	fmt.Println("final count is :", finalCount)
	defer f.Close()
}

func fileReader(filepath string) (*csv.Reader, *os.File) {
	// takes filepath as input and returns csv object

	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(f)

	fmt.Println(reflect.TypeOf(csvReader))

	return csvReader, f

}

func QuestionAnswer(csvReader *csv.Reader) int {
	// takes csv object as input and loop over each row using first row as a question, taking input as user, and compraing to it 2nd columsn for answer while keeping a count
	var count int
	done := make(chan bool)
	// timer1 := time.NewTimer(10 * time.Second)

	go func() {
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
			}
			value, err := strconv.Atoi(rec[1])
			if err != nil {
				log.Fatal(err)
			}
			if input == value {
				count += 1
			}

		}
		done <- true
	}()

	timer := time.NewTimer(30 * time.Second)

	select {
	case <-done:
		fmt.Println("All Questions Completed!")
		return count
	case <-timer.C:
		fmt.Println("Timer ran out!")
		return count
	}

}
