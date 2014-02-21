package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kelseyhightower/csv2json"
)

var (
	infile string
	columns = []string{"Name","Date","Title"}
)

func init() {
	flag.StringVar(&infile, "infile", "", "input file")
}

func main() {
	flag.Parse()
	if infile == "" {
		log.Fatal("Input file required")
	}
	f, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	jsonData, err := csv2json.Convert(f, columns)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}
