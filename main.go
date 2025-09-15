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

	//hello world handler:
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Server")
	})

	//atomic number handler:
	http.HandleFunc("/element", func(w http.ResponseWriter, r *http.Request) {
		number := r.URL.Query().Get("number")
		n := parseNumber(number)
		fmt.Println("request received and processed for element " + number)
		fmt.Fprintf(w, "%v", elementByNumber(n, records))
	})

	fmt.Println("trying to start server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func elementByNumber(in int, records [][]string) Element {
	freeze := parseNumber(records[in][2])
	boil := parseNumber(records[in][3])
	return Element{in, records[in][1], freeze, boil}
}

func parseNumber(in string) int {
	i, err := strconv.Atoi(in)
	if err != nil {
		fmt.Println("error bad input " + err.Error())
	}
	return i
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
