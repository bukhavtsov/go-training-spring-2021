package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang-sql/civil"
)

type Song struct {
	Name        string
	Signer      string
	DateCreated civil.Date
}

func main() {
	const filename = "./lesson_5/error_handling/example.txt"
	createdFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error: File has not been created! Filename:", filename)
		os.Exit(1)
	}
	song := Song{
		Name:        "Slipknot",
		Signer:      "Duality",
		DateCreated: civil.Date{Year: 2004, Month: 4, Day: 4},
	}
	songBytes, err := json.Marshal(song)
	if err != nil {
		fmt.Println("Error: File has not been Marshalled:", err)
		os.Exit(1)
	}
	err = ioutil.WriteFile(filename, songBytes, os.ModeAppend)
	if err != nil {
		fmt.Println("Error: File has not been Written:", err)
		os.Exit(1)
	}
	fileBytes, err := ioutil.ReadFile(createdFile.Name())
	if err != nil {
		fmt.Println("Error: File has not been redden. Filename:", filename)
		os.Exit(1)
	}
	var songFromFile Song
	err = json.Unmarshal(fileBytes, &songFromFile)
	if err != nil {
		fmt.Println("Error: File has not been Unmarshalled:", err)
		os.Exit(1)
	}
	fmt.Printf("%v", songFromFile)
}
