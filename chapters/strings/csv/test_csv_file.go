package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var fileName string
	flag.StringVar(&fileName,"file", "", "csv file path")
	flag.Parse()

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Open file failed")
	}
	defer f.Close()

	r := csv.NewReader(f)
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(row)
	}

}
