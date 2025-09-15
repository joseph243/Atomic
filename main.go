package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	fmt.Println("atomic data online...")
	var records = load()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Server")
	})
	fmt.Println("trying to start server")
	log.Fatal(http.ListenAndServe(":8080", nil))

	//input
	input := consoleRequest()
	fmt.Println("element " + input + " requested")
	i, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("error bad input " + err.Error())
	} else {
		fmt.Println("element " + input + " is a valid element.  building..")
	}
	freeze, err := strconv.Atoi(records[i][2])
	if err != nil {
		fmt.Println("error bad data for element " + input)
	}
	boil, err := strconv.Atoi(records[i][3])
	if err != nil {
		fmt.Println("error bad data for element " + input)
	}
	a := Element{i, records[i][1], freeze, boil}
	fmt.Println(a)
}

func consoleRequest() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("input an element atomic number: ")
	scanner.Scan()
	input := scanner.Text()
	return input
}

func load() [][]string {
	content, error := os.Open("./data.txt")
	if error != nil {
		fmt.Println("error occurred with open file: " + error.Error())
		return nil
	} else {
		fmt.Println("read file success.")
	}
	reader := csv.NewReader(content)
	records, csvErr := reader.ReadAll()
	if csvErr != nil {
		fmt.Println("error occurred with parse csv: " + csvErr.Error())
		return nil
	} else {
		fmt.Println("parse file success.")
	}
	return records
}
